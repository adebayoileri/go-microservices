package main

import (
    "log"
    "github.com/streadway/amqp"
)

func main() {
    consumeOrderMessages()
}

func consumeOrderMessages() {
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

    q, err := ch.QueueDeclare(
        "order_queue", // Queue name
        true,          // Durable
        false,         // Delete when unused
        false,         // Exclusive
        false,         // No-wait
        nil,           // Arguments
    )
   
    log.Printf("Queue %v", q)

    msgs, err := ch.Consume(
        "order_queue",
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
            log.Printf("Received order message: %s", d.Body)

            // Mocked payment processing here
            paymentStatus := "payment.success"

            publishPaymentStatus(paymentStatus)
        }
    }()
    log.Printf("Waiting for messages.")
    <-make(chan bool)
}

func publishPaymentStatus(message string) {
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

    q, err := ch.QueueDeclare(
        "payment_queue",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }

    err = ch.Publish(
        "",         // exchange
        q.Name,     // routing key
        false,      // mandatory
        false,      // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        },
    )
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Sent payment status message: %s", message)
}
