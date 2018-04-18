package configmap

import (
	"k8s.io/client-go/kubernetes"
)

func ConfigMapDelete(client *kubernetes.Clientset, namespace string, name string) error {
	return client.CoreV1().ConfigMaps(namespace).Delete(name, nil)
}
