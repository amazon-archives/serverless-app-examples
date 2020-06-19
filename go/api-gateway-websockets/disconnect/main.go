// AWS Samples / Serverless Application Examples / Go WebSockets
//
// Go source code, Infrastructure as Code templates, build scripts, and configuration
// files for deploying a minimal example demonstrating how to use WebSockets with Amazon
// API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.
package main

import (
	"context"

	"com.aws-samples/golang-websockets/lib/apigw"

	"com.aws-samples/golang-websockets/lib/logger"
	"com.aws-samples/golang-websockets/lib/redis"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
	radix "github.com/mediocregopher/radix/v3"
	"go.uber.org/zap"
)

func main() {
	lambda.Start(handler)
}

// handler receives a synchronous invocation from API Gateway when a new connection has been disconnected from the
// application's API. The connection details are removed in the application's Redis cache which cleans up the connection
// details. This handler is not guaranteed to be called when the WebSocket connection is closed.
func handler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer func() {
		_ = logger.Instance.Sync()
	}()

	logger.Instance.Info("websocket disconnect",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	var result string
	err := redis.Client.Do(radix.Cmd(&result, "SREM", "connections", req.RequestContext.ConnectionID))
	if err != nil {
		logger.Instance.Error("failed to delete connection details from cache",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	logger.Instance.Info("websocket connection deleted from cache",
		zap.String("result", result),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	return apigw.OkResponse(), nil
}
