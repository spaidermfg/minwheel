package test

import (
	"errors"
	"testing"
)

func foo() (int, error) {
	return 11, nil
}

func bar() (int, error) {
	return 88, errors.New("this is a error info")
}

func TestShadow(t *testing.T) {
	var err error
	defer func() {
		if err != nil {
			println("error in defer:", err.Error())
		}
	}()

	i, err := foo()
	if err != nil {
		return
	}
	println("i=", i)

	if i == 11 {
		//var j int
		j, err := bar()
		if err != nil {
			return
		}
		println("j=", j)
	}
}
