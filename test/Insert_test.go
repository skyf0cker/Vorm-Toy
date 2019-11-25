package test

import (
	"Vorm/Basic"
	"testing"
)

func TestBasic_Insert(t *testing.T)  {
	db, err := Basic.Connect("root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		if err := db.Insert(&User{
			Name: "yy",
			Age:  20,
		}); err != nil {
			panic("fuck again")
		}
	}
}
