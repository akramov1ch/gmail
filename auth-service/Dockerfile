FROM golang:1.23.0-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o auth-service .

EXPOSE 9001

CMD ["./auth-service"]
