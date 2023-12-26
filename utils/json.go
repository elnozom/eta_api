package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func Serialize(documentStructure interface{}) string {
	var serializedString string

	switch reflect.TypeOf(documentStructure).Kind() {
	case reflect.String, reflect.Int, reflect.Float64, reflect.Bool:
		return fmt.Sprintf("\"%v\"", documentStructure) // Wrap simple values in quotes
	case reflect.Slice:
		// Handle arrays of any type
		s := reflect.ValueOf(documentStructure)
		for i := 0; i < s.Len(); i++ {
			if i > 0 {
				serializedString += "," // Add comma separators between elements
			}
			serializedString += Serialize(s.Index(i).Interface()) // Recursively serialize elements of any type
		}
	case reflect.Struct:
		t := reflect.TypeOf(documentStructure)
		v := reflect.ValueOf(documentStructure)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)
			serializedString += fmt.Sprintf("\"%s\"%s",
				strings.ToUpper(field.Name), // Convert key to uppercase and wrap in quotes
				Serialize(value.Interface()))
		}
	}

	return serializedString
}
