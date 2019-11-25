package test

import (
	"Vorm/Basic"
	"testing"
)

func TestBasic_Update(t *testing.T)  {
	db, err := Basic.Connect("root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		if err := db.Where("Name = '%s'", "yy").Update(&User{
			Age:  18,
		}); err != nil {
			panic("fuck again")
		}
	}
}
