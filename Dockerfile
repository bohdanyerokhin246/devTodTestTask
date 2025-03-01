FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN chmod +x ./migrate.sh

RUN go build -o myapp .


CMD ["./myapp"]
