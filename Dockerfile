FROM golang:1.21-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /sheriff ./cmd/main.go

EXPOSE 8080

ENTRYPOINT ["/sheriff", "server"]
