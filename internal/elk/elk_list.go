package elk

import (
	apiv1 "k8s.io/api/core/v1"
	//"encoding/json"
	//cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	svc "github.com/marek5050/kube-elk/internal/service"
	//pvc "github.com/marek5050/kube-elk/internal/pvc"
	//pv "github.com/marek5050/kube-elk/internal/pv"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/extensions/v1beta1"
)

func ServicesList() (*apiv1.ServiceList, error) {
	items, err := svc.ServiceList(Clientset, namespace, Elkconfig.Org)
	if err != nil {
		log.Info("Failed to retrieve Services List %s", err)
	} else {
		log.Info("List:Services: %s", items)
	}

	return items, err
}

func DeployList() (*v1beta1.DeploymentList, error) {
	items, err := deploy.DeploymentList(Clientset, namespace, Elkconfig.Org)
	if err != nil {
		log.Info("Failed to retrieve Deploy List %s", err)
	} else {
		log.Info("List:Deploy: %s", items)
	}
	return items, err
}

func ElkServiceList(_namespace string, elkconfig *ElkConfig) (*apiv1.ServiceList, error) {
	namespace = _namespace
	Elkconfig = elkconfig
	sl, err := ServicesList()
	if err != nil {
		print(sl)
	}

	return sl, err
}

func ElkDeployList(_namespace string, elkconfig *ElkConfig) (*v1beta1.DeploymentList, error) {
	namespace = _namespace
	Elkconfig = elkconfig
	items, err := DeployList()
	if err != nil {
		print(items)
	}
	return items, err
}
