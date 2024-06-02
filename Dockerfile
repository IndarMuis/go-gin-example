FROM golang:1.18-alpine AS builder
WORKDIR /build
COPY go.mod go.sum .env ./

RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o go-example .

FROM scratch
COPY --from=builder ["/build/go-example", "/"]
EXPOSE 8888
ENTRYPOINT ["/go-example"]