docker-build:
	docker build -t todo-app .
docker-up:
	docker run -d -p 8080:8080 --name todo-app todo-app
docker-down:
	docker stop todo-app
tests:
	go test ./internal/storage ./internal/service ./internal/handlers -v