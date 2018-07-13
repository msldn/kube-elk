package elk

import (
	"encoding/json"
	cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	pvc "github.com/marek5050/kube-elk/internal/pvc"
	svc "github.com/marek5050/kube-elk/internal/service"
	ns "github.com/marek5050/kube-elk/internal/namespace"
	apiv1 "k8s.io/api/core/v1"
	//pv "github.com/marek5050/kube-elk/internal/pv"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/extensions/v1beta1"
	"github.com/marek5050/kube-elk/internal/ingress"
	"github.com/marek5050/kube-elk/internal/secret"
)

func ServicesDelete(elkconfig *ElkConfig ) {
	var org = Elkconfig.Org
	raw := GetConfig("./base/kib-service.json", org)

	var _svc = &apiv1.Service{}
	var err error

	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(Clientset, elkconfig.Org, _svc.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Delete: Kib")
	}

	raw = GetConfig("./base/es-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(Clientset, elkconfig.Org, _svc.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Delete: ES")
	}

	raw = GetConfig("./base/ls-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(Clientset, elkconfig.Org, _svc.Name)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Delete: LS")
	}
}

func DeploymentDelete(elkconfig *ElkConfig) {
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-deploy.json", org)

	var item = &v1beta1.Deployment{}
	var err error

	json.Unmarshal(raw, &item)
	err = deploy.DeploymentDelete(Clientset, elkconfig.Org, item.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Delete: Kib")
	}

	raw = GetConfig("./base/es-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	err = deploy.DeploymentDelete(Clientset, elkconfig.Org, item.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Delete: ES")
	}

	raw = GetConfig("./base/ls-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	err = deploy.DeploymentDelete(Clientset, elkconfig.Org, item.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Delete: LS")
	}

}

func ConfigMapDelete(elkconfig *ElkConfig) {
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-config.json", org)

	var item = &apiv1.ConfigMap{}
	var err error

	json.Unmarshal(raw, &item)
	err = cm.ConfigMapDelete(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("CM: Delete: Kibana")
	}

	raw = GetConfig("./base/ls-config.json", org)

	item = &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)

	err = cm.ConfigMapDelete(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to delete Logstash ConfigMap")
	} else {
		log.Info("Deploy: Delete: LS")
	}
}

func NamespaceDelete(elkconfig *ElkConfig) {
	var org = elkconfig.Org

	err := ns.NamespaceDelete(Clientset, org)

	if err != nil {
		log.Errorf("failed to delete Namespace: %d", org )
	} else {
		log.Info("Deploy: Delete: LS")
	}
}

func PVCDelete(elkconfig *ElkConfig) {
	var org = Elkconfig.Org

	raw := GetConfig("./base/pvclaim-data.json", org)

	var item = &apiv1.PersistentVolumeClaim{}
	var err error

	json.Unmarshal(raw, &item)
	err = pvc.PVCDelete(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to delete PVClaim-Data")
	} else {
		log.Info("PVC:Data: Delete")
	}

	raw = GetConfig("./base/pvclaim-logs.json", org)

	item = &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	err = pvc.PVCDelete(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to Delete PVClaim-logs")
	} else {
		log.Info("PVC:Logs: Delete")
	}

	raw = GetConfig("./base/pvclaim-org.json", org)

	item = &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	err = pvc.PVCDelete(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to Delete PVClaim-org")
	} else {
		log.Info("PVC:Org: Delete")
	}

}

func IngressDelete(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	var err error

	raw := GetConfig("./base/x-ingress.json", org)

	var _pv = &v1beta1.Ingress{}
	json.Unmarshal(raw, &_pv)
	err = ingress.IngressDelete(Clientset,org,_pv.Name)

	if err != nil {
		log.Error("failed: IngressCreate")
	} else {
		log.Info("Ingress: Create")
	}
}

func UserDelete(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	var err error

	raw := GetConfig("./base/x-useraccess.json", org)

	var _pv = &apiv1.Secret{}
	json.Unmarshal(raw, &_pv)
	err = secret.SecretDelete(Clientset,org, _pv.Name)

	if err != nil {
		log.Error("failed: UserCreate")
	} else {
		log.Info("User: Create")
	}
}


func ElkDelete(elkconfig *ElkConfig) error {
	IngressDelete(elkconfig)
	UserDelete(elkconfig)
	ServicesDelete(elkconfig)
	DeploymentDelete(elkconfig)
	ConfigMapDelete(elkconfig)
	PVCDelete(elkconfig)
	NamespaceDelete(elkconfig)
	return nil
}
