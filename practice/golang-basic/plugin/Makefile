docker_image = golang:1.13

plugin:
	docker run -v $(shell pwd):/go/src/  $(docker_image) go build -buildmode=plugin -o /go/src/plugin.so /go/src/plugin.go
