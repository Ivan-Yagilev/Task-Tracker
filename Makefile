build:
	docker-compose build todo-app
run:
	docker-compose up todo-app
migrate:
	migrate -path ./schema -database 'postgres://postgres:bruh@0.0.0.0:5436/postgres?sslmode=disable' up