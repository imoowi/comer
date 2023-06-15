package copy

import "reflect"

func CopySctruct(from, to interface{}) {
	sval := reflect.ValueOf(from).Elem()
	dval := reflect.ValueOf(to).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}
