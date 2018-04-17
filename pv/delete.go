package pv

import (
	"k8s.io/client-go/kubernetes"
)

func PVDelete(client *kubernetes.Clientset,name string) (error) {
	return client.CoreV1().PersistentVolumes().Delete(name,nil)
}
