/*
Copyright 2017 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"path/filepath"

	//appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//"k8s.io/apimachinery/pkg/api/resource"
	//"k8s.io/client-go/util/retry"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	//"k8s.io/apimachinery/pkg/util/intstr"
	//corev1 "k8s.io/api/core/v1"
	"encoding/json"
	"k8s.io/api/extensions/v1beta1"
	"testing"
	cm "github.com/marek5050/kube-elk/internal/configmap"
	deploy "github.com/marek5050/kube-elk/internal/deploy"
	svc  "github.com/marek5050/kube-elk/internal/service"
	pvc "github.com/marek5050/kube-elk/internal/pvc"
	pv "github.com/marek5050/kube-elk/internal/pv"
	"github.com/marek5050/kube-elk/internal/elk"
	"os"
)

var  clientset *kubernetes.Clientset

func init() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}else{
		os.Exit(1)
	}

	//flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
	panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
	panic(err)
}

}

//func TestCreateConfigMap(t *testing.T) {
//	raw, err := ioutil.ReadFile("github.com/marek5050/kube-elk/internal/base/kib-config.json", "testorg")
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//	var item =  &apiv1.ConfigMap{}
//	json.Unmarshal(raw, &item)
//	_,err = ConfigMapCreate(clientset, item)
//
//	if err != nil {
//		print(err)
//		t.Fatal("Failed to create ConfigMap")
//	}else{
//		print(item)
//	}
//}
var err error

func TestCreateConfigMaps(t *testing.T) {
	raw := elk.GetConfig("github.com/marek5050/kube-elk/base/kib-config.json", "testorg")

	var item =  &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)
	_,err = cm.ConfigMapCreate(clientset,"default", item)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Kibana ConfigMap")
	}

	raw = elk.GetConfig("github.com/marek5050/kube-elk/base/ls-config.json", "testorg")

	item =  &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)
	_,err = cm.ConfigMapCreate(clientset,"default", item)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Logstash ConfigMap")
	}
}

//func TestCreateDeploy(t *testing.T) {
//	raw, err := ioutil.ReadFile("github.com/marek5050/kube-elk/internal/base/kib-deploy.json", "testorg")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	var item =  &v1beta1.Deployment{}
//	json.Unmarshal(raw, &item)
//	_, err = DeploymentCreate(clientset, item)
//	if err != nil {
//		print(err)
//		t.Fatal("Failed to create Deployment")
//	}else{
//		print(item)
//	}
//}

func TestCreateDeployments(t *testing.T) {
	raw := elk.GetConfig("github.com/marek5050/kube-elk/base/kib-deploy.json", "testorg")

	var item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(clientset,"default", item)
	if err != nil {
		t.Errorf("Failed to create Kibana Deployment")
	}

	raw = elk.GetConfig("github.com/marek5050/kube-elk/base/es-deploy.json", "testorg")

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(clientset,"default", item)
	if err != nil {
		t.Errorf("Failed to create Elasticsearch Deployment")
	}

	raw = elk.GetConfig("github.com/marek5050/kube-elk/base/ls-deploy.json", "testorg")

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(clientset,"default", item)
	if err != nil {
		t.Errorf("Failed to create Logstash Deployment")
	}
}

//func TestCreateService(t *testing.T) {
//	raw, err := ioutil.ReadFile("github.com/marek5050/kube-elk/internal/base/kib-service.json", "testorg")
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//	var svc =  &apiv1.Service{}
//	json.Unmarshal(raw, &svc)
//	_,err = ServiceCreate(clientset, svc)
//
//	if err != nil {
//		print(err)
//		t.Fatal("Failed to create Service")
//	}else{
//		print(svc)
//	}
//}


func TestCreateServices(t *testing.T) {
	raw := elk.GetConfig("github.com/marek5050/kube-elk/base/kib-service.json", "testorg")

	var _svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceCreate(clientset,"default", _svc)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Kibana Service")
	}

	raw=elk.GetConfig("github.com/marek5050/kube-elk/base/es-service.json", "testorg")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceCreate(clientset,"default", _svc)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Elasticsearch Service")
	}

	raw=elk.GetConfig("github.com/marek5050/kube-elk/base/ls-service.json", "testorg")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceCreate(clientset,"default", _svc)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Logstash Service")
	}
}

func TestDeleteServices(t *testing.T) {
	raw := elk.GetConfig("github.com/marek5050/kube-elk/base/kib-service.json", "testorg")

	var _svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(clientset,"default", _svc.Name)

	if err != nil {
		t.Fatal("Failed to create Kibana Service")
	}

	raw = elk.GetConfig("github.com/marek5050/kube-elk/base/es-service.json", "testorg")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(clientset,"default", _svc.Name)

	if err != nil {
		t.Fatal("Failed to create Elasticsearch Service")
	}

	raw = elk.GetConfig("github.com/marek5050/kube-elk/base/ls-service.json", "testorg")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(clientset,"default", _svc.Name)

	if err != nil {
		t.Fatal("Failed to create Logstash Service")
	}
}

func TestElkDelete(t *testing.T) {
	elk.Clientset = clientset
	//elk.ElkDelete("default","0")
}
//
//func TestDeletePVC(t *testing.T) {
//	raw:=elk.GetConfig("github.com/marek5050/kube-elk/base/pvclaim.json", "testorg")
//
//	var svc =  &apiv1.PersistentVolumeClaim{}
//	json.Unmarshal(raw, &svc)
//	err = pvc.PVCDelete(clientset,"default", svc.Name)
//
//	if err != nil {
//		t.Fatal("Failed to delete PVClaim")
//	}
//
//	raw= elk.GetConfig("github.com/marek5050/kube-elk/base/pvstore.json", "testorg")
//
//	var _pv =  &apiv1.PersistentVolume{}
//	json.Unmarshal(raw, &_pv)
//	err = pv.PVDelete(clientset, _pv.Name)
//
//	if err != nil {
//		t.Fatal("Failed to delete PV")
//	}
//}

func TestCreatePersistentVolume(t *testing.T) {
	//raw:=elk.GetConfig("github.com/marek5050/kube-elk/base/pvstore.json", "testorg")
	//
	//var item =  &apiv1.PersistentVolume{}
	//json.Unmarshal(raw, &item)
	//_,err = pv.PVCreate(clientset, item)
	//
	//if err != nil {
	//	t.Fatal("Failed to create PVolume")
	//}
}

func TestPersistentVolumeClaim(t *testing.T) {
	rawPV:=elk.GetConfig("../../base/pvstore.json", "testorg")

	var pervol =  &apiv1.PersistentVolume{}
	json.Unmarshal(rawPV, &pervol)
	_,err = pv.PVCreate(clientset, pervol)

	if err != nil {
		t.Fatal("Failed to create PVolume: %s", err)
	}

	raw:=elk.GetConfig("../../base/pvclaim-data.json", "testorg")

	var item =  &apiv1.PersistentVolumeClaim{}
	json.Unmarshal(raw, &item)
	_,err = pvc.PVCCreate(clientset,"default", item)

	if err != nil {
		t.Fatal("Failed to create PVolumeClaim")
	}

	err = pvc.PVCDelete(clientset,"default", item.Name)
	if err != nil {
		t.Fatal("Failed to delete PV")
	}

	err = pv.PVDelete(clientset, pervol.Name)
	if err != nil {
		t.Fatal("Failed to Delete PVolume")
	}
}