package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ttdennis/kontainer.io/user"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = user.NewService(db)
	if err != nil {
		panic(err)
	}
}
