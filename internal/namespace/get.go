package configmap

import (
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NamespaceGet(client *kubernetes.Clientset, name string) (*apiv1.Namespace, error) {
	return client.CoreV1().Namespaces().Get(name, meta_v1.GetOptions{IncludeUninitialized: true})
}
