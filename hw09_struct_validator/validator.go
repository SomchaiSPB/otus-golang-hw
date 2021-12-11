package hw09structvalidator

import (
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	iv := reflect.ValueOf(v)

	if iv.Kind() != reflect.Struct {
		return nil
	}

	t := iv.Type()

	structMap := make(map[string]interface{}, iv.NumField())

	for i := 0; i < iv.NumField(); i++ {
		field := t.Field(i)
		val := iv.Field(i)

		if val.CanInterface() {
			structMap[field.Name] = val.Interface()
		}
	}

	fmt.Println(structMap)

	return nil
}
