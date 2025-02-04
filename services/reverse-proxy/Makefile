all: stop start view

stop: 
	kubectl delete -f proxy.yaml -n pji3
start: 
	kubectl create -f proxy.yaml -n pji3

view: 
	kubectl get proxy.yaml -n pji3
logs: 
	kubectl logs deployment/proxy -n pji3