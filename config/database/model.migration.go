package database

import (
	"log"
	"sewa-supir/internal/users"
)

func MigrateMode() {

	models := []interface{}{
		&users.Users{},
	}

	for _, dataModel := range models {

		err := DB.AutoMigrate(dataModel)

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("😎 Database migration successful 😎")

}
