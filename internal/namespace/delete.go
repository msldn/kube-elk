package configmap

import (
	"k8s.io/client-go/kubernetes"
)

func NamespaceDelete(client *kubernetes.Clientset, name string) error {
	return client.CoreV1().Namespaces().Delete(name, nil)
}
