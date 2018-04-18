package service

import (
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceGet is a wrapper
func ServiceList(client *kubernetes.Clientset, namespace string, name string) (*apiv1.ServiceList, error) {
	var selector = "org=" + name
	return client.CoreV1().Services(namespace).List(meta_v1.ListOptions{LabelSelector: selector})
}
