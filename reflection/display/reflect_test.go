package display

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

// Object Relational Mapping
// 适合处理的问题：输入参数的类型无法提前确定
// 缺点： 代码逻辑混乱；使代码变慢；可能会导致运行时panic

// 三大法则
// 1.反射入口： 经由interface{}类型变量值进入，获得对应的反射对象reflect.Value or reflect.Type
// reflect.Type包含了被反射对象的所有类型信息
// reflect.Value包含了被反射对象的值信息，并且通过Type方法可以获得到与reflect.TypeOf等价的类型信息

// 2.反射出口：反射对象reflect.Value化身接口泪习惯变量值退出
// 3.修改反射对象的前提： reflect.Value必须是可设置的

type Product struct {
	Id        uint32
	Name      string
	Price     uint32
	LeftCount uint32 `orm:"left_count"`
	Batch     string `orm:"batch_number"`
	Updated   time.Time
}

type Person struct {
	ID      string
	Name    string
	Age     uint32
	Gender  string
	Addr    string `orm:"address"`
	Updated time.Time
}

// generate sql
func ConstructQueryStmt(obj interface{}) (string, error) {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return "", errors.New("only struct is supported")
	}

	buf := bytes.NewBufferString("")
	buf.WriteString("SELECT ")

	if typ.NumField() == 0 {
		return "", fmt.Errorf("the type[%s] has no fields", typ.Name())
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if i != 0 {
			buf.WriteString(", ")
		}

		column := field.Name
		if tag := field.Tag.Get("orm"); tag != "" {
			column = tag
		}
		buf.WriteString(column)
	}

	return fmt.Sprintf("%s FROM %s", buf.String(), typ.Name()), nil
}

func TestConstructQueryStmt(t *testing.T) {
	stmt, err := ConstructQueryStmt(&Product{})
	if err != nil {
		log.Fatal("construct query stmt for Product error:", err)
	}
	fmt.Println(stmt)

	stmt1, err1 := ConstructQueryStmt(Person{})
	if err != nil {
		log.Fatal("construct query stmt for Person error:", err1)
	}
	fmt.Println(stmt1)
}

// reflect.ValueOf通过Type方法可以获得到与reflect.TypeOf等价的类型信息
func TestTypeEqual(t *testing.T) {
	i := 67
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)

	fmt.Println(reflect.DeepEqual(val.Type(), typ))
}

// 通过Type和Value对反射实例进行值信息和类型信息的检视
func TestCheckTyp(t *testing.T) {
	// bool
	b := true
	boolVal := reflect.ValueOf(b)
	boolTyp := reflect.TypeOf(b)
	fmt.Println(boolTyp.Name(), boolVal.Bool())

	// int
	i := 67
	intVal := reflect.ValueOf(i)
	intTyp := reflect.TypeOf(i)
	fmt.Println(intTyp.Name(), intVal.Int())

	// string
	s := "less is more and more"
	stringTyp := reflect.TypeOf(s)
	stringVal := reflect.ValueOf(s)
	fmt.Println(stringTyp.Name(), stringVal.String())

	// float
	f := 3.1415926
	floatTyp := reflect.TypeOf(f)
	floatVal := reflect.ValueOf(f)
	fmt.Println(floatTyp.Name(), floatVal.Float())

	// function
	add := func(a, b int) int {
		return a + b
	}
	funcTyp := reflect.TypeOf(add)
	funcVal := reflect.ValueOf(add)
	fmt.Println(funcTyp.Kind(), "|", funcTyp.String(), "|", funcVal.String())
}
