package slices

import (
	"fmt"
	"reflect"
	"testing"
)

// https://golang.org/cmd/go/#hdr-Description_of_testing_flags
// go test -v -bench string_test.go -benchtime 2s

var ()

func expect(args ...interface{}) error {
	switch len(args) {
	case 0, 1:
		return fmt.Errorf("Not enough arguments to _expect. Args passed: %v", args)
	case 2:
		return _expectBasic(args[0], args[1])
	default:
		return _expectLabelled(args[0], args[1], args[2])
	}
}

func _expectBasic(result interface{}, expected interface{}) error {
	if !_areEqual(result, expected) {
		return fmt.Errorf("Expected '%v' to equal '%v'", result, expected)
	}
	return nil
}

func _expectLabelled(result interface{}, expected interface{}, label interface{}) error {
	if !_areEqual(result, expected) {
		return fmt.Errorf("%v: Expected '%v' to equal '%v'", label, result, expected)
	}
	return nil
}

func _areEqual(result interface{}, expected interface{}) bool {
	return reflect.DeepEqual(result, expected)
}

func expected(args ...interface{}) error {
	switch len(args) {
	case 0, 1:
		return fmt.Errorf("Not enough arguments to _errExpected. Args passed: %v", args)
	case 2:
		return fmt.Errorf("Expected '%v' to equal '%v'", args[0], args[1])
	default:
		return fmt.Errorf("%v: Expected '%v' to equal '%v'", args[0], args[1], args[2])
	}
}

func TestIsEmpty(t *testing.T) {

	if err := expect(IsEmpty([]string{}), true); err != nil {
		t.Error(err)
	}

	if err := expect(IsEmpty([]string{""}), true); err != nil {
		t.Error(err)
	}

	if err := expect(IsEmpty([]string{"a", "b", "c"}), false); err != nil {
		t.Error(err)
	}
}

func TestConcat(t *testing.T) {

	if err := expect(Concat([]string{}), ""); err != nil {
		t.Error(err)
	}

	if err := expect(Concat([]string{""}), ""); err != nil {
		t.Error(err)
	}

	if err := expect(Concat([]string{"a", "b", "c"}), "abc"); err != nil {
		t.Error(err)
	}
}

func TestJoin(t *testing.T) {

	if err := expect(Join([]string{}, "-"), ""); err != nil {
		t.Error(err)
	}

	if err := expect(Join([]string{""}, "-"), ""); err != nil {
		t.Error(err)
	}

	if err := expect(Join([]string{"a", "b", "c"}, "-"), "a-b-c"); err != nil {
		t.Error(err)
	}
}

func toPointer(s []string) *[]string {
	return &s
}

func TestCut(t *testing.T) {

	s := []string{}
	u := []string{}
	t.Logf("1%#v:", toPointer(s))
	t.Logf("2%#v:", toPointer(u))

	if err := expect(Cut([]string{}, 0, -1), []string{}); err != nil {
		t.Error(err)
	}

	if err := expect(Cut([]string{}, 0, -1), []string{""}); err != nil {
		t.Error(err)
	}

	// if err := expect(Cut([]string{""}, "-"), ""); err != nil {
	// 	t.Error(err)
	// }

	// if err := expect(Cut([]string{"a", "b", "c"}, "-"), "a-b-c"); err != nil {
	// 	t.Error(err)
	// }
}

// func BenchmarkStringsReplaceAll(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		strings.Replace(TestFileContent, ToFind, ToReplace, -1)
// 	}
// }