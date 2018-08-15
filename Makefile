all:
	docker-compose down
	docker-compose up -d --build
	curl "http://localhost:9628/"

logs:
	docker-compose logs -f
