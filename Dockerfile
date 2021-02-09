# Build the manager binary
FROM golang:1.15 as builder

WORKDIR /workspace
COPY go.sum go.sum
COPY go.mod go.mod
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o wait-for-http main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /workspace/wait-for-http /
ENTRYPOINT [ "/wait-for-http" ]
