# kube-scheduler-extender

> 参照 https://github.com/everpeace/k8s-scheduler-extender-example.git  
> 解决 https://github.com/kubernetes/kubernetes/issues/3312

### go build

```bash
VERSION=0.0.1
GO111MODULE=on go mod download
go build -ldflags "-s -w -X main.version=$VERSION" kube-scheduler-extender
```

### docker build

```bash
docker build . -t kube-scheduler-extender:0.0.1
```

### run kube-scheduler-extender

修改默认的 kube-scheduler.yaml，加入 kube-scheduler-extender Container。

如果使用 kubeadm 初始化，kube-scheduler.yaml 在 /etc/kubernetes/manifests 路径下，将自定义的 kube-scheduler.yaml 覆盖过去即可。

> 注意： 需要修改 kube-scheduler.yaml 中的镜像为实际使用的镜像

在 kube-schedule Pod 中加入 kube-scheduler-extender Container 之后，可以看到如下状态：

### 创建测试 pod

```bash
cd kubernetes

docker build . -t hello-node
kubectl apply -f hello-node.yaml
```

将会看到 7 个 Pod 分 4 批启动，每次仅启动 2 个，待 Ready 后继续启动后续 Pod。

```bash
hello-node-6d686f5784-4ksmz   0/1     Pending     0          7s
hello-node-6d686f5784-55w27   0/1     Pending     0          7s
hello-node-6d686f5784-8w2jz   0/1     Pending     0          7s
hello-node-6d686f5784-ckj7t   0/1     Pending     0          7s
hello-node-6d686f5784-fphjb   0/1     Pending     0          7s
hello-node-6d686f5784-frkf9   0/1     Running     0          7s
hello-node-6d686f5784-nppgx   0/1     Running     0          7s
```
