package utils

import "reflect"

func GetStructKeys(s interface{}) []string {
	var keys []string
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		keys = append(keys, field.Name)
	}

	return keys
}
