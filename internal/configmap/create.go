package configmap

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

func ConfigMapCreate(client *kubernetes.Clientset, namespace string, new *apiv1.ConfigMap) (*apiv1.ConfigMap, error) {
	svc, err := client.CoreV1().ConfigMaps(namespace).Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
