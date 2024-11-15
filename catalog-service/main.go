package main

import (
    "log"
    "net/http"
    "go-ecommerce/catalog-service/handler"
    "github.com/julienschmidt/httprouter"
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
    session := initCassandra()
    defer session.Close()

    router := httprouter.New()
    productHandler := handler.NewProductHandler(session)

    router.GET("/products", productHandler.GetProducts)
    router.GET("/products/:id", productHandler.GetProductByID)
    router.POST("/products", productHandler.CreateProduct)

    log.Println("Catalog Service is running on port 8002")
    log.Fatal(http.ListenAndServe(":8002", router))
}