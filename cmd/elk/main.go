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
	"github.com/marek5050/kube-elk/internal/elk"
	//"github.com/marek5050/kube-elk/internal/httplog"
	"github.com/marek5050/kube-elk/internal/web"
	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"net/http"
	"path/filepath"
)

var log logrus.Logger
var namespace = "default"
var Clientset *kubernetes.Clientset

func init() {
	//cfg := httplog.Config{
	//	MinLevel: logrus.InfoLevel,
	//}
	//h := httplog.NewHook(cfg, "http://org1.log.example.com/key/1233")
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//logrus.SetLevel(logrus.InfoLevel)
	//logrus.AddHook(h)

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	Clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	elk.Clientset = Clientset
}

func main() {

	ctx := logrus.WithFields(logrus.Fields{
		"method": "main",
	})
	ctx.Info("Starting on port 8080")

	router := web.NewRouter()

	logrus.Fatal(http.ListenAndServe(":8080", router))
}
