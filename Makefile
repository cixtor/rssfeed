hotfix: rebuild register

logs:
	docker-compose logs -f

register:
	curl "http://localhost:9628/register?token=$$(cat config.json)"

build:
	docker-compose down
	docker-compose up -d --build
	curl "http://localhost:9628/"

rebuild:
	env GOOS=linux GOARCH=amd64 go build -o /tmp/rssfeed
	docker cp /tmp/rssfeed rssfeed:/go/bin/rssfeed
	docker-compose restart
	rm -- /tmp/rssfeed
