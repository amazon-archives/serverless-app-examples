// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// AWS Samples / Serverless Application Examples / Go WebSockets
//
// Go source code, Infrastructure as Code templates, build scripts, and configuration
// files for deploying a minimal example demonstrating how to use WebSockets with Amazon
// API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.

// Package logger provides a singleton logging instance which is used across the same AWS Lambda execution contexts.
// Reusing the client across execution contexts provides some performance enhancements and ensures the logger
// configuration is consistent across all handlers.
package logger

import "go.uber.org/zap"

// Instance used across Lambda invocations for this execution context.
var Instance *zap.Logger

// initialize the logger used across Lambda invocations for the same execution context.
func init() {
	Instance, _ = zap.NewProduction()
}
