package main

import (
    "log"
    "net/http"
    "go-ecommerce/user-service/handler"
)

func main() {
    http.HandleFunc("/register", handler.Register)
    http.HandleFunc("/login", handler.Login)
    log.Fatal(http.ListenAndServe(":8001", nil))
}
