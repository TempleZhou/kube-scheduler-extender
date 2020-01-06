package detecttoomanypodsinitializing

import (
	"github.com/prometheus/common/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-scheduler-extender/kubeclientset"
)

func DetectTooManyPodsInitializing(node v1.Node, maxInitializingPods int) (res bool) {
	clientset := kubeclientset.Client
	cnt := 0

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{FieldSelector: "spec.nodeName=" + node.Name})
	if err != nil {
		panic(err.Error())
	}

	for _, podItem := range pods.Items {
		podIsReady := false
		for _, condition := range podItem.Status.Conditions {
			if condition.Type == v1.PodReady &&
				condition.Status == v1.ConditionTrue {
				podIsReady = true
			}

		}
		if !podIsReady {
			log.Info("pod ", podItem.Name, " on node:", podItem.Status.HostIP, " is not ready...")
			cnt++
		}
	}

	if cnt >= maxInitializingPods {
		log.Warn("there are ", cnt, " pods are initializing on node ", node.Name, ", stop schedule...")
		return false
	}
	return true
}
