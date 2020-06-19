// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// AWS Samples / Serverless Application Examples / Go WebSockets
//
// Go source code, Infrastructure as Code templates, build scripts, and configuration
// files for deploying a minimal example demonstrating how to use WebSockets with Amazon
// API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.

// Package apigw provides common resources for working with Amazon API Gateway
package apigw

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Response is a typedef for the response type provided by the SDK.
type Response = events.APIGatewayProxyResponse

// InternalServerErrorResponse returns an Amazon API Gateway Proxy Response configured with the correct HTTP status
// code.
func InternalServerErrorResponse() Response {
	return Response{StatusCode: http.StatusInternalServerError}
}

// BadRequestResponse returns an Amazon API Gateway Proxy Response configured with the correct HTTP status code.
func BadRequestResponse() Response {
	return Response{StatusCode: http.StatusBadRequest}
}

// OkResponse returns an Amazon API Gateway Proxy Response configured with the correct HTTP status code.
func OkResponse() Response {
	return Response{StatusCode: http.StatusOK}
}
