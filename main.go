package main

import (
	"assignment/routes"
	"assignment/services"
)

func main() {
	db, err := services.Connect("root:root@(127.0.0.1:3306)/assignment_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	g := routes.Create(db)
	g.Run("127.0.0.1:3000")
}
