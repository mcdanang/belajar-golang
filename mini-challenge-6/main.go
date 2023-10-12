package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// USING MYSQL
const (
	user     = "root"
	password = ""
	dbname   = "mini-challenge-6"
)

var (
	db   *sql.DB
	err  error
	rows *sql.Rows
)

type Product struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
	Variants  []Variant
}

type Variant struct {
	ID          int
	VariantName string
	Quantity    int
	ProductID   int
	CreatedAt   string
	UpdatedAt   string
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	createProduct("Jaket")
	// updateProduct("Celana Pendek", 3)
	getProductById(1)
	// updateVariantById("Ungu", 35, 5, 2)
	// createVariant("Putih", 5, 3)
	// getProductWithVariant(3)
	// deleteVariantById(4)
}

func createProduct(name string) {
	var product = Product{}

	sqlStatement := `
	INSERT INTO products (name)
	VALUES (?)
	`

	result, err := db.Exec(sqlStatement, name)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Retrieve the inserted row using the lastInsertID
	sqlRetrieve := `
    SELECT * FROM products WHERE id = ?
	`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Product Data: %+v\n", product)
}

func updateProduct(name string, id int) {
	sqlStatement := `UPDATE products SET name = ? WHERE id = ?;`
	res, err := db.Exec(sqlStatement, name, id)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}

func getProductById(id int) {
	var product = Product{}
	sqlRetrieve := `
    SELECT * FROM products WHERE id = ?
	`

	err = db.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get Product ID %d: %+v\n", id, product)
}

func updateVariantById(name string, qty int, productId int, id int) {
	sqlStatement := `UPDATE variants SET variant_name = ?, quantity = ?, product_id = ? WHERE id = ?;`
	res, err := db.Exec(sqlStatement, name, qty, productId, id)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}

func createVariant(name string, qty int, productId int) {
	var variant = Variant{}

	sqlStatement := `
	INSERT INTO variants (variant_name, quantity, product_id)
	VALUES (?, ?, ?)
	`

	result, err := db.Exec(sqlStatement, name, qty, productId)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Retrieve the inserted row using the lastInsertID
	sqlRetrieve := `
    SELECT * FROM variants WHERE id = ?
	`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Variant Data: %+v\n", variant)
}

func getProductWithVariant(id int) {
	var product = Product{}
	sqlRetrieve := `
    SELECT * FROM products WHERE id = ?
	`

	err = db.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	var variants = []Variant{}

	sqlStatement := `
    SELECT * FROM variants WHERE product_id = ?
	`

	rows, err = db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var variant = Variant{}

		err = rows.Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			panic(err)
		}

		variants = append(variants, variant)
	}

	product.Variants = variants

	fmt.Printf("Get Product ID %d with Variant: %+v\n", id, product)
}

func deleteVariantById(id int) {
	sqlStatement := `DELETE from variants WHERE id = ?;`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount:", count)
}
