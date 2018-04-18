package deploy

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeploymentCreate is a wrapper which will attempt to create and/or up a deployment.
func DeploymentGet(client *kubernetes.Clientset, namespace string, name string) (*v1beta1.Deployment, error) {
	return client.ExtensionsV1beta1().Deployments(namespace).Get(name, v1.GetOptions{IncludeUninitialized: true})
}
