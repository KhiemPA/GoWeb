package main

import (
	"fmt"
	"log"
	"net/http"
	"prj/routes"
	"prj/utils"
	"prj/models"
	"prj/sessions"
)

const PORT = ":8080"

func main() {
	models.TestConnection()
	fmt.Println("Start on 8080")
	utils.LoadTemplates("views/*.html")
	sessions.SessionOptions("localhost", "/", 3600, true)
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

