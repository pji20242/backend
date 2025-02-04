all: stop start view

stop: 
	kubectl delete -f mosquitto.yaml -n pji3
start: 
	kubectl create -f mosquitto.yaml -n pji3

view: 
	kubectl get mosquitto.yaml -n pji3
logs: 
	kubectl logs deployment/mosquitto -n pji3