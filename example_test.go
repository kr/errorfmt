package errorfmt_test

import (
	"errors"
	"fmt"

	"kr.dev/errorfmt"
)

func ExampleHandlef() {
	err := Frob("baz")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// frob widget baz: step2 failed
}

func Frob(name string) (err error) {
	defer errorfmt.Handlef("frob widget %s: %w", name, &err)

	err = step1()
	if err != nil {
		return err
	}

	err = step2()
	if err != nil {
		return err
	}
	return nil
}

func step1() error { return nil }
func step2() error { return errors.New("step2 failed") }
