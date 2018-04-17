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
	"flag"
	"fmt"
	"os"
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
	"io/ioutil"
	"encoding/json"
	"k8s.io/api/extensions/v1beta1"
	"bytes"
	"testing"
	cm "./configmap"
	deploy "./deploy"
	svc  "./service"
	pvc "./pvc"
	pv "./pv"
	elk "./elk"
)


func getJSON(name string) []byte {
	raw, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return raw
}

func getListJSON(version, kind string, items ...[]byte) []byte {
	json := fmt.Sprintf(`{"apiVersion": %q, "kind": %q, "items": [%s]}`,
		version, kind, bytes.Join(items, []byte(",")))
	return []byte(json)
}

var  clientset *kubernetes.Clientset

func init() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	//flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
	panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
	panic(err)
}

}

//func TestCreateConfigMap(t *testing.T) {
//	raw, err := ioutil.ReadFile("./base/kib-config.json")
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
	raw := getJSON("./base/kib-config.json")

	var item =  &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)
	_,err = cm.ConfigMapCreate(clientset,"default", item)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Kibana ConfigMap")
	}

	raw = getJSON("./base/ls-config.json")

	item =  &apiv1.ConfigMap{}
	json.Unmarshal(raw, &item)
	_,err = cm.ConfigMapCreate(clientset,"default", item)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Logstash ConfigMap")
	}
}

//func TestCreateDeploy(t *testing.T) {
//	raw, err := ioutil.ReadFile("./base/kib-deploy.json")
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
	raw := getJSON("./base/kib-deploy.json")

	var item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(clientset,"default", item)
	if err != nil {
		t.Errorf("Failed to create Kibana Deployment")
	}

	raw = getJSON("./base/es-deploy.json")

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(clientset,"default", item)
	if err != nil {
		t.Errorf("Failed to create Elasticsearch Deployment")
	}

	raw = getJSON("./base/ls-deploy.json")

	item =  &v1beta1.Deployment{}
	json.Unmarshal(raw, &item)
	_, err = deploy.DeploymentCreate(clientset,"default", item)
	if err != nil {
		t.Errorf("Failed to create Logstash Deployment")
	}
}

//func TestCreateService(t *testing.T) {
//	raw, err := ioutil.ReadFile("./base/kib-service.json")
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
	raw := getJSON("./base/kib-service.json")

	var _svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceCreate(clientset,"default", _svc)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Kibana Service")
	}

	raw=getJSON("./base/es-service.json")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceCreate(clientset,"default", _svc)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Elasticsearch Service")
	}

	raw=getJSON("./base/ls-service.json")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	_,err = svc.ServiceCreate(clientset,"default", _svc)

	if err != nil {
		print(err)
		t.Fatal("Failed to create Logstash Service")
	}
}

func TestDeleteServices(t *testing.T) {
	raw := getJSON("./base/kib-service.json")

	var _svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(clientset,"default", _svc)

	if err != nil {
		t.Fatal("Failed to create Kibana Service")
	}

	raw = getJSON("./base/es-service.json")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(clientset,"default", _svc)

	if err != nil {
		t.Fatal("Failed to create Elasticsearch Service")
	}

	raw = getJSON("./base/ls-service.json")

	_svc =  &apiv1.Service{}
	json.Unmarshal(raw, &_svc)
	err = svc.ServiceDelete(clientset,"default", _svc)

	if err != nil {
		t.Fatal("Failed to create Logstash Service")
	}
}

func TestElkDelete(t *testing.T) {
	elk.Clientset = clientset
	elk.ElkDelete("default","0")
}

func TestDeletePVC(t *testing.T) {
	raw:=getJSON("./base/pvclaim.json")

	var svc =  &apiv1.PersistentVolumeClaim{}
	json.Unmarshal(raw, &svc)
	err = pvc.PVCDelete(clientset,"default", svc)

	if err != nil {
		t.Fatal("Failed to delete PVClaim")
	}

	raw= getJSON("./base/pvstore.json")

	var _pv =  &apiv1.PersistentVolume{}
	json.Unmarshal(raw, &_pv)
	err = pv.PVDelete(clientset, _pv)

	if err != nil {
		t.Fatal("Failed to delete PV")
	}
}

func TestCreatePersistentVolume(t *testing.T) {
	raw:=getJSON("./base/pvstore.json")

	var item =  &apiv1.PersistentVolume{}
	json.Unmarshal(raw, &item)
	_,err = pv.PVCreate(clientset, item)

	if err != nil {
		t.Fatal("Failed to create PVolume")
	}
}

func TestCreatePersistentVolumeClaim(t *testing.T) {
	raw:=getJSON("./base/pvclaim.json")

	var item =  &apiv1.PersistentVolumeClaim{}
	json.Unmarshal(raw, &item)
	_,err = pvc.PVCCreate(clientset,"default", item)

	if err != nil {
		t.Fatal("Failed to create PVolumeClaim")
	}
}