package main

import (
	"fmt"
	"time"

	"github.com/lib/pq"
)

func getBook(bookID int) (Customer, error) {
	//Retrieve
	res := Customer{}

	var id int
	var name string
	var address string
	var tel int
	var registerDate pq.NullTime

	err := db.QueryRow(`SELECT id, name, address, tel, registerdate FROM books where id = $1`, bookID).Scan(&id, &name, &address, &tel, &registerDate)
	if err == nil {
		res = Customer{ID: id, Name: name, Address: address, Tel: tel, RegisterDate: registerDate.Time}
	}

	return res, err
}

func allBooks() ([]Customer, error) {
	//Retrieve
	books := []Customer{}

	rows, err := db.Query(`SELECT id, name, address, tel, registerdate FROM customers order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var address string
		var tel int
		var registerDate pq.NullTime

		err = rows.Scan(&id, &name, &address, &tel, &registerDate)
		if err != nil {
			return books, err
		}

		currentBook := Customer{ID: id, Name: name, Address: address, Tel: tel}
		if registerDate.Valid {
			currentBook.RegisterDate = registerDate.Time
		}

		books = append(books, currentBook)
	}

	return books, err
}

func insertBook(name, address string, tel int, registerDate time.Time) (int, error) {
	//Create
	var bookID int
	err := db.QueryRow(`INSERT INTO customers(name, address, tel, registerdate) VALUES($1, $2, $3, $4) RETURNING id`, name, address, tel, registerDate).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", bookID)
	return bookID, err
}

func updateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE books set name=$1, adress=$2, tel=$3, registerdate=$4 where id=$5 RETURNING id`, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from customers where id = $1`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
