package database

import (
	"database/sql"
)

func InitDB() *sql.DB {
	connectionString := "root:12345678@tcp(10.211.55.4:3306)/northwind"

	databaseConnection, err := sql.Open("mysql", connectionString)

	println(databaseConnection)
	println(err == nil)
	if err != nil {

		panic(err.Error()) //Erro Handling = manejo de errores
	}

	return databaseConnection
}
