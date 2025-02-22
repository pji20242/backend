all: stop start

stop:
	docker compose down

start:
	docker compose up & 

clean:
	docker ps -a -q | xargs docker rm ; docker images -q | xargs docker rmi -f ; docker volume ls -q | xargs docker volume rm
	
restart:
	docker compose restart apiconnector

erase: stop clean start