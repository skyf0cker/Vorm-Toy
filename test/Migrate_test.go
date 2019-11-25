package test

import (
	"Vorm/Basic"
	"testing"
)

type User struct {
	Name string
	Age int
}

func TestConnection_Migrate(t *testing.T) {
	db, err := Basic.Connect("root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		if err := db.AutoMigrate(&User{}); err != nil {
			panic("fuck again")
		}
	}

}
