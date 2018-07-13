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

func ServicesList(elkconfig *ElkConfig) (*apiv1.ServiceList, error) {
	items, err := svc.ServiceList(Clientset, elkconfig.Org, Elkconfig.Org)
	if err != nil {
		log.Info("Failed to retrieve Services List %s", err)
	} else {
		log.Info("List:Services: %s", items)
	}

	return items, err
}

func DeployList(elkconfig *ElkConfig) (*v1beta1.DeploymentList, error) {
	items, err := deploy.DeploymentList(Clientset, elkconfig.Org, Elkconfig.Org)
	if err != nil {
		log.Info("Failed to retrieve Deploy List %s", err)
	} else {
		log.Info("List:Deploy: %s", items)
	}
	return items, err
}

func ElkServiceList(elkconfig *ElkConfig) (*apiv1.ServiceList, error) {
	Elkconfig = elkconfig
	sl, err := ServicesList(elkconfig)
	if err != nil {
		print(sl)
	}

	return sl, err
}

func ElkDeployList(elkconfig *ElkConfig) (*v1beta1.DeploymentList, error) {
	Elkconfig = elkconfig
	items, err := DeployList(elkconfig)
	if err != nil {
		print(items)
	}
	return items, err
}
