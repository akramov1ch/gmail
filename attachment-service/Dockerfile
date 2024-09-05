FROM golang:1.23.0-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o attachment-service .

EXPOSE 9003

CMD ["./attachment-service"]
