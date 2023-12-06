mod-gen:
	@echo "Running generate go mod..."
	go mod tidy && go mod vendor -v

docker-start:
	@echo "Running docker compose up..."
	docker-compose up

docker-stop:
	@echo "Running docker compose down..."
	docker-compose down

docker-remove:
	@echo "Running remove docker compose..."
	docker-compose down -v