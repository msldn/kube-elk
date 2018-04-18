package pvc

import (
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func PVCGet(client *kubernetes.Clientset, namespace string, name string) (*apiv1.PersistentVolumeClaim, error) {
	svc, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(name, meta_v1.GetOptions{IncludeUninitialized: true})
	return svc, err
}
