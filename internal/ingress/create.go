package ingress

import (
	apiv1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

func IngressCreate(client *kubernetes.Clientset, namespace string, new *apiv1.Ingress) (*apiv1.Ingress, error) {
	svc, err := client.ExtensionsV1beta1().Ingresses(namespace).Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
