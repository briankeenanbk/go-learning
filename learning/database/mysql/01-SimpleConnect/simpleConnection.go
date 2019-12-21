package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	db      *sql.DB
	insert  *sql.Rows
	results *sql.Rows
}
type data struct {
	f1 string
	f2 string
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	var con database
	var err error

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	con.db, err = sql.Open("mysql", "brian:_Agadub52!@tcp(127.0.0.1:3306)/testschema")
	fmt.Printf("db =%T\n", con.db)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer con.db.Close()

	con.insert, err = con.db.Query("INSERT INTO testtable VALUES ( 'TEST9', 'TEST' )")
	fmt.Printf("insert =%T\n", con.insert)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer con.insert.Close()

	// Execute the query
	con.results, err = con.db.Query("SELECT * FROM testtable WHERE test_one = ?", "TEST6")
	fmt.Printf("results =%T\n", con.results)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for con.results.Next() {
		var tag data
		// for each row, scan the result into our tag composite object
		err = con.results.Scan(&tag.f1, &tag.f2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tag)
	}

}
