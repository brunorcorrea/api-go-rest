package main

import (
	"fmt"

	"github.com/brunorcorrea/go-rest-api/database"
	"github.com/brunorcorrea/go-rest-api/routes"
	"github.com/brunorcorrea/go-rest-api/routes/models"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
		{Id: 3, Nome: "Nome 3", Historia: "Historia 3"},
		{Id: 4, Nome: "Nome 4", Historia: "Historia 4"},
	}
	
	database.ConectaBanco()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}