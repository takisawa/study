EXECUTABLE="app"

DOCKER_IMAGE="amazon/dynamodb-local"
DOCKER_CONTAINER="dynamodb-local"

.PHONY: all
all: build

.PHONY: build
build:
	go build -o $(EXECUTABLE) main.go

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	rm -f $(EXECUTABLE)

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

.PHONY: create-table
create-table:
	aws dynamodb create-table --table-name Persons \
		--attribute-definitions AttributeName=Id,AttributeType=N \
		--key-schema AttributeName=Id,KeyType=HASH \
		--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
		--endpoint-url http://localhost:8000
