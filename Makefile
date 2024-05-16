build:
	@go build -o bin/api

run: build
	@./bin/api

test:
	@go test -v ./...

up: 
	@docker-compose up --build

stop:
	@docker-compose stop

down:
	@docker-compose down
