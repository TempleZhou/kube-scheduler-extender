package kubeclientset

import (
	clientset "k8s.io/client-go/kubernetes"
)

var (
	Client   *clientset.Clientset
	Complete = make(chan int, 1)
)
