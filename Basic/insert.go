package Basic

import (
	"Vorm/utils"
	"fmt"
	"reflect"
	"strings"
)

//insert into user (first_name, last_name) values ('Tom', 'Cat'), ('Tom', 'Cruise')

func (d *Db)Insert(inf interface{}) error {
	var keys, values []string
	t := reflect.TypeOf(inf)
	v := reflect.ValueOf(inf)
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		keys, values = utils.SKV(v)
	default:
		panic("not implement")
	}
	splList := strings.Split(t.String(), ".")
	name := splList[len(splList)-1]

	keyStr := "(" + strings.Join(keys, ",") + ")"
	valStr := "(" + strings.Join(values, ",") + ")"

	sql_sen := fmt.Sprintf("insert into %s %s values %s", name, keyStr, valStr)
	fmt.Println(sql_sen)
	result, err := d.DB.Exec(sql_sen)
	if err == nil {
		fmt.Println(result.RowsAffected())
	} else {
		fmt.Println(err)
	}
	return err
}