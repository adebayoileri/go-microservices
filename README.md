# Ecommerce Microservice

Event-driven microservice system using Go, RabbitMQ, and Cassandra

## Services
- User Service
- Catalog Service
- Order Service
- Payment Service
- Notification Service

## Technologies

- Go
- RabbitMQ
- Cassandra

## How to run

1. Install Docker
2. Run `docker-compose up --build`

## Test

1. Run `curl -X POST http://localhost:8002/products -H "Content-Type: application/json" -d '{"name": "Product 1", "description": "Description 1", "price": 100, "stock": 10}'`
2. Run `curl -X POST http://localhost:8003/orders -H "Content-Type: application/json" -d '{"product_id": 1, "quantity": 1, "total_price": 100}'`
3. Check the payment service logs to see the payment status
4. Check the notification service logs to see the notification status

## Notes
- The payment service is mocked to always return a successful payment
- The notification service is mocked to always send a notification
