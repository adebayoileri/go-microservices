package repository

import (
    "errors"
    "go-ecommerce/order-service/model"
    "github.com/gocql/gocql"
)

type OrderRepository struct {
    session *gocql.Session
}

func NewOrderRepository(session *gocql.Session) *OrderRepository {
    return &OrderRepository{session: session}
}

func (r *OrderRepository) GetProductByID(productID int) (model.Product, error) {
    var product model.Product
    query := "SELECT id, name, description, price, stock FROM products WHERE id = ?"
    if err := r.session.Query(query, productID).Scan(
        &product.ID,
        &product.Name,
        &product.Description,
        &product.Price,
        &product.Stock,
    ); err != nil {
        return product, err
    }
    return product, nil
}

func (r *OrderRepository) CreateOrder(order model.Order) error {
    if order.ID == (gocql.UUID{}) {
        order.ID = gocql.TimeUUID()
    }
    
    query := `INSERT INTO orders (id, product_id, quantity, total_price, status, created_at) 
              VALUES (?, ?, ?, ?, ?, ?)`
    
    return r.session.Query(query,
        order.ID,
        order.ProductID,
        order.Quantity,
        float64(order.TotalPrice),
        order.Status,
        order.CreatedAt,
    ).Exec()
}

func (r *OrderRepository) UpdateProductStock(productID int, quantity int) error {
    var currentStock int
    
    if err := r.session.Query("SELECT stock FROM products WHERE id = ?", productID).Scan(&currentStock); err != nil {
        return err
    }
    
    if currentStock < quantity {
        return errors.New("insufficient stock")
    }
    
    applied, err := r.session.Query(
        "UPDATE products SET stock = ? WHERE id = ? IF stock = ?",
        currentStock - quantity,
        productID,
        currentStock,
    ).ScanCAS(nil)
    
    if err != nil {
        return err
    }
    
    if !applied {
        return errors.New("concurrent stock update detected, please retry")
    }
    
    return nil
}