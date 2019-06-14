.PHONY: clean

build:
	CGO_ENABLED=0 GOOS=linux go build -o kube-operator

image: build
	docker build -t kube-operator .

push: image
docker push kube-operators:latest

clean:
	rm -rf kube-operator
