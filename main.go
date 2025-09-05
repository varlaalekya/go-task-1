package main

import (
    "database/sql"
    "log"
    "net/http"

     //"github.com/go-sql-driver/mysql"
	   "go-task-1/api"
)
	

func main() {
	dsn := "username:password@tcp(localhost(127.0.0.1)dbname?)parseTime=true"
	db, err :=sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.RegisterRoutes(db)

	//start the HTTP server
	log.Println("Listening on  : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}



