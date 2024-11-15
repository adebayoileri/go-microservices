package handler

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "go-ecommerce/order-service/model"
    "go-ecommerce/order-service/repository"

    "github.com/julienschmidt/httprouter"
    "github.com/streadway/amqp"
    "github.com/gocql/gocql"
)

type OrderHandler struct {
    repo           *repository.OrderRepository
    rabbitMQClient *amqp.Connection
}

func NewOrderHandler(repo *repository.OrderRepository, rabbitMQClient *amqp.Connection) *OrderHandler {
    return &OrderHandler{repo: repo, rabbitMQClient: rabbitMQClient}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var req struct {
        ProductID int `json:"product_id"`
        Quantity  int `json:"quantity"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    product, err := h.repo.GetProductByID(req.ProductID)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    if product.Stock < req.Quantity {
        http.Error(w, "Insufficient stock", http.StatusConflict)
        return
    }

    totalPrice := float64(product.Price * float64(req.Quantity))

    order := model.Order{
        ID:         gocql.TimeUUID(),
        ProductID:  req.ProductID,
        Quantity:   req.Quantity,
        TotalPrice: totalPrice,
        Status:     "pending",
        CreatedAt:  time.Now(),
    }

    if err := h.repo.CreateOrder(order); err != nil {
        log.Printf("Error creating order: %v", err)    
        http.Error(w, "Error creating order", http.StatusInternalServerError)
        return
    }

    if err := h.repo.UpdateProductStock(req.ProductID, req.Quantity); err != nil {
        log.Printf("Error updating stock: %v", err)
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }

    if err := h.publishOrderMessage(order); err != nil {
        log.Printf("Failed to publish order message: %v", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message":     "Order created successfully",
        "order_id":    order.ID.String(),
        "total_price": totalPrice,
    })
    log.Printf("Order created and message published for product ID %d with quantity %d", req.ProductID, req.Quantity)
}

func (h *OrderHandler) publishOrderMessage(order model.Order) error {
    ch, err := h.rabbitMQClient.Channel()
    if err != nil {
        return err
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
    if err != nil {
        return err
    }

    body, err := json.Marshal(order)
    if err != nil {
        return err
    }

    err = ch.Publish(
        "",           // Exchange
        q.Name,       // Routing key (queue name)
        false,        // Mandatory
        false,        // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    return err
}
