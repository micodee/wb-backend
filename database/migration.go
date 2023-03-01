package database

import (
	"backend/models"
	"backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.ConnDB.AutoMigrate(&models.User{}, &models.Product{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}