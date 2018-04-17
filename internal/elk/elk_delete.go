package elk

import (
	apiv1 "k8s.io/api/core/v1"
	"encoding/json"
	cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	svc "github.com/marek5050/kube-elk/internal/service"
	pvc "github.com/marek5050/kube-elk/internal/pvc"
	//pv "github.com/marek5050/kube-elk/internal/pv"
	"k8s.io/api/extensions/v1beta1"
	log "github.com/sirupsen/logrus"
)

func ServicesDelete(){
	var org = Elkconfig.Org
	raw := GetConfig("./base/kib-service.json", org)

	var _svc =  &apiv1.Service{}
	var err error

	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(Clientset,namespace, _svc.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Service: Delete: Kib")
	}

	raw = GetConfig("./base/es-service.json", org)

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(Clientset,namespace, _svc.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Service: Delete: ES")
	}

	raw = GetConfig("./base/ls-service.json",org)

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(Clientset,namespace, _svc.Name)

	if err != nil {
		log.Error(err)
	}else{
		log.Info("Service: Delete: LS")
	}
}

func DeploymentDelete(){
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-deploy.json",org)

	var item =  &v1beta1.Deployment{}
	var err error

	json.Unmarshal(raw, &item)
	err = deploy.DeploymentDelete(Clientset,namespace, item.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Deploy: Delete: Kib")
	}

	raw = GetConfig("./base/es-deploy.json",org)

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	err = deploy.DeploymentDelete(Clientset,namespace, item.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Deploy: Delete: ES")
	}

	raw = GetConfig("./base/ls-deploy.json",org)

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	err = deploy.DeploymentDelete(Clientset,namespace, item.Name)
	if err != nil {
		log.Error(err)
	}else{
		log.Info("Deploy: Delete: LS")
	}

}


func ConfigMapDelete(){
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-config.json",org)

		var item =  &apiv1.ConfigMap{}
		var err error

		json.Unmarshal(raw, &item)
		err = cm.ConfigMapDelete(Clientset,namespace, item.Name)

		if err != nil {
			log.Error(err)
		}else{
			log.Info("CM: Delete: Kibana")
		}

		raw = GetConfig("./base/ls-config.json",org)

		item =  &apiv1.ConfigMap{}
		json.Unmarshal(raw, &item)

		err = cm.ConfigMapDelete(Clientset,namespace, item.Name)

		if err != nil {
			log.Error("Failed to delete Logstash ConfigMap")
		}else{
			log.Info("Deploy: Delete: LS")
		}
}

func PVCDelete () {
	var org = Elkconfig.Org

	raw:=GetConfig("./base/pvclaim-data.json",org)

	var item =  &apiv1.PersistentVolumeClaim{}
	var err error

	json.Unmarshal(raw, &item)
	err = pvc.PVCDelete(Clientset,namespace, item.Name)

	if err != nil {
		log.Error("Failed to delete PVClaim-Data")
	}else{
		log.Info("PVC:Data: Delete")
	}

	raw=GetConfig("./base/pvclaim-logs.json", org)

	item =  &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	err = pvc.PVCDelete(Clientset,namespace, item.Name)

	if err != nil {
		log.Error("Failed to Delete PVClaim-logs")
	}else{
		log.Info("PVC:Logs: Delete")
	}

	raw=GetConfig("./base/pvclaim-org.json", org)

	item =  &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	err = pvc.PVCDelete(Clientset,namespace, item.Name)

	if err != nil {
		log.Error("Failed to Delete PVClaim-org")
	}else{
		log.Info("PVC:Org: Delete")
	}

}

func ElkDelete(_namespace string, elkconfig *ElkConfig) (error) {
	namespace = _namespace
	Elkconfig = elkconfig
	
	ServicesDelete()
	DeploymentDelete()
	ConfigMapDelete()
	PVCDelete()

	return nil
}
