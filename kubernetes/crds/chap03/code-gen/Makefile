gen:
	export GO111MODULE=off
	go get k8s.io/code-generator
	go get k8s.io/apimachinery
	${GOPATH}/src/k8s.io/code-generator/generate-groups.sh all \
    github.com/liuliqiang/blog-demos/kubernetes/crds/chap03/code-gen/client \
    github.com/liuliqiang/blog-demos/kubernetes/crds/chap03/code-gen/apis \
    admin:v1 \
