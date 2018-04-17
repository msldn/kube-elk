package pv

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/api/errors"
	apiv1 "k8s.io/api/core/v1"
)

func PVCreate(client *kubernetes.Clientset, new *apiv1.PersistentVolume) (*apiv1.PersistentVolume, error) {
	svc, err := client.CoreV1().PersistentVolumes().Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
