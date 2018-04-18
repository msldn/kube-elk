FROM golang:1.10-alpine


WORKDIR /go/src/github.com/marek5050/kube-elk
COPY . .

RUN apk --no-cache add -t build-deps build-base git \
	&& apk --no-cache add ca-certificates
RUN go get -d -v ./...
RUN go install -v ./...