FROM golang:1.23.3

WORKDIR /app

# Install netcat
RUN apt-get update && apt-get install -y netcat-openbsd

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Convert CRLF to LF for the script
RUN sed -i 's/\r$//' wait-for-it.sh
RUN chmod +x wait-for-it.sh

RUN go build -o order-service .

EXPOSE 8003

CMD ["./wait-for-it.sh", "cassandra:9042", "--", "./order-service"]
