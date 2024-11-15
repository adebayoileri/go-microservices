package repository

import (
    "go-ecommerce/catalog-service/model"
    "github.com/gocql/gocql"
)

type ProductRepository struct {
    session *gocql.Session
}

func NewProductRepository(session *gocql.Session) *ProductRepository {
    return &ProductRepository{session: session}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
    query := "SELECT id, name, description, price, stock FROM products"
    scanner := r.session.Query(query).Iter().Scanner()
    
    var products []model.Product
    for scanner.Next() {
        var product model.Product
        
        if err := scanner.Scan(
            &product.ID,
            &product.Name,
            &product.Description,
            &product.Price,
            &product.Stock,
        ); err != nil {
            return nil, err
        }
        
        products = append(products, product)
    }
    
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    
    return products, nil
}

func (r *ProductRepository) GetByID(id int) (model.Product, error) {
    var product model.Product
    query := "SELECT id, name, description, price, stock FROM products WHERE id = ?"
    err := r.session.Query(query, id).Scan(
        &product.ID,
        &product.Name,
        &product.Description,
        &product.Price,
        &product.Stock,
    )
    return product, err
}

func (r *ProductRepository) Create(product model.Product) error {
    query := `INSERT INTO products (id, name, description, price, stock) 
              VALUES (?, ?, ?, ?, ?)`
    
    return r.session.Query(query,
        product.ID,
        product.Name,
        product.Description,
        float64(product.Price),
        product.Stock,
    ).Exec()
}
