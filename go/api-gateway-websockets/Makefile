# AWS Samples / Serverless Application Examples / Go WebSockets
#
# Go source code, Infrastructure as Code templates, build scripts, and configuration
# files for deploying a minimal example demonstrating how to use WebSockets with Amazon
# API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.
root	:=		$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: check clean test build deploy lint install-golangci

install-golangci:
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

lint:
	"${root}/bin/golangci-lint" run 
	cfn-lint --template "${root}/template.yml"
	cfn_nag "${root}/template.yml" --blacklist-path "${root}/.cfnnag.yml"

test:
	go test -v ./...

clean:
	$(MAKE) -C "${root}/publish" clean
	$(MAKE) -C "${root}/connect" clean
	$(MAKE) -C "${root}/disconnect" clean

build: clean
	@echo "building handlers for aws lambda"
	$(MAKE) -C "${root}/publish" build
	$(MAKE) -C "${root}/connect" build
	$(MAKE) -C "${root}/disconnect" build

deploy: check
	@echo "deploying infrastructure and code"
	sam package --output-template-file "${root}"/packaged.yml --s3-bucket "${bucket}"
	sam deploy --stack-name "${stack}" --capabilities CAPABILITY_IAM --template-file "${root}/packaged.yml"

check:
ifndef bucket
	$(error bucket was not provided)
endif
ifndef stack
	$(error stack was not provided)
endif
