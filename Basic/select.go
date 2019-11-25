package Basic

import (
	"fmt"
	"reflect"
	"strings"
)

func (d *Db)Select(inf interface{}) error {
	var where string
	if d.where != ""{
		where = d.where
	}
	t := reflect.TypeOf(inf)
	v := reflect.ValueOf(inf)

	if t.Kind() == reflect.Ptr{
		v = v.Elem()
	}

	splList := strings.Split(t.String(), ".")
	name := splList[len(splList)-1]

	sql_sen := fmt.Sprintf("SELECT * FROM %s WHERE %s;", name, where)
	fmt.Println(sql_sen)
	rows, err := d.DB.Query(sql_sen)
	if err != nil {
		fmt.Println(err)
	} else {	
		for rows.Next(){
			dest := reflect.New(v.Type())
			dest = dest.Elem()
			dt := dest.Type()
			addrs := make([]interface{}, 0)

			switch dt.Kind() {
			case reflect.Struct:
				for n := 0; n < dt.NumField(); n++ {
					tf := dt.Field(n)
					vf := dest.Field(n)
					if tf.Anonymous {
						continue
					}
					for vf.Type().Kind() == reflect.Ptr {
						vf = vf.Elem()
					}
					column := vf.Type().Name()
					if column == "" {
						continue
					}
					//只取选定的字段的地址
					addrs = append(addrs, vf.Addr().Interface())
					}
				}
			if err := rows.Scan(addrs...); err != nil {
				fmt.Println(err)
			} else {
				v.Set(dest)
			}
		}
	}
	return err
}
