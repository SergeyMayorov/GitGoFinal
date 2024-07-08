FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/main.go

EXPOSE 80

CMD ["./main"]

