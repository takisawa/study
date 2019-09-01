EXECUTABLE="app"

DOCKER_IMAGE="amazon/dynamodb-local"
DOCKER_CONTAINER="dynamodb-local"

TABLE_NAME="ClimbingGeos"

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
	aws dynamodb create-table --table-name $(TABLE_NAME) \
		--attribute-definitions AttributeName=ClimbingId,AttributeType=S AttributeName=Time,AttributeType=S \
		--key-schema AttributeName=ClimbingId,KeyType=HASH AttributeName=Time,KeyType=RANGE \
		--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
		--endpoint-url http://localhost:8000

.PHONY: delete-table
delete-table:
	aws dynamodb delete-table --table-name $(TABLE_NAME) --endpoint-url http://localhost:8000

.PHONY: list-tables
list-tables:
	aws dynamodb list-tables --endpoint-url http://localhost:8000
