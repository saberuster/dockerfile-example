FROM golang:1.11.0-alpine3.8  as builder
WORKDIR /go/project/example
COPY .  .
RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags "-s -w " -a -installsuffix cgo -o ./example main.go
FROM alpine:3.8
WORKDIR /go/project/example
COPY --from=builder /go/project/example/example .
COPY --from=builder /go/project/example/file .
ENTRYPOINT ["./example"]


