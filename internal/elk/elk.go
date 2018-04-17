package elk

import (
	"k8s.io/client-go/kubernetes"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

var namespace = "default"
var OrgId = "ORG"

var Clientset *kubernetes.Clientset

func init(){
	print("init elk")
}

func GetJSON(name string) []byte {
	raw, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return raw
}

func GetConfig(name, org string) []byte {
	raw, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ret := []byte(strings.Replace(string(raw), "$ORG",org, -1))

	return ret
}


type ElkConfig struct{
	Org string
	Kib_p int32
	Ls_p int32
}