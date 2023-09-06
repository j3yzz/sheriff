FROM golang:1.21-alpine AS build-stage

WORKDIR /app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /sheriff ./cmd/main.go

EXPOSE 8080

#ENTRYPOINT ["/sheriff", "server"]
CMD ["air"]
