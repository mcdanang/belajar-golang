package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	password = ""
	dbname   = "go-sql"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID     int
	Title  string
	Author string
	Stock  int
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sucessfully connected to database")

	CreateBook()
	// GetBooks()
	// UpdateBook()
	// DeleteBook()
}

func CreateBook() {
	var book = Book{}

	sqlStatement :=
		`INSERT INTO books (title, author, stock) VALUES (?, ?, ?)`

	fmt.Println("ERRRRORR GAAN")
	result, err := db.Exec(sqlStatement, "Laskar Pelangi", "Andrea Hirata", 30)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// retrieve inserted row
	sqlRetrieve :=
		`SELECT * FROM books WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Printf("New book data: %v\n", book)
}

func GetBooks() {
	var results = []Book{}

	sqlRetrieve := `SELECT * FROM books`

	rows, err := db.Query(sqlRetrieve)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	fmt.Println("Book datas: ", results)
}

func UpdateBook() {
	sqlStatement :=
		`UPDATE books SET title = ?, author = ?, stock = ? WHERE id = ?;`

	result, err := db.Exec(sqlStatement, "Laskar Pelangi Updated", "Andrea Hirata Update", 300, 2)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated book: ", count)
}

func DeleteBook() {
	sqlStatement :=
		`DELETE from books WHERE id = ?;`

	result, err := db.Exec(sqlStatement, "Laskar Pelangi Updated", "Andrea Hirata Update", 300, 2)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated book: ", count)
}
