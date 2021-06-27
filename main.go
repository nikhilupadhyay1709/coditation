package main

import (
	//	"github.com/dwahyudi/inventory/configs/database"
	//	"github.com/dwahyudi/inventory/configs/routes"
	db "coditation/database"
	"coditation/routes"
	"log"
)

func main() {
	log.Println("test")
	db.Prepare()
	web := routes.GenerateRoutes()

	log.Fatal(web.Run())
}
