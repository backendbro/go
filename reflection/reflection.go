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

func walk3(x interface{}, fn func(input string)) {
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
		walk3(getField(i).Interface(), fn)
	}
}

func walk4(x interface{}, fn func(input string)) {
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
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk4(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValue; i++ {
		walk4(getField(i).Interface(), fn)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
