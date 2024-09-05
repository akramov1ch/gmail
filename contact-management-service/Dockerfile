FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o contact-management-service .

EXPOSE 9005

CMD ["./contact-management-service"]
