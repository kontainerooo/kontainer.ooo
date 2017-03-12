package testutils

import (
	"errors"
	"reflect"
)

func merge(dst, src reflect.Value, overwriteID bool, depth int) error {
	if depth > 7 {
		return errors.New("too deep")
	}
	for i := 0; i < src.NumField(); i++ {
		if !overwriteID && src.Type().Field(i).Name == "ID" {
			continue
		}

		var field reflect.Value
		if dst.Kind() == reflect.Ptr {
			field = dst.Elem().Field(i)
		} else {
			field = dst.Field(i)
		}

		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			merge(field, src.Field(i), overwriteID, depth+1)
		} else if field.CanSet() {
			srcVal := src.Field(i)
			if srcVal.IsValid() && (!isZero(srcVal)) {
				field.Set(srcVal)
			}
		} else {
			return errors.New("unable to set value")
		}
	}
	return nil
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}
