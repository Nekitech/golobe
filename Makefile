migrate_up:
	migrate -path internal/database/migrations -database 'postgres://postgres:347389@localhost:5436/golobe_db?sslmode=disable' up
migrate_down:
	migrate -path internal/database/migrations -database 'postgres://postgres:347389@localhost:5436/golobe_db?sslmode=disable' down
docker_backend_build:
	docker compose up -d --build golobe-backend