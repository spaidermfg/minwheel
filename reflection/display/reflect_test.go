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
// Name方法返回有确定定义的类型的名字，没有返回空
// String方法返回的类型描述可能包含包名
// Kind方法返回类型的特定类别
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
	fmt.Println(funcTyp.Kind(), "|", funcTyp.String(), "|", funcTyp.Name(), "|", funcVal.String())
}

// 对原生复合类型和自定义类型进行检视
func TestCheckCompoundTyp(t *testing.T) {
	// 复合类型
	// slice
	sl := []int{6, 7}
	sliceVal := reflect.ValueOf(sl)
	sliceTyp := reflect.TypeOf(sl)
	fmt.Println(sliceTyp.Kind(), sliceTyp.Name(), sliceTyp.String())
	fmt.Println(sliceVal.Kind(), sliceVal.String())
	fmt.Println(sliceVal.Index(0).Int(), sliceVal.Index(1).Int())

	// array
	arr := [3]int{2, 6, 7}
	arrayVal := reflect.ValueOf(arr)
	arrayTyp := reflect.TypeOf(arr)
	fmt.Println(arrayTyp.Kind(), arrayTyp.Name(), arrayTyp.String())
	fmt.Println(arrayVal.Index(0).Int(), arrayVal.Index(1).Int(), arrayVal.Index(2).Int())

	// map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	mapVal := reflect.ValueOf(m)
	mapTyp := reflect.TypeOf(m)

	iter := mapVal.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("[%s:%d]\n", k.String(), v.Int())
	}
	fmt.Println(mapTyp.Kind(), mapTyp.String())

	// struct
	type Person struct {
		Name string
		Age  int
	}

	p := Person{"mark", 22}
	structVal := reflect.ValueOf(p)
	structTyp := reflect.TypeOf(p)
	fmt.Println(structTyp.Kind(), structTyp.Name(), structTyp.String())
	fmt.Println(structVal.Field(0).String(), structVal.Field(1).Int())

	// chan
	ch := make(chan int, 1)
	chanVal := reflect.ValueOf(ch)
	chanTyp := reflect.TypeOf(ch)
	ch <- 67
	x, ok := chanVal.TryRecv()
	if ok {
		fmt.Println(x.Int())
	}
	fmt.Println(chanTyp.Kind(), chanTyp.Name(), chanTyp.String())

	//自定义类型
	type myInt int
	var mi myInt = 67
	myVal := reflect.ValueOf(mi)
	myTyp := reflect.TypeOf(mi)
	fmt.Println(myTyp.Kind(), myTyp.Name(), myTyp.String())
	fmt.Println(myVal.Int())
}
