all: stop start view

stop:
	make -C components/mqtt-broker stop
	make -C kubernetes stop

start:
	make -C kubernetes start
	make -C components/mqtt-broker start

view:
	kubectl get ns
	kubectl get svc -n ingress-nginx

view-all:
	kubectl get all --all-namespaces
