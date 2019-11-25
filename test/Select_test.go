package test

import (
	"Vorm/Basic"
	"fmt"
	"testing"
)

func TestBasic_Select(t *testing.T)  {
	db, err := Basic.Connect("root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		var user User
		if err := db.Where("Name = '%s'", "yy").Select(&user); err != nil {
			panic("fuck again")
		} else {
			fmt.Println(user.Age)
		}
	}
}
