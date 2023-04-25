FROM ghcr.io/roadrunner-server/velox:latest AS velox

FROM golang:1.20 AS golang

# Personal access tokens fot github and gitlab respectively
ARG RT_TOKEN
ARG GL_TOKEN

ENV CGO_ENABLED=0

COPY --from=velox /usr/bin/vx /usr/local/bin/vx

WORKDIR /go/src
COPY config.go .
COPY interfaces.go .
COPY plugin.go .
COPY rpc.go .
COPY go.mod .
COPY go.sum .
COPY ./velox.toml .

RUN go get .

WORKDIR /go/bin
COPY .rr.yaml .
RUN /usr/local/bin/vx build -c /go/src/velox.toml
RUN /go/bin/rr --version

CMD ["/go/bin/rr"]