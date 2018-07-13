package ingress

import (
	"k8s.io/client-go/kubernetes"
)

func IngressDelete(client *kubernetes.Clientset, namespace string, name string) error {
	return client.ExtensionsV1beta1().Ingresses(namespace).Delete(name, nil)
}
