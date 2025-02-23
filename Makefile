all: stop start

stop:
	docker compose down

start:
	docker compose up & 

clean: stop
	docker ps -a -q | xargs docker rm ; docker images -q | xargs docker rmi -f ; docker volume ls -q | xargs docker volume rm

restart-go: down-go up-go

down-go:
	docker compose down apiconnector & 
	docker compose down connector & 
	
up-go:
	docker compose up apiconnector & 
	docker compose up connector & 

erase: stop clean start