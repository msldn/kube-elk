package secret

import (
	"k8s.io/client-go/kubernetes"
)

func SecretDelete(client *kubernetes.Clientset, namespace string, name string) error {
	return client.CoreV1().Secrets(namespace).Delete(name, nil)
}
