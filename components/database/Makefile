all: stop start view

stop: 
	kubectl delete -f Deployment.yaml -n pji3
	kubectl delete -f Service.yaml -n pji3
	kubectl delete -f configMap.yaml -n pji3
	kubectl delete -f PVC.yaml -n pji3
	kubectl delete -f PV.yaml

start: 
	kubectl create -f PV.yaml
	kubectl create -f PVC.yaml -n pji3
	kubectl create -f configMap.yaml -n pji3
	kubectl create -f Service.yaml -n pji3
	kubectl create -f Deployment.yaml -n pji3

view: 
	kubectl get configmap -n pji3
	kubectl get svc -n pji3
	# kubectl get ingress -n pji3
	kubectl get deployment -n pji3

logs: 
	kubectl logs deployment/mosquitto -n pji3
