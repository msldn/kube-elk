package elk

import (
	"encoding/json"
	cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	pvc "github.com/marek5050/kube-elk/internal/pvc"
	svc "github.com/marek5050/kube-elk/internal/service"
	apiv1 "k8s.io/api/core/v1"
	//pv "github.com/marek5050/kube-elk/internal/pv"
	"fmt"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/extensions/v1beta1"
)

func ServicesGet(elkconfig *ElkConfig) {
	var org = Elkconfig.Org
	raw := GetConfig("./base/kib-service.json", org)

	var _svc = &apiv1.Service{}
	var err error

	json.Unmarshal(raw, &_svc)
	_, err = svc.ServiceGet(Clientset, elkconfig.Org, _svc.Name)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: GET: Kib")
	}

	raw = GetConfig("./base/es-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_, err = svc.ServiceGet(Clientset, elkconfig.Org, _svc.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Get: ES")
	}

	raw = GetConfig("./base/ls-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_, err = svc.ServiceGet(Clientset, elkconfig.Org, _svc.Name)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Get: LS")
	}
}

func DeploymentGet(elkconfig *ElkConfig) {
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-deploy.json", org)

	var item = &v1beta1.Deployment{}
	var err error

	json.Unmarshal(raw, &item)
	item, err = deploy.DeploymentGet(Clientset, elkconfig.Org, item.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Get: Kib")
	}

	raw = GetConfig("./base/es-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	item, err = deploy.DeploymentGet(Clientset, elkconfig.Org, item.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Get: ES")
	}

	raw = GetConfig("./base/ls-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	item, err = deploy.DeploymentGet(Clientset, elkconfig.Org, item.Name)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Get: LS")
	}

}

func ConfigMapGet(elkconfig *ElkConfig) {
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-config.json", org)

	var item = &apiv1.ConfigMap{}
	var err error

	json.Unmarshal(raw, &item)
	item, err = cm.ConfigMapGet(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("CM: Get: Kibana")
	}

	raw = GetConfig("./base/ls-config.json", org)

	item = &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)
	_, err = cm.ConfigMapGet(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to Get Logstash ConfigMap")
	} else {
		log.Info("Deploy: Get: LS")
	}
}

func PVCGet(elkconfig *ElkConfig) {
	var org = Elkconfig.Org

	raw := GetConfig("./base/pvclaim-data.json", org)

	var item = &apiv1.PersistentVolumeClaim{}
	var err error

	json.Unmarshal(raw, &item)
	item, err = pvc.PVCGet(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to Get PVClaim-Data")
	} else {
		log.Info("PVC:Data: Get")
	}

	raw = GetConfig("./base/pvclaim-logs.json", org)

	item = &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCGet(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to Get PVClaim-logs")
	} else {
		log.Info("PVC:Logs: Get")
	}

	raw = GetConfig("./base/pvclaim-org.json", org)

	item = &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCGet(Clientset, elkconfig.Org, item.Name)

	if err != nil {
		log.Error("Failed to Get PVClaim-org")
	} else {
		log.Info("PVC:Org: Get")
	}

}

func ElkGet(elkconfig *ElkConfig) (*Elk, error) {
	var elkRoot = newElk(elkconfig.Org)
	svcs, _ := ServicesList(elkconfig)
	deploys, _ := DeployList(elkconfig)
	elkRoot.KibanaUrl = "http://localhost"
	elkRoot.LogUrl = "http://localhost"

	for item := range svcs.Items {
		var port = svcs.Items[item].Spec.Ports[0].Name

		if port == "5601" {
			elkRoot.KibanaUrl += fmt.Sprintf(":%d", svcs.Items[item].Spec.Ports[0].NodePort)
		} else if port == "8080" {
			elkRoot.LogUrl += fmt.Sprintf(":%d", svcs.Items[item].Spec.Ports[0].NodePort)
		}
	}

	elkRoot.SetServices(svcs)
	elkRoot.SetDeploy(deploys)

	return elkRoot, nil
}
