FROM golang:1.15 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN mkdir store
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' ./cmd/golang-rest-api-example

FROM scratch

WORKDIR /app
ENTRYPOINT ["/app/golang-rest-api-example"]

COPY --from=builder /app/golang-rest-api-example /app/golang-rest-api-example
COPY --from=builder /app/store /app/store