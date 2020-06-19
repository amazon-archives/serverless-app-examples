// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// AWS Samples / Serverless Application Examples / Go WebSockets
//
// Go source code, Infrastructure as Code templates, build scripts, and configuration
// files for deploying a minimal example demonstrating how to use WebSockets with Amazon
// API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.

// Package redis provides a singleton Redis client instance which is used across the same AWS Lambda execution contexts.
// Reusing the client across execution contexts provides some performance enhancements due to reusing the underlying
// connection to the Redis cluster.
package redis

import (
	"net"
	"os"

	"com.aws-samples/golang-websockets/lib/logger"
	"github.com/mediocregopher/radix/v3"
	"go.uber.org/zap"
)

// Client is the single client instance shared across the same Lambda execution contexts.
var Client *radix.Pool

func init() {
	var err error
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	Client, err = radix.NewPool("tcp", net.JoinHostPort(host, port), 1)
	if err != nil {
		logger.Instance.Panic("unable to create redis connection pool", zap.Error(err))
	}
}
