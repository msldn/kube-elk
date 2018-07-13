package secret

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

func SecretCreate(client *kubernetes.Clientset, namespace string, new *apiv1.Secret) (*apiv1.Secret, error) {
	svc, err := client.CoreV1().Secrets(namespace).Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
