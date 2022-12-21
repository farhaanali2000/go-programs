package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpass = "root"
const dbname = "movies"

func AddMovietoDB(mov Film) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO moviess (name,genre,director) VALUES (?,?,?)",
		mov.Name, mov.Genre, mov.Director)

	/*
		// Or use fmt.Sprintf to concatenate SQL statement if prepared statement isn't worth here
		sqlstm :=
			fmt.Sprintf("INSERT INTO product (code,name,qty,last_updated)"+
				" VALUES ('%s','%s',%d, now())",
				product.Code, product.Name, product.Qty)
		insert, err := db.Query(sqlstm)
	*/

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func GetMovies() []Film {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		// fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM moviess")

	if err != nil {

		// fmt.Println("Err", err.Error())

		return nil

	}

	moviesss := []Film{}

	for results.Next() {

		var fl Film

		err = results.Scan(&fl.Name, &fl.Genre, &fl.Director)

		if err != nil {
			panic(err.Error())
		}

		moviesss = append(moviesss, fl)

		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)
	}

	return moviesss

}
