ARG GO_VER=1.15
FROM golang:${GO_VER}-alpine AS builder

WORKDIR /app
COPY ./go.mod ./main.go ./
RUN CGO_ENABLED=0 go build ./...


FROM scratch
COPY --from=builder /app/go-hello-world /go-hello-world
ENTRYPOINT [ "/go-hello-world" ]
EXPOSE 8080
