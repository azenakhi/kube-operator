FROM golang:latest as build
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR $GOPATH/src/github.com/azenakhi/kube-operator
COPY . .
RUN dep ensure
RUN make build 

FROM alpine
WORKDIR /root
COPY --from=build /go/src/github.com/azenakhi/kube-operator/build/kube-operator .
CMD ["./kube-operator"]


