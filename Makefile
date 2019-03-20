.PHONY: clean

build:
	CGO_ENABLED=0 GOOS=linux go build -o kube-operators

image: build
	docker build -t kube-operators .

push: image
docker push kube-operators:latest

clean:
	rm -rf kube-operators