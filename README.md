# kube-scheduler-extender

> 参照 https://github.com/everpeace/k8s-scheduler-extender-example.git

### go build

```bash
VERSION=0.0.1
GO111MODULE=on go mod download
go build -ldflags "-s -w -X main.version=$VERSION" kube-scheduler-extender
```

### docker build

```
docker build . -t kube-scheduler-extender:0.0.1
```

### schedule test pod

you will see `test-pod` will be scheduled by `my-scheduler`.

```
$ kubectl create -f test-pod.yaml

$ kubectl describe pod test-pod
Name:         test-pod
...
Events:
  Type    Reason                 Age   From               Message
  ----    ------                 ----  ----               -------
  Normal  Scheduled              25s   my-scheduler       Successfully assigned test-pod to minikube
  Normal  SuccessfulMountVolume  25s   kubelet, minikube  MountVolume.SetUp succeeded for volume "default-token-wrk5s"
  Normal  Pulling                24s   kubelet, minikube  pulling image "nginx"
  Normal  Pulled                 8s    kubelet, minikube  Successfully pulled image "nginx"
  Normal  Created                8s    kubelet, minikube  Created container
  Normal  Started                8s    kubelet, minikube  Started container
```
