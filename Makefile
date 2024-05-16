test:
	@cd ./api && go test -v ./...

up: 
	@docker-compose up --build

stop:
	@docker-compose stop

down:
	@docker-compose down
