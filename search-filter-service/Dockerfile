FROM golang:1.23.0-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o search-filter-service .

EXPOSE 9004

CMD ["./search-filter-service"]
