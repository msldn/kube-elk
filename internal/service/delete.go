package service

import (
	"k8s.io/client-go/kubernetes"
)

// DeploymentCreate is a wrapper which will attempt to create and/or up a deployment.
func ServiceDelete(client *kubernetes.Clientset, namespace string, name string) error {
	return client.CoreV1().Services(namespace).Delete(name, nil)
}
