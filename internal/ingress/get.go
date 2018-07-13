package ingress

import (
	apiv1 "k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func IngressGet(client *kubernetes.Clientset, namespace string, name string) (*apiv1.Ingress, error) {
	return client.ExtensionsV1beta1().Ingresses(namespace).Get(name, meta_v1.GetOptions{IncludeUninitialized: true})
}
