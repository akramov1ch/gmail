FROM golang:1.23.0-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o email-delivery-service .

EXPOSE 9007

CMD ["./email-delivery-service"]
