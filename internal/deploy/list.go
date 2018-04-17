package deploy

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentCreate is a wrapper which will attempt to create and/or up a deployment.
func DeploymentList(client *kubernetes.Clientset, namespace string, name string) (*v1beta1.DeploymentList, error) {
	var selector = "org=" + name
	return  client.ExtensionsV1beta1().Deployments(namespace).List(v1.ListOptions{ LabelSelector: selector})
}