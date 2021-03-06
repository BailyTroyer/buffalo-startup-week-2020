# bflobox Makefile

# ============================================================================
# Globals
# ============================================================================
DOCKER_IMAGE:=418118555604.dkr.ecr.us-east-1.amazonaws.com/bflobox
CONTAINER_IMAGE:=bflobox

S3_URL:=s3://bflobox.com

VERSION:=latest
.DEFAULT_GOAL := build

# ============================================================================
# Local Development Commands
# ============================================================================

.PHONY: build
build:
	docker build $(CACHE_FROM) -t $(DOCKER_IMAGE):$(VERSION) .

.PHONY: run
run:
	docker-compose stop
	docker-compose up


.PHONY: lint
lint:
	docker run -v $(shell pwd)/src:/node/src -v $(shell pwd)/public:/node/public $(DOCKER_IMAGE):$(VERSION) yarn lint

.PHONY: test
test:
	docker run \
		-v $(shell pwd)/src:/node/src \
		-v $(shell pwd)/public:/node/public \
		$(DOCKER_IMAGE):$(VERSION) yarn test

.PHONY: docker-push
docker-push: aws-login
	docker push $(DOCKER_IMAGE):$(VERSION)

.PHONY: build-static-assets
build-static-assets: aws-login
	docker run \
    -v $(shell pwd)/build:/node/build \
		-v $(shell pwd)/src:/node/src \
		-v $(shell pwd)/public:/node/public \
		-v $(shell pwd)/.env.production:/node/.env.production \
		$(DOCKER_IMAGE):$(VERSION) yarn build

.PHONY: push-static-assets
push-static-assets:
	aws s3 sync build $(S3_URL) --cache-control max-age=30

# ============================================================================
# CI Commands
# ============================================================================

.PHONY: aws-login
aws-login:
	@eval $(shell aws ecr get-login --no-include-email --region us-east-1)

.PHONY: pull-cache
pull-cache: aws-login
	$(eval RECENT_TAG := $(shell aws ecr describe-images --repository-name $(CONTAINER_IMAGE) --query 'reverse(sort_by(imageDetails, &imagePushedAt))[:1].imageTags[:1]' | tr -d ' []"\t\n\r\f' || true))
	$(eval CACHE_URL := $(DOCKER_IMAGE):$(RECENT_TAG))
	docker pull $(CACHE_URL) || true
	$(eval CACHE_FROM := --cache-from $(CACHE_URL))

.PHONY: clean
clean:
	rm -rf build/*
