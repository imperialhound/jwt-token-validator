FROM golang:1.21.4 as builder
WORKDIR /src
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal/ ./internal/
COPY cmd/server/ ./cmd/server/

RUN go build -o server cmd/server/main.go

FROM alpine:3.18
COPY --from=builder /src/server /server
ENTRYPOINT ["/server"]
