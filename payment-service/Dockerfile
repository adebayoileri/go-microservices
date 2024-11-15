FROM golang:1.23.3

WORKDIR /app

COPY go.mod .
COPY go.sum . 

RUN go mod download

COPY . .

RUN go build -o payment-service .

EXPOSE 8004

CMD ["./payment-service"]
