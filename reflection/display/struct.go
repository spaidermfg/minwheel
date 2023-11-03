package display

import (
	"fmt"
	"reflect"
)

type Persons struct {
	Name string
	Age  int
}

func Use_persons() {
	bob := &Persons{"Bob", 22}
	v := reflect.ValueOf(bob)

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i).Interface()

		fmt.Println(fieldName, fieldValue)
	}
}
