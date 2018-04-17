package elk

import (
	apiv1 "k8s.io/api/core/v1"
	"encoding/json"
	cm "../configmap"
	deploy "../deploy"
	svc "../service"
	pvc "../pvc"
	//pv "../pv"
	"k8s.io/api/extensions/v1beta1"
	log "github.com/sirupsen/logrus"
)

func ServicesGet(){
	var org = Elkconfig.Org
	raw := GetConfig("./base/kib-service.json", org)

	var _svc =  &apiv1.Service{}
	var err error

	json.Unmarshal(raw, &_svc)
	_, err = svc.ServiceGet(Clientset,namespace, _svc.Name)

	if err != nil {
		log.Error(err)
	}else{
		log.Info("Service: GET: Kib")
	}

	raw = GetConfig("./base/es-service.json", org)

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceGet(Clientset,namespace, _svc.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Service: Get: ES")
	}

	raw = GetConfig("./base/ls-service.json",org)

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceGet(Clientset,namespace, _svc.Name)

	if err != nil {
		log.Error(err)
	}else{
		log.Info("Service: Get: LS")
	}
}

func DeploymentGet(){
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-deploy.json",org)

	var item =  &v1beta1.Deployment{}
	var err error

	json.Unmarshal(raw, &item)
	item, err = deploy.DeploymentGet(Clientset,namespace, item.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Deploy: Get: Kib")
	}

	raw = GetConfig("./base/es-deploy.json",org)

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	item, err = deploy.DeploymentGet(Clientset,namespace, item.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Deploy: Get: ES")
	}

	raw = GetConfig("./base/ls-deploy.json",org)

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	item, err = deploy.DeploymentGet(Clientset,namespace, item.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Deploy: Get: LS")
	}

}


func ConfigMapGet(){
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-config.json",org)

		var item =  &apiv1.ConfigMap{}
		var err error

		json.Unmarshal(raw, &item)
		item, err = cm.ConfigMapGet(Clientset,namespace, item.Name)

		if err != nil {
			log.Error(err)
		}else{
			log.Info("CM: Get: Kibana")
		}

		raw = GetConfig("./base/ls-config.json",org)

		item =  &apiv1.ConfigMap{}
		json.Unmarshal(raw, &item)
		_,err = cm.ConfigMapGet(Clientset,namespace, item.Name)

		if err != nil {
			log.Error("Failed to Get Logstash ConfigMap")
		}else{
			log.Info("Deploy: Get: LS")
		}
}

func PVCGet () {
	var org = Elkconfig.Org

	raw:=GetConfig("./base/pvclaim-data.json",org)

	var item =  &apiv1.PersistentVolumeClaim{}
	var err error

	json.Unmarshal(raw, &item)
	item, err = pvc.PVCGet(Clientset,namespace, item.Name)

	if err != nil {
		log.Error("Failed to Get PVClaim-Data")
	}else{
		log.Info("PVC:Data: Get")
	}

	raw=GetConfig("./base/pvclaim-logs.json", org)

	item =  &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCGet(Clientset,namespace, item.Name)

	if err != nil {
		log.Error("Failed to Get PVClaim-logs")
	}else{
		log.Info("PVC:Logs: Get")
	}

	raw=GetConfig("./base/pvclaim-org.json", org)

	item =  &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCGet(Clientset,namespace, item.Name)

	if err != nil {
		log.Error("Failed to Get PVClaim-org")
	}else{
		log.Info("PVC:Org: Get")
	}

}

func ElkGet(_namespace string, elkconfig *ElkConfig) (error) {
	namespace = _namespace
	Elkconfig = elkconfig
	
	ServicesGet()
	DeploymentGet()
	ConfigMapGet()
	PVCGet()

	return nil
}
