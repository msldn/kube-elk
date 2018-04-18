package service

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

// DeploymentCreate is a wrapper which will attempt to create and/or up a deployment.
func ServiceCreate(client *kubernetes.Clientset, namespace string, new *apiv1.Service) (*apiv1.Service, error) {
	svc, err := client.CoreV1().Services(namespace).Create(new)
	// We don't do anything if this is an existing resource.
	if errors.IsAlreadyExists(err) {
		return svc, nil
	}

	return svc, err
}
