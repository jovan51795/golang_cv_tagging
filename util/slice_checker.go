package util

import (
	"fmt"
	"reflect"
)

func Contains(slice interface{}, search any) bool {

	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		fmt.Println("Provided data is not a slice")
		return false
	}

	searchValue := reflect.ValueOf(search)
	if searchValue.Kind() == reflect.Ptr {
		searchValue = searchValue.Elem()
	}

	for i := 0; i < s.Len(); i++ {
		sliceField := s.Index(i).Field(1).Interface()
		searchField := searchValue.Field(1).Interface()

		if reflect.DeepEqual(sliceField, searchField) {
			return true
		}
	}

	return false
}
