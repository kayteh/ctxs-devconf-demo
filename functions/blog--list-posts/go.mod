module function/handler

go 1.13

require (
	cloud.google.com/go/firestore v1.1.1
	github.com/golang/protobuf v1.3.3
	github.com/kayteh/ctxs-devconf-demo v0.0.0-20200203224928-7b700b5f969f
	github.com/kayteh/ctxs-devconf-rpc v1.5.100
	google.golang.org/grpc v1.27.0
	k8s.io/klog v1.0.0
)

replace github.com/kayteh/ctxs-devconf-demo => ../../
