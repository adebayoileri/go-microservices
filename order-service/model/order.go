package model

import (
    "github.com/gocql/gocql"
    "time"
)

type Order struct {
    ID          gocql.UUID `json:"id" cql:"id"`
    ProductID   int        `json:"product_id" cql:"product_id"`
    Quantity    int        `json:"quantity" cql:"quantity"`
    TotalPrice  float64    `json:"total_price" cql:"total_price"`
    Status      string     `json:"status" cql:"status"`
    CreatedAt   time.Time  `json:"created_at" cql:"created_at"`
}

type Product struct {
    ID          int     `json:"id" cql:"id"`
    Name        string  `json:"name" cql:"name"`
    Description string  `json:"description" cql:"description"`
    Price       float64 `json:"price" cql:"price"`
    Stock       int     `json:"stock" cql:"stock"`
}
