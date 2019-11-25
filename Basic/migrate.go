package Basic

import (
	"Vorm/utils"
	"fmt"
	"reflect"
	"strings"
)

func (d *Db) AutoMigrate(st interface{}) error {
	var keys, values []string
	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)

	name := t.String()
	spl_list := strings.Split(name, ".")
	name = spl_list[len(spl_list)-1]

	for v.Kind() == reflect.Ptr{
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		keys, values = utils.GetTableStruct(v)
	default:
		panic("not implement")
	}

	var strList []string
	for index := 0;index < len(keys); index++{
		str := keys[index] + " " + values[index] + ","
		strList = append(strList, str)
	}

	payload := "(" + strings.TrimRight(strings.Join(strList, ""), ",") + ")"
	fmt.Println(payload)
	sql_sen := fmt.Sprintf("create table %s %s", name, payload)
	fmt.Println(sql_sen)
	result, err := d.DB.Exec(sql_sen)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
	return err
}
