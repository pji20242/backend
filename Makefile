all: stop start

stop:
	docker compose down

start:
	docker compose up & 

clean:
	docker images -q | xargs docker rmi -f
	docker ps -a -q | xargs docker rm
	docker volume rm $(docker volume ls -q)

erase: stop clean start