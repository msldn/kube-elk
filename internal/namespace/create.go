package configmap

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

func NamespaceCreate(client *kubernetes.Clientset, new *apiv1.Namespace) (*apiv1.Namespace, error) {
	svc, err := client.CoreV1().Namespaces().Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
