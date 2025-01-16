SHELL := /bin/sh
CLOUD_REGION_NAME := $(if $(CLOUD_REGION_NAME),$(CLOUD_REGION_NAME),us-east-2)
ENV := 	$(if $(ENV),$(ENV),prod)
STACK := $(if $(STACK),$(STACK),interviews)
PLATFORM := $(if $(PLATFORM),$(PLATFORM),amd64)
VERSION := $(if $(VERSION),$(VERSION),)
ECR_URI := $(shell aws cloudformation --region $(CLOUD_REGION_NAME) describe-stacks --stack-name interviews-prod --query  "Stacks[0].Outputs[?ExportName=='interview-ecr-uri'].OutputValue" --output text)

deploy:
	aws cloudformation deploy --region $(CLOUD_REGION_NAME) --template-file template.yml --stack-name $(STACK)-items-$(ENV) --parameter-overrides Env=$(ENV) ServiceImageVersion=$(VERSION) --capabilities CAPABILITY_NAMED_IAM

build:
	docker build -t $(STACK)-$(ENV)-$(VERSION) .

ecr-login:
	aws ecr get-login-password --region $(CLOUD_REGION_NAME) | docker login --username AWS --password-stdin 732596568988.dkr.ecr.$(CLOUD_REGION_NAME).amazonaws.com

ecr-deploy: ecr-login
	docker buildx build --platform=linux/$(PLATFORM) --build-arg CMD="lambda" --build-arg GOARCH="$(PLATFORM)" --provenance=false -t $(STACK):$(VERSION) --load . \
	&& docker tag $(STACK):$(VERSION) $(ECR_URI):$(VERSION) \
	&& docker push $(ECR_URI):$(VERSION)

docker-buildx-driver:
	docker buildx create --use

publish: ecr-deploy deploy

run-http:
	userDB=theAdmin passDB=MyVeryPrivatePassword21 dbnameDB=itemsDB urlDB=localhost go run cmd/http/main.go




