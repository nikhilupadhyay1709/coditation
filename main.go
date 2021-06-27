package main

import (
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
