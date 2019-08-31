DOCKER_IMAGE="amazon/dynamodb-local"
DOCKER_CONTAINER="dynamodb-local"

.PHONY: all
all: build

.PHONY: build
build:
	go build -o main main.go

.PHONY: docker-create
docker-create:
	docker container create -p 8000:8000 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

.PHONY: docker-rm
docker-rm:
	docker container rm $(DOCKER_CONTAINER)

.PHONY: docker-start
docker-start:
	docker container start $(DOCKER_CONTAINER)

.PHONY: docker-stop
docker-stop:
	docker container stop $(DOCKER_CONTAINER)
