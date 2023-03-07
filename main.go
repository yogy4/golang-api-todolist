package main

import (
	"fmt"
	"golang-api-todolist/config"
	"golang-api-todolist/route"
)

func main() {
	fmt.Println("hello api golang")
	db, _ := config.DBConn()
	route.SetupRouteApi(db)
}
