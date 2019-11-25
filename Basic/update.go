package Basic

import (
	"Vorm/utils"
	"fmt"
	"reflect"
	"strings"
)

func (d *Db)Update(inf interface{}) error {
	var where string
	var keys, values []string
	if d.where != ""{
		where = d.where
	}
	t := reflect.TypeOf(inf)
	v := reflect.ValueOf(inf)

	if t.Kind() == reflect.Ptr{
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct{
		keys, values = utils.SKV(v)
	} else {
		panic("not implement")
	}

	splList := strings.Split(t.String(), ".")
	name := splList[len(splList)-1]

	var str string
	for i:=0;i<len(keys);i++ {
		str += keys[i] + "=" + values[i] + " , "
	}
	str = strings.TrimRight(str, " , ")
	sql_sen := fmt.Sprintf("UPDATE %s SET %s WHERE %s", name, str, where)
	result, err := d.DB.Exec(sql_sen)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.RowsAffected())
	}
	return err
}