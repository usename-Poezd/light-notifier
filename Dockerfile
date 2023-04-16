FROM golang:1.19-alpine3.16
WORKDIR /src
RUN set -ex && \
    apk add --no-cache gcc musl-dev
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app ./cmd/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 8080

CMD ["/app"]
