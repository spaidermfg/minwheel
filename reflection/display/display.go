package display

import (
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// 不导出，隐藏实现细节
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			s := fmt.Sprintf("%s[%d]", path, i)
			display(s, v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			s := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(s, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			s := fmt.Sprintf("%s[%s]", path, formatAtom(key))
			display(s, v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			s := fmt.Sprintf("(*%s)", path)
			display(s, v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s =  nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// Any 把任何值格式转化为string
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 0, 0, 0)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Map, reflect.Slice:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}
