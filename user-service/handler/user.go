package handler

import (
    "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("User registered successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("User logged in successfully"))
}
