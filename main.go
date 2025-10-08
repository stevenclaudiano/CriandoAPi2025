package main

import (
	"github.com/stevenclaudiano/CriandoAPi2025/database"
	"github.com/stevenclaudiano/CriandoAPi2025/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
