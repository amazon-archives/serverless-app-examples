// AWS Samples / Serverless Application Examples / Go WebSockets
//
// Go source code, Infrastructure as Code templates, build scripts, and configuration
// files for deploying a minimal example demonstrating how to use WebSockets with Amazon
// API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.
package main

import (
	"context"
	"time"

	"com.aws-samples/golang-websockets/lib/apigw"
	"com.aws-samples/golang-websockets/lib/apigw/ws"
	"com.aws-samples/golang-websockets/lib/logger"
	"com.aws-samples/golang-websockets/lib/redis"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	radix "github.com/mediocregopher/radix/v3"
	"go.uber.org/zap"
)

// cfg is the base or parent AWS configuration for this lambda.
var cfg aws.Config

// apiClient provides access to the Amazon API Gateway management functions.
var apiClient *apigatewaymanagementapi.Client

// Use the SDK default configuration, loading additional config and credentials values from the environment variables,
// shared credentials, and shared configuration files.
func init() {
	var err error
	cfg, err = external.LoadDefaultAWSConfig()
	if err != nil {
		logger.Instance.Panic("unable to load SDK config", zap.Error(err))
	}
}

func main() {
	lambda.Start(handler)
}

// handler is the hook AWS Lambda calls to invoke the function as an Amazon API Gateway Proxy. This handlers reads the
// request and echos the request back out to all connected clients. This demonstrates looking up connected clients from
// the Redis cache and calling the Amazon API Gateway Management API to send data to the connected clients.
func handler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer func() {
		_ = logger.Instance.Sync()
	}()

	// Lazily initialize the API Gateway Management client. This enables setting the service's endpoint to our API
	// endpoint. These values are provided from the synchronous request, thus the client can only be created upon the
	// first invocation.
	if apiClient == nil {
		apiClient = apigw.NewAPIGatewayManagementClient(&cfg, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	logger.Instance.Info("websocket publish",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	input, err := new(ws.InputEnvelop).Decode([]byte(req.Body))
	if err != nil {
		logger.Instance.Error("failed to parse client input",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.BadRequestResponse(), err
	}

	output := &ws.OutputEnvelop{
		Data:     input.Data,
		Type:     input.Type,
		Received: time.Now().Unix(),
	}

	data, err := output.Encode()
	if err != nil {
		logger.Instance.Error("failed to encode output",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	var connections []string
	err = redis.Client.Do(radix.Cmd(&connections, "SMEMBERS", "connections"))
	if err != nil {
		logger.Instance.Error("failed to read connections from cache",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	logger.Instance.Info("websocket connections read from cache",
		zap.Int("connections", len(connections)),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	for _, connection := range connections {
		if connection == req.RequestContext.ConnectionID && !input.Echo {
			continue
		}

		_, err := apiClient.PostToConnectionRequest(&apigatewaymanagementapi.PostToConnectionInput{
			Data:         data,
			ConnectionId: aws.String(connection),
		}).Send(ctx)

		if err != nil {
			logger.Instance.Error("failed to publish to connection",
				zap.String("receiver", connection),
				zap.String("requestId", req.RequestContext.RequestID),
				zap.String("sender", req.RequestContext.ConnectionID),
				zap.Error(err))
		}
	}

	return apigw.OkResponse(), nil
}
