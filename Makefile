all: stop start

stop:
	docker compose down

start:
	docker compose up -d

clean:
	docker images -q | xargs docker rmi -f
	docker ps -a -q | xargs docker rm
