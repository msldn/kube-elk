package deploy

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeploymentCreate is a wrapper which will attempt to create and/or up a deployment.
func DeploymentDelete(client *kubernetes.Clientset, namespace string, name string) error {
	var prop v1.DeletionPropagation = "Foreground"
	err := client.ExtensionsV1beta1().Deployments(namespace).Delete(name, &v1.DeleteOptions{
		PropagationPolicy: &prop,
	})
	return err
}
