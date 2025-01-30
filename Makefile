api: docker
	air

docker:
	docker-compose up -d

web: 
	cd frontend && npm run dev 	

stop:
	docker-compose down

clean:
	docker-compose down --volumes --remove-orphans

docs:
	swag fmt && swag init -g cmd/api/main.go

test:
	go test ./... -v -cover

.PHONY: web docs test
