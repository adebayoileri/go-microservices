package model

type Product struct {
    ID          int     `json:"id" cql:"id"`
    Name        string  `json:"name" cql:"name"`
    Description string  `json:"description" cql:"description"`
    Price       float64 `json:"price" cql:"price"`
    Stock       int     `json:"stock" cql:"stock"`
}
