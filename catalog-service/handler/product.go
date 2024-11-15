package handler

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "go-ecommerce/catalog-service/model"
    "go-ecommerce/catalog-service/repository"
    "github.com/julienschmidt/httprouter"
    "github.com/gocql/gocql"
)

type ProductHandler struct {
    repo *repository.ProductRepository
}

func NewProductHandler(session *gocql.Session) *ProductHandler {
    return &ProductHandler{
        repo: repository.NewProductRepository(session),
    }
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    products, err := h.repo.GetAll()
    log.Println(products)
    if err != nil {
        log.Println(err)
        http.Error(w, "Error retrieving products", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.Atoi(ps.ByName("id"))
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
    product, err := h.repo.GetByID(id)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var product model.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    if err := h.repo.Create(product); err != nil {
        log.Println(err)
        http.Error(w, "Error creating product", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}
