package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"toko/controller"
	"toko/query"
	"toko/usecase"

	_ "github.com/lib/pq" // here
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "pengguna", "pengguna", "pengguna")

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productQuery := query.NewDBQuery(db)

	uc := usecase.NewUsecase(productQuery)

	c := controller.NewControllerHttp(uc)

	http.HandleFunc("/buy", c.Buy)

	log.Fatal(http.ListenAndServe(":10000", nil))
}
