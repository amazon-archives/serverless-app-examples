# AWS WebSockets with Golang and Redis

This project contains a reference implementation for using AWS VPC, Amazon API Gateway WebSockets, AWS Lambda and Amazon ElastiCache for Redis.

Three AWS Lambda handlers are included in the project:

- **ConnectFunction**: Invoked by API Gateway when a new WebSocket connection is established. The connection information is cached in the ElastiCache for Redis instance.

- **DisconnectFunction**: Invoked by API Gateway when a WebSocket connection is terminated. The connection information is removed from the ElastiCache for Redis instance.

- **PublishFunction**: Invoked by API Gateway when data is sent from the client over the WebSocket connection. The data is "published" to all connected clients.

## Building and Deploying

### Compilation

The AWS Lambda handlers are written in Go. See the following link for more information about the Go programming language including installation instructions.

Go >= 1.13 is required.

<https://golang.org/>

Compilation of the AWS Lambda handlers is managed with the included Makefile. The Makefile ensures the binaries are cross-compiled to run in AWS Lambda. Run the following command to build the binaries for deployment.

```bash
make clean build
```

### Deploying

Deployments are managed using the AWS Serverless Application Model. See the following links for more information about AWS SAM including installation instructions.

<https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html>

<https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html>

The included Makefile includes a target for deploying the infrastructure and code using AWS SAM. The following example demonstrates how to deploy the infrastructure and code.

```bash
AWS_PROFILE={profile} AWS_DEFAULT_REGION={region} make bucket={deployment bucket} stack=severless-app-examples-go-websockets deploy
```

## Using wscat for Testing

Query the deployed WebSocket endpoint by running the following command:

```bash
aws cloudformation describe-stacks --stack-name severless-app-examples-go-websockets --query "Stacks[0].Outputs[7].OutputValue" --output text
```

Now knowing the WebSocket endpoint, wscat can be used to test the deployment. Establish a WebSocket connection with the following command, and the endpoint value from the previous step:

```bash
wscat -c {endpoint}
```

The "backend service" only accepts messages which adhere to the below schema. Use the following format To publish a message to all connected clients:

```json
{ "echo": false, "type": 99, "data": "data to publish" }
```
