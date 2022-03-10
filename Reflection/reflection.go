package main

import "reflect"

// This code is very unsafe and very naive
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// NumField on a pointer Value
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)

		}

	}
}
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
