package display

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func Use_person() {
	bob := &Person{"Bob", 22}
	v := reflect.ValueOf(bob)

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i).Interface()

		fmt.Println(fieldName, fieldValue)
	}
}
