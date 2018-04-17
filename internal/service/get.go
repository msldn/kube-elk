package service

import (
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

// ServiceGet is a wrapper
func ServiceGet(client *kubernetes.Clientset, namespace string, name string) (*apiv1.Service, error) {
	return client.CoreV1().Services(namespace).Get(name,meta_v1.GetOptions{IncludeUninitialized:true})
}