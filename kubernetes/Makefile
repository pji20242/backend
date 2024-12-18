# Ver https://github.com/kubernetes-csi/csi-driver-nfs/blob/master/deploy/example/README.md#storage-class-usage-dynamic-provisioning

all: stop start view

stop:
	kubectl delete -f ns.yaml
	# kubectl delete -f ingressController.yaml
	# kubectl delete -f storageClass.yaml
	#
	# Ver https://github.com/kubernetes-csi/csi-driver-nfs/blob/master/docs/install-csi-driver-master.md#clean-up-nfs-csi-driver
	# curl -skSL https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/uninstall-driver.sh | bash -s master --

start:
	# Ver https://github.com/kubernetes-csi/csi-driver-nfs/blob/master/docs/install-csi-driver-master.md#install-with-kubectl
	# curl -skSL https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/install-driver.sh | bash -s master --
	# kubectl create -f storageClass.yaml
	# kubectl create -f ingressController.yaml
	kubectl create -f ns.yaml

view:
	kubectl get ns
	kubectl get svc -n ingress-nginx

view-all:
	kubectl get all --all-namespaces
