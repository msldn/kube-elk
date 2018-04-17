package pvc

import (
	"k8s.io/client-go/kubernetes"
)

func PVCDelete(client *kubernetes.Clientset, namespace string, name string) (error) {
	return client.CoreV1().PersistentVolumeClaims(namespace).Delete(name,nil)
}
