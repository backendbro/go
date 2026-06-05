package reflection

import "reflect"

func walk1(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk1(field.Interface(), fn)
		}
	}
}

func walk2(x interface{}, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk2(val.Index(i).Interface(), fn)
		}
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk2(field.Interface(), fn)
		}
	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValue := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValue = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValue = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValue; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
