.PHONY: build

REPO  = docker.io/azenakhi
IMAGE = kube-operator
TAG   = latest

build:
	@go build -o ./build/kube-operator -i cmd/main.go

login:
	@docker login

image: 
	@docker build -t ${REPO}/${IMAGE}:${TAG} .

push:
	@docker push ${REPO}/${IMAGE}:${TAG}

clean:
	@rm -rf build/kube-operator
