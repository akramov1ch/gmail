FROM golang:1.23.0-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o notification-service .

EXPOSE 9006

CMD ["./notification-service"]
