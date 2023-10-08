package errorutil

import (
	"errors"
	"testing"
)

func TestAddParentError(t *testing.T) {
	if AddParentError(nil, nil) != nil {
		t.Error("must return nil when parameters are nil")
	}
	e1 := errors.New("test")
	e2 := errors.New("test2")
	if r := AddParentError(e1, nil); r != e1 {
		t.Errorf("expected %v but got %v", e1, r)
	}
	if AddParentError(e1, e2) == nil {
		t.Error("expected result to be non-nil")
	}
}

func TestAddParentErrorToErrors(t *testing.T) {
	errs := []error{errors.New("foo")}
	if l := len(AddParentErrorToErrors(errs, nil)); l != 1 {
		t.Errorf("expected to have an array of length 1 but got %v elements", l)
	}
	errs = append(errs, nil)
	if l := len(AddParentErrorToErrors(errs, nil)); l != 1 {
		t.Errorf("expected to have an array of length 1 but got %v elements", l)
	}
	errs = append(errs, errors.New("bar"))
	if l := len(AddParentErrorToErrors(errs, nil)); l != 2 {
		t.Errorf("expected to have an array of length 2 but got %v elements", l)
	}
}
