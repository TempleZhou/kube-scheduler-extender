package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/common/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	detector "kube-scheduler-extender/detecttoomanypodsinitializing"
	"kube-scheduler-extender/kubeclientset"
	"net/http"

	"k8s.io/api/core/v1"
)

const (
	versionPath         = "/version"
	apiPrefix           = "/scheduler"
	predicatesPrefix    = apiPrefix + "/predicates"
	maxInitializingPods = 2
)

var (
	version string // injected via ldflags at build time

	TooManyPodsInitializingPredicate = Predicate{
		Name: "detect_too_many_pods_initializing",
		Func: func(pod v1.Pod, node v1.Node) (bool, error) {
			return detector.DetectTooManyPodsInitializing(node, maxInitializingPods), nil
		},
	}
)

func main() {
	// 初始化 k8 client
	var kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	kubeclientset.Client, _ = kubernetes.NewForConfig(config)

	// 初始化 http server
	router := httprouter.New()
	AddVersion(router)

	predicates := []Predicate{TooManyPodsInitializingPredicate}
	for _, p := range predicates {
		AddPredicate(router, p)
	}

	log.Info("server starting on the port :80")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}
