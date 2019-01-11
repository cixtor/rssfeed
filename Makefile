hotfix: rebuild logs

logs:
	docker-compose logs -f --tail=1

build:
	docker-compose down
	docker-compose up -d --build
	curl "http://localhost:9628/"

rebuild:
	env GOOS=linux GOARCH=amd64 go build -o /tmp/rssfeed
	docker cp /tmp/rssfeed rssfeed:/go/bin/rssfeed
	docker-compose restart
	rm -- /tmp/rssfeed
