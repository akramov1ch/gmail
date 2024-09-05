FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o email-service .

EXPOSE 9002

CMD ["./email-service"]
