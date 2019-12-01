package utils

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// ListOfErrors error handler
func ListOfErrors(object interface{}, e error) []map[string]string {
	ve := e.(validator.ValidationErrors)
	InvalidFields := make([]map[string]string, 0)

	for _, e := range ve {
		errors := map[string]string{}
		// field := reflect.TypeOf(e.NameNamespace)
		field, _ := reflect.TypeOf(object).Elem().FieldByName(e.Name)
		jsonTag := string(field.Tag.Get("json"))
		errors[jsonTag] = e.Tag
		InvalidFields = append(InvalidFields, errors)
	}

	return InvalidFields
}
