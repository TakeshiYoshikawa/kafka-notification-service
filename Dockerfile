FROM golang:1.23-alpine

WORKDIR /docker-app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "producer_client.go"]

EXPOSE 8080