all: build register

logs:
	docker-compose logs -f

register:
	curl "http://localhost:9628/register?token=$$(cat config.json)"

build:
	docker-compose down
	docker-compose up -d --build
	curl "http://localhost:9628/"
