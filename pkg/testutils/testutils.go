package testutils

import (
	"errors"
	"reflect"
)

func merge(dst, src reflect.Value, depth int) error {
	if depth > 7 {
		return errors.New("too deep")
	}
	for i := 0; i < src.NumField(); i++ {
		if src.Type().Field(i).Name == "ID" {
			continue
		}

		var field reflect.Value
		if dst.Kind() == reflect.Ptr {
			field = dst.Elem().Field(i)
		} else {
			field = dst.Field(i)
		}

		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			merge(field, src.Field(i), depth+1)
		} else if field.CanSet() && src.Field(i).IsValid() {
			field.Set(src.Field(i))
		} else {
			return errors.New("unable to set value")
		}
	}
	return nil
}
