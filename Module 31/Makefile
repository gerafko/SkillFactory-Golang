docker-build:
	docker build . -t anton-go-service

up: docker-build
	docker-compose up --build

migrate-up:
	goose --dir migrations postgres "host=localhost port=5432 dbname=postgres user=postgres password=postgres" up

migrate-down:
	goose --dir migrations postgres "host=localhost port=5432 dbname=postgres user=postgres password=postgres" down
