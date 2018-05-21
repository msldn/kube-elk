package elk

import (
	"fmt"
	"io/ioutil"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"strings"

	"k8s.io/api/extensions/v1beta1"
)

var namespace = "default"
var OrgId = "ORG"

var Clientset *kubernetes.Clientset

func init() {
	print("init elk")
}

type Elk struct {
	Org       string     `json: org`
	Deploy    Deployment `json: deploy`
	Svc       Service    `json: svc`
	KibanaUrl string     `json: kibanaurl`
	LogUrl    string     `json: logurl`
	Status    Status     `json: status`
}

type Deployment struct {
	Status int `json: status`
}

type Service struct {
	Status int `json: status`
}

type Status struct {
	Status int `json: status`
}

func newElk(org string) (ret *Elk) {
	return &Elk{
		Org:    org,
		Deploy: Deployment{},
		Svc:    Service{},
		Status: Status{},
	}
}

func GetConfig(name, org string) []byte {
	raw, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ret := []byte(strings.Replace(string(raw), "$ORG", org, -1))

	return ret
}

type ElkConfig struct {
	Org   string
	Kib_p int32
	Ls_p  int32
}

func (e *Elk) SetStatus() bool {
	return true
}

func (e *Elk) SetServices(svcs *apiv1.ServiceList) bool {
	e.Svc.Status = len(svcs.Items)
	return true
}

func (e *Elk) SetDeploy(deploy *v1beta1.DeploymentList) bool {
	e.Deploy.Status = len(deploy.Items)
	return true
}
