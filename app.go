package main

import (
	"fmt"
	"log"
	"net/http"
	"prj/routes"
	"prj/utils"
	"prj/models"
)

const PORT = ":8080"

func main() {
	models.TestConnection()
	fmt.Println(models.GetProducts())
	fmt.Println(models.GetCategories())
	fmt.Println("Start on 8080")
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

