package deploy

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

// DeploymentCreate is a wrapper which will attempt to create and/or up a deployment.
func DeploymentCreate(client *kubernetes.Clientset, namespace string, new *v1beta1.Deployment) (*v1beta1.Deployment, error) {
	dply, err := client.ExtensionsV1beta1().Deployments(namespace).Create(new)
	if errors.IsAlreadyExists(err) {
		return client.ExtensionsV1beta1().Deployments(namespace).Update(new)
	}

	return dply, err
}
