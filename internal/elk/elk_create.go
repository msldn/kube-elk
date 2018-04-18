package elk

import (
	"encoding/json"
	cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	pv "github.com/marek5050/kube-elk/internal/pv"
	pvc "github.com/marek5050/kube-elk/internal/pvc"
	svc "github.com/marek5050/kube-elk/internal/service"
	log "github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

var Elkconfig *ElkConfig

func init() {
	//print("init elk_create\n")
	//cfg := httplog.Config{
	//	MinLevel:       logrus.InfoLevel,
	//}
	//
	//h := httplog.NewHook(cfg, "http://192.168.99.100:31523/key/elk_Create")
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.InfoLevel)
	//log.AddHook(h)
}

func ServicesCreate() {
	var org = Elkconfig.Org
	raw := GetConfig("./base/kib-service.json", org)

	var _svc = &apiv1.Service{}
	var err error

	json.Unmarshal(raw, &_svc)
	_svc.Spec.Ports[0].NodePort = Elkconfig.Kib_p

	_, err = svc.ServiceCreate(Clientset, namespace, _svc)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: Kib")
	}

	raw = GetConfig("./base/es-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)

	_, err = svc.ServiceCreate(Clientset, namespace, _svc)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: ES")
	}

	raw = GetConfig("./base/ls-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)

	_svc.Spec.Ports[0].NodePort = Elkconfig.Ls_p

	_, err = svc.ServiceCreate(Clientset, namespace, _svc)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: LS")
	}
}

func DeploymentCreate() {
	var org = Elkconfig.Org
	raw := GetConfig("./base/kib-deploy.json", org)

	var item = &v1beta1.Deployment{}
	var err error

	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, namespace, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: Kib")
	}

	raw = GetConfig("./base/es-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, namespace, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: ES")
	}

	raw = GetConfig("./base/ls-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, namespace, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: LS")
	}

}

func ConfigMapCreate() {
	var org = Elkconfig.Org

	raw := GetConfig("./base/kib-config.json", org)

	var item = &apiv1.ConfigMap{}
	var err error

	json.Unmarshal(raw, &item)

	_, err = cm.ConfigMapCreate(Clientset, namespace, item)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("CM: Create: Kibana")
	}

	raw = GetConfig("./base/ls-config.json", org)

	item = &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)

	_, err = cm.ConfigMapCreate(Clientset, namespace, item)

	if err != nil {
		log.Error("Failed to create Logstash ConfigMap")
	} else {
		log.Info("Deploy: Create: LS")
	}
}

func PVCCreate() {
	var org = Elkconfig.Org
	raw := GetConfig("./base/pvclaim-data.json", org)

	var item = &apiv1.PersistentVolumeClaim{}
	var err error

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCCreate(Clientset, namespace, item)

	if err != nil {
		log.Error("Failed to Create PVClaim-Data")
	} else {
		log.Info("PVC:Data: Create")
	}

	raw = GetConfig("./base/pvclaim-logs.json", org)

	item = &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCCreate(Clientset, namespace, item)

	if err != nil {
		log.Error("Failed to Create PVClaim-Logs")
	} else {
		log.Info("PVC:Logs: Create")
	}

	raw = GetConfig("./base/pvclaim-org.json", org)

	item = &apiv1.PersistentVolumeClaim{}

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCCreate(Clientset, namespace, item)

	if err != nil {
		log.Error("Failed to Create PVClaim-Org")
	} else {
		log.Info("PVC:Org: Create")
	}
}

func PVCreate() {
	var org = Elkconfig.Org
	var err error

	raw := GetConfig("./base/pvstore.json", org)

	var _pv = &apiv1.PersistentVolume{}
	json.Unmarshal(raw, &_pv)
	_, err = pv.PVCreate(Clientset, _pv)

	if err != nil {
		log.Error("Failed to Create PV")
	} else {
		log.Info("PV: Create")
	}
}

func ElkCreate(_namespace string, elkconfig *ElkConfig) error {
	namespace = _namespace
	Elkconfig = elkconfig
	ConfigMapCreate()
	PVCreate()
	PVCCreate()
	DeploymentCreate()
	ServicesCreate()

	return nil
}
