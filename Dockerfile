#FROM golang:1.12-alpine as builder
#ARG VERSION=0.0.1
#
#ENV GO111MODULE=on
#ENV CGO_ENABLED=0
#ENV GOOS=linux
#ENV GOARCH=amd64
#
#WORKDIR /go/src/kube-scheduler-extender
#COPY go.mod .
#COPY go.sum .
#RUN GO111MODULE=on go mod download
#COPY . .
#RUN go install -ldflags "-s -w -X main.version=$VERSION" kube-scheduler-extender
#
#FROM gcr.io/google_containers/ubuntu-slim:0.14
#COPY --from=builder /go/bin/kube-scheduler-extender /usr/bin/kube-scheduler-extender
#ENTRYPOINT ["kube-scheduler-extender"]

FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY kube-scheduler-extender /usr/bin/kube-scheduler-extender
ENTRYPOINT ["kube-scheduler-extender"]
