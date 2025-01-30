api: docker
	air

docker:
	docker-compose up -d

stop:
	docker-compose down

clean:
	docker-compose down --volumes --remove-orphans

docs:
	swag fmt && swag init -g cmd/api/main.go

test:
	go test ./... -v -cover

.PHONY: docs test
