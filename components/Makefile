all: stop start

stop:
	kubectl delete -f database.yaml

start:
	kubectl create -f database.yaml
