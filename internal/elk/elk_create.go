package elk

import (
	"encoding/json"
	cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	pv "github.com/marek5050/kube-elk/internal/pv"
	pvc "github.com/marek5050/kube-elk/internal/pvc"
	svc "github.com/marek5050/kube-elk/internal/service"
	ns "github.com/marek5050/kube-elk/internal/namespace"
	log "github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"github.com/marek5050/kube-elk/internal/ingress"
	"github.com/marek5050/kube-elk/internal/secret"
)


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

func ServicesCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	raw := GetConfig("./base/kib-service.json", org)

	var _svc = &apiv1.Service{}
	var err error

	json.Unmarshal(raw, &_svc)
	_svc.Spec.Ports[0].NodePort = elkconfig.Kib_p

	_, err = svc.ServiceCreate(Clientset, elkconfig.Org, _svc)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: Kib")
	}

	raw = GetConfig("./base/es-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)

	_, err = svc.ServiceCreate(Clientset, elkconfig.Org, _svc)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: ES")
	}

	raw = GetConfig("./base/ls-service.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)

	_svc.Spec.Ports[0].NodePort = elkconfig.Ls_p

	_, err = svc.ServiceCreate(Clientset, elkconfig.Org, _svc)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: LS")
	}

	raw = GetConfig("./base/3-basic-auth-svc.json", org)

	_svc = &apiv1.Service{}
	json.Unmarshal(raw, &_svc)

	_, err = svc.ServiceCreate(Clientset, elkconfig.Org, _svc)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("Service: Create: Basic-Auth")
	}
}

func DeploymentCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	raw := GetConfig("./base/kib-deploy.json", org)

	var item = &v1beta1.Deployment{}
	var err error

	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, elkconfig.Org, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: Kib")
	}

	raw = GetConfig("./base/es-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, elkconfig.Org, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: ES")
	}

	raw = GetConfig("./base/ls-deploy.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, elkconfig.Org, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: LS")
	}

	raw = GetConfig("./base/3-basic-auth.json", org)

	item = &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(Clientset, elkconfig.Org, item)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Deploy: Create: Basic Auth")
	}

}

func ConfigMapCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org

	raw := GetConfig("./base/kib-config.json", org)

	var item = &apiv1.ConfigMap{}
	var err error

	json.Unmarshal(raw, &item)

	_, err = cm.ConfigMapCreate(Clientset, elkconfig.Org, item)

	if err != nil {
		log.Error(err)
	} else {
		log.Info("CM: Create: Kibana")
	}

	raw = GetConfig("./base/ls-config.json", org)

	item = &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)

	_, err = cm.ConfigMapCreate(Clientset, elkconfig.Org, item)

	if err != nil {
		log.Error("Failed to create Logstash ConfigMap")
	} else {
		log.Info("Deploy: Create: LS")
	}
}

func CreateNamespace(elkconfig *ElkConfig) {
	raw := GetConfig("./base/1-Namespace.json", elkconfig.Org)

	var item = &apiv1.Namespace{}
	var err error

	json.Unmarshal(raw, &item)

	_, err = ns.NamespaceCreate(Clientset, item)

	if err != nil {
		log.Error(err)
	} else {
		log.Infof("CM: Create: Namespace: %s", elkconfig.Org)
	}
}

func PVCCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	raw := GetConfig("./base/pvclaim-data.json", org)

	var item = &apiv1.PersistentVolumeClaim{}
	var err error

	json.Unmarshal(raw, &item)
	_, err = pvc.PVCCreate(Clientset, elkconfig.Org, item)

	if err != nil {
		log.Error("Failed to Create PVClaim-Data")
	} else {
		log.Info("PVC:Data: Create")
	}

	//raw = GetConfig("./base/pvclaim-logs.json", org)
	//
	//item = &apiv1.PersistentVolumeClaim{}
	//
	//json.Unmarshal(raw, &item)
	//_, err = pvc.PVCCreate(Clientset, elkconfig.Org, item)
	//
	//if err != nil {
	//	log.Error("Failed to Create PVClaim-Logs")
	//} else {
	//	log.Info("PVC:Logs: Create")
	//}
	//
	//raw = GetConfig("./base/pvclaim-org.json", org)
	//
	//item = &apiv1.PersistentVolumeClaim{}
	//
	//json.Unmarshal(raw, &item)
	//_, err = pvc.PVCCreate(Clientset, elkconfig.Org, item)
	//
	//if err != nil {
	//	log.Error("Failed to Create PVClaim-Org")
	//} else {
	//	log.Info("PVC:Org: Create")
	//}
}

func PVCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org
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

func IngressCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	var err error

	raw := GetConfig("./base/x-ingress.json", org)

	var _pv = &v1beta1.Ingress{}
	json.Unmarshal(raw, &_pv)
	_, err = ingress.IngressCreate(Clientset,org, _pv)

	if err != nil {
		log.Error("failed: IngressCreate")
	} else {
		log.Info("Ingress: Create")
	}
}

func UserCreate(elkconfig *ElkConfig) {
	var org = elkconfig.Org
	var err error

	raw := GetConfig("./base/x-useraccess.json", org)

	var _pv = &apiv1.Secret{}
	json.Unmarshal(raw, &_pv)
	_, err = secret.SecretCreate(Clientset,org, _pv)

	if err != nil {
		log.Error("failed: UserCreate")
	} else {
		log.Info("User: Create")
	}
}

func ElkCreate(elkconfig *ElkConfig) error {
	CreateNamespace(elkconfig)
	ConfigMapCreate(elkconfig)
	//PVCreate(elkconfig)
	PVCCreate(elkconfig)
	DeploymentCreate(elkconfig)
	ServicesCreate(elkconfig)
	UserCreate(elkconfig)
	IngressCreate(elkconfig)
	return nil
}
