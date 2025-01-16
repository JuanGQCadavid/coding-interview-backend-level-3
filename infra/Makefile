SHELL := /bin/sh
CLOUD_REGION_NAME := $(if $(CLOUD_REGION_NAME),$(CLOUD_REGION_NAME),us-east-2)
ENV := 	$(if $(ENV),$(ENV),prod)
STACK := $(if $(STACK),$(STACK),interviews)

deploy-core:
	aws cloudformation deploy --region $(CLOUD_REGION_NAME) --template-file template.yml --stack-name $(STACK)-$(ENV) --parameter-overrides Env=$(ENV) --capabilities CAPABILITY_NAMED_IAM
