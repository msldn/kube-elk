FROM golang:1.10-alpine AS build

WORKDIR /go/src/github.com/marek5050/kube-elk
COPY . .

RUN apk --no-cache add -t git \
	&& apk --no-cache add ca-certificates
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main ./cmd/elk



FROM golang:1.10-alpine
COPY --from=build /go/src/github.com/marek5050/kube-elk/main /usr/local/bin/
COPY --from=build /go/src/github.com/marek5050/kube-elk/cfg/.kube /root/.kube/