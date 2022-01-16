package errorfmt

import (
	"errors"
	"testing"
)

func TestNil(t *testing.T) {
	var err error
	Handlef("foo: %w", &err)
	if err != nil {
		t.Errorf("err == %v, want nil", err)
	}
}

func TestWrap(t *testing.T) {
	err := errors.New("value")
	Handlef("foo: %w", &err)
	g := err.Error()
	w := "foo: value"
	if g != w {
		t.Errorf("g == %q, want %q", g, w)
	}
}

func TestPanicMissingArg(t *testing.T) {
	defer func() {
		recover()
	}()

	Handlef("foo")
	t.Errorf("want panic")
}

func TestPanicNil(t *testing.T) {
	defer func() {
		recover()
	}()

	var badPointer *error
	Handlef("foo", badPointer)
	t.Errorf("want panic")
}
