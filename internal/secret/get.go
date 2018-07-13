package secret

import (
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func SecretGet(client *kubernetes.Clientset, namespace string, name string) (*apiv1.Secret, error) {
	return client.CoreV1().Secrets(namespace).Get(name, meta_v1.GetOptions{IncludeUninitialized: true})
}
