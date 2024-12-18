all: stop start view

stop:
	kubectl delete -f ns.yaml
	# kubectl delete -f ingressController.yaml

start:
	# kubectl create -f ingressController.yaml
	kubectl create -f ns.yaml

view:
	kubectl get ns
	kubectl get svc -n ingress-nginx

view-all:
	kubectl get all --all-namespaces


# all: build view
# 
# build:
# 	kubectl apply -f ingressController.yaml
# 	kubectl apply -f ns.yaml
# 	kubectl apply -f persistentVolume.yaml
# view:
# 	kubectl get ns
# 	kubectl get svc -n ingress-nginx
# 	kubectl get pv
# 
# view-all:
# 	kubectl get all --all-namespaces