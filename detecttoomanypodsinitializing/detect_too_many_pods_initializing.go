package detecttoomanypodsinitializing

import (
	"github.com/prometheus/common/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-scheduler-extender/kubeclientset"
	"strings"
)

func DetectTooManyPodsInitializing(node v1.Node, maxInitializingPods int) (res bool) {
	clientset := kubeclientset.Client
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	cnt := 0

	for _, item := range namespaces.Items {
		namespace := item.Name
		if strings.HasPrefix(namespace, "kube-") {
			continue
		}
		pods, _ := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{FieldSelector: "spec.nodeName=" + node.Name})

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
	}

	if cnt >= maxInitializingPods {
		log.Warn("there are ", cnt, " pods are initializing on node ", node.Name, ", stop schedule...")
		return false
	}
	return true
}
