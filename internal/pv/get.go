package pv

import (
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func PVGet(client *kubernetes.Clientset, name string) (*apiv1.PersistentVolume, error) {
	svc, err := client.CoreV1().PersistentVolumes().Get(name, meta_v1.GetOptions{IncludeUninitialized: true})
	return svc, err
}
