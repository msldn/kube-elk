package configmap

import (
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ConfigMapGet(client *kubernetes.Clientset, namespace string, name string) (*apiv1.ConfigMap, error) {
	return client.CoreV1().ConfigMaps(namespace).Get(name, meta_v1.GetOptions{IncludeUninitialized: true})
}
