package main

import (
	"fmt"
	"go-book-management-system/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Started server at port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
