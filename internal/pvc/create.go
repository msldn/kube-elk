package pvc

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

func PVCCreate(client *kubernetes.Clientset, namespace string, new *apiv1.PersistentVolumeClaim) (*apiv1.PersistentVolumeClaim, error) {
	svc, err := client.CoreV1().PersistentVolumeClaims(namespace).Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
