all: stop start

stop:
	docker compose down

start:
	docker compose up & 

clean:
	docker compose down 
	docker rm -f $(docker ps -aq) 
	docker rmi -f $(docker images -aq)
	docker volume rm $(docker volume ls -q)

erase: stop clean start