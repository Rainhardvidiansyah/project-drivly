package main

import (
	"sewa-supir/config/database"
)

func main() {

	cfg := database.LoadConfig()

	database.InitializeDatabase(cfg)

}
