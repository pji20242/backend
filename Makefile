all: build view

build:
	make -C kubernetes
	make -C components/mqtt-broker
view:
	kubectl get ns
	kubectl get svc -n ingress-nginx

view-all:
	kubectl get all --all-namespaces
