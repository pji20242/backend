all: build view

build: 
	kubectl apply -f configMap.yaml -n pji3
	kubectl apply -f Service.yaml -n pji3
	# kubectl apply -f Ingress.yaml -n pji3
	kubectl apply -f Deployment.yaml -n pji3

view: 
	kubectl get configmap -n pji3
	kubectl get svc -n pji3
	# kubectl get ingress -n pji3
	kubectl get deployment -n pji3

logs: 
	kubectl logs deployment/mosquitto -n pji3

remove: 
	kubectl delete -f Deployment.yaml -n pji3
	# kubectl delete -f Ingress.yaml -n pji3
	kubectl delete -f Service.yaml -n pji3
	kubectl delete -f configMap.yaml -n pji3