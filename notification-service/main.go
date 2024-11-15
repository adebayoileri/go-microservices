package main

import (
    "log"
    "github.com/streadway/amqp"
)

func main() {
    consumePaymentStatus()
}

func consumePaymentStatus() {
    conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatal(err)
    }
    defer ch.Close()

       // Declare the queue (must match the consumer’s queue configuration)
       q, err := ch.QueueDeclare(
        "payment_queue", // Queue name
        true,          // Durable
        false,         // Delete when unused
        false,         // Exclusive
        false,         // No-wait
        nil,           // Arguments
    )
   
    log.Printf("Queue %v", q)

    msgs, err := ch.Consume(
        "payment_queue",
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for d := range msgs {
            log.Printf("Received payment status: %s", d.Body)

            log.Printf("Sending notification for: %s", d.Body)
        }
    }()
    log.Printf("Waiting for payment status messages.")
    <-make(chan bool)
}
