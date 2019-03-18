FROM golang:alpine as builder
RUN apk add --no-cache git
COPY . /go/src/github.com/pivotal-cf/azure-blobstore-resource
ENV CGO_ENABLED 0
RUN go build -o /assets/in github.com/pivotal-cf/azure-blobstore-resource/cmd/in
RUN go build -o /assets/out github.com/pivotal-cf/azure-blobstore-resource/cmd/out
RUN go build -o /assets/check github.com/pivotal-cf/azure-blobstore-resource/cmd/check
WORKDIR /go/src/github.com/pivotal-cf/azure-blobstore-resource
RUN go get github.com/onsi/ginkgo/ginkgo

ARG TEST_STORAGE_ACCOUNT_NAME
ENV TEST_STORAGE_ACCOUNT_NAME=${TEST_STORAGE_ACCOUNT_NAME}

ARG TEST_STORAGE_ACCOUNT_KEY
ENV TEST_STORAGE_ACCOUNT_KEY=${TEST_STORAGE_ACCOUNT_KEY}

RUN ginkgo -r /go/src/github.com/pivotal-cf/azure-blobstore-resource

FROM alpine:edge AS resource
RUN apk add --no-cache bash tzdata ca-certificates unzip zip gzip tar
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource
