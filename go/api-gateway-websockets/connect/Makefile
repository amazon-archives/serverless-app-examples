# AWS Samples / Serverless Application Examples / Go WebSockets
#
# Go source code, Infrastructure as Code templates, build scripts, and configuration
# files for deploying a minimal example demonstrating how to use WebSockets with Amazon
# API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.

.PHONY: clean build

clean:
	rm -rfv bin

build:
	 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/handler