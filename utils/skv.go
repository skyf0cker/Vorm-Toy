package utils

import (
	"fmt"
	"reflect"
	"time"
)

func SKV(v reflect.Value) ([]string, []string) {
	var keys, values []string
	t := v.Type()
	for n := 0; n < t.NumField(); n++ {
		tf := t.Field(n)
		vf := v.Field(n)

		if tf.Anonymous {
			continue
		}

		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type()).Interface()) {
			continue
		}
		for vf.Type().Kind() == reflect.Ptr {
			vf = vf.Elem()
		}
		if vf.Kind() == reflect.Struct && tf.Type.Name() != "Time"{
			cKeys, cValues := SKV(vf)
			keys = append(keys, cKeys...)
			values = append(values, cValues...)
			continue
		}
		key := tf.Name
		if key == ""{
			continue
		}
		value := format(vf)
		if value != ""{
			keys = append(keys, key)
			values = append(values, value)
		}
	}
	return keys, values
}

func format(v reflect.Value) string {
	if t, ok := v.Interface().(time.Time); ok {
		return fmt.Sprintf("FROM_UNIXTIME(%d)", t.Unix())
	}

	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("'%s'", v.Interface())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return fmt.Sprintf("%d", v.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Interface())
	default:
		panic("not imprlement")
	}
}

func GetTableStruct(v reflect.Value) ([]string, []string) {
	var keys, values []string
	t := v.Type()
	for n := 0; n < t.NumField(); n++ {
		tf := t.Field(n)

		if tf.Anonymous {
			continue
		}
		keys = append(keys, tf.Name)
		var tmpVal string
		switch tf.Type.Name() {
		case "string":
			tmpVal = "TEXT"
		case "int", "int8", "int16", "int32", "int64":
			tmpVal = "INT"
		case "float", "float32", "float64":
			tmpVal = "DOUBLE"
		}
		values = append(values, tmpVal)
	}
	return keys, values
}

