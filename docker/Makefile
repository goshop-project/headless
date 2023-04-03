.PHONY: all prepare
.PHONY: build build-docker build-bin

DOCKER_COMPOSE ?= docker-compose
DOCKER_COMPOSE_NAMESPACE ?= headless

DOCKER_COMPOSE_CALL := $(DOCKER_COMPOSE) -p $(DOCKER_COMPOSE_NAMESPACE)

NODES := ca auth bff

all: build-bin

prepare:
	@mkdir -p $(addsuffix -data, $(NODES)) bin

build: build-docker build-bin

build-docker: prepare
	@$(DOCKER_COMPOSE_CALL) build --no-cache --parallel

build-bin: prepare
	@$(MAKE) -C $(CURDIR)/.. OUTDIR=$(CURDIR)/bin

.PHONY: dev start down logs ps
export COMPOSE_PROJECT_NAME=headless

dev:
	@$(DOCKER_COMPOSE_CALL) up

start:
	@$(DOCKER_COMPOSE_CALL) up -d

stop:
	@$(DOCKER_COMPOSE_CALL) down

logs:
	@$(DOCKER_COMPOSE_CALL) logs -f

ps:
	@$(DOCKER_COMPOSE_CALL) ps
