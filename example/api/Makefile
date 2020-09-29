# ============================================================================
# Globals
# ============================================================================

ECR_REPOSITORY:=418118555604.dkr.ecr.us-east-1.amazonaws.com
CONTAINER_NAME:=bflobox-api
GO_UTILS_CONTAINER_NAME:=go-utils

VERSION := $(shell git log -1 --pretty=format:"%H")
.DEFAULT_GOAL := build

# ============================================================================
# Shared/Development Commands
# ============================================================================

.PHONY: build
build:
	docker build $(CACHE_FROM) -t $(CONTAINER_NAME):$(VERSION) -t $(CONTAINER_NAME):latest .

.PHONY: build-utils
build-utils:
	docker build -t $(GO_UTILS_CONTAINER_NAME):latest -f DockerfileGoUtils .

.PHONY: run
run:
	docker-compose stop
	docker-compose up

.PHONY: lint
lint:
	$(eval LINT_OUTPUT := "/tmp/$(shell uuidgen).out")
	docker run -v $(shell pwd):/lint -w /lint -i $(GO_UTILS_CONTAINER_NAME):latest golint ./... | tee $(LINT_OUTPUT)
	test -z "$$(cat $(LINT_OUTPUT))"

# .PHONY: lint
# lint:
# 	$(eval LINT_OUTPUT := "/tmp/$(shell uuidgen).out")
# 	docker run \
# 		-v $(shell pwd):/lint \
# 		-w /lint \
# 		-i $(GO_UTILS_CONTAINER_NAME):latest golint ./... \
# 	| tee $(LINT_OUTPUT)
# 	test -z "$$(cat $(LINT_OUTPUT))"

.PHONY: test
test:
	docker run -v $(shell pwd):/test -w /test -i $(GO_UTILS_CONTAINER_NAME):latest /bin/sh -c "go test -v ./internal/handler/... -coverprofile=coverage.out | go-junit-report > report.xml"

.PHONY: coverage-html
coverage-html:
	docker run -v $(shell pwd):/test -w /test -i $(GO_UTILS_CONTAINER_NAME):latest /bin/sh -c "go tool cover -html=coverage.out -o index.html"

.PHONY: coverage-percent
coverage-percent:
	@docker run -v $(shell pwd):/test -w /test -i $(GO_UTILS_CONTAINER_NAME):latest /bin/sh -c "go tool cover -func coverage.out" | grep total | awk '{print $$3}'

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

.PHONY: push
push:
	eval $(shell aws ecr get-login --no-include-email --region us-east-1)
	docker tag $(CONTAINER_NAME):$(VERSION) $(ECR_REPOSITORY)/$(CONTAINER_NAME):$(VERSION)
	docker push $(ECR_REPOSITORY)/$(CONTAINER_NAME):$(VERSION)
