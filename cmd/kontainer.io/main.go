package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ttdennis/kontainer.io/dbwrap"
	"github.com/ttdennis/kontainer.io/user"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbWrapper := dbwrap.NewWrapper(db)

	_, err = user.NewService(dbWrapper)
	if err != nil {
		panic(err)
	}
}
