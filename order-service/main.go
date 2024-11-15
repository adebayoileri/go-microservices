package main

import (
	"log"
    "net/http"
    "go-ecommerce/order-service/handler"
    "go-ecommerce/order-service/repository"
    "github.com/julienschmidt/httprouter"
    "github.com/streadway/amqp"
	"github.com/gocql/gocql"
)

func initCassandra() *gocql.Session {
    cluster := gocql.NewCluster("cassandra")
    cluster.Keyspace = "ecommerce"
    cluster.Consistency = gocql.Quorum
    cluster.ProtoVersion = 4
    
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatalf("Failed to connect to Cassandra: %v", err)
    }
    
    return session
}

func main() {
    rabbitMQConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer rabbitMQConn.Close()

	session := initCassandra()
    defer session.Close()

    orderRepo := repository.NewOrderRepository(session)
    orderHandler := handler.NewOrderHandler(orderRepo, rabbitMQConn)

    router := httprouter.New()
    router.POST("/orders", orderHandler.CreateOrder)

    log.Println("Order Service is running on port 8003")
    log.Fatal(http.ListenAndServe(":8003", router))
}
