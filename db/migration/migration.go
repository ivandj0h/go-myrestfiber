package migration

import (
	"fmt"
	"log"

	"ivandjoh.online/m/v2/db"
	"ivandjoh.online/m/v2/model/entity"
)

func RunMigration() {
	err := db.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Migration complete!")
}
