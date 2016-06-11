package slices

import (
	"fmt"
	"reflect"
	"testing"
)

// https://golang.org/cmd/go/#hdr-Description_of_testing_flags
// go test -v -bench string_test.go -benchtime 2s

var ()

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

func TestCut(t *testing.T) {

	if err := expect(Cut([]string{}, 0, -1), []string{}); err != nil {
		t.Error(err)
	}

	if err := expectNot(Cut([]string{}, 0, -1), []string{""}); err != nil {
		t.Error(err)
	}

	if err := expectNot(Cut([]string{""}, 0, -1), []string{""}); err != nil {
		t.Error(err)
	}
}

// func BenchmarkStringsReplaceAll(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		strings.Replace(TestFileContent, ToFind, ToReplace, -1)
// 	}
// }

func expect(args ...interface{}) error {
	switch len(args) {
	case 0, 1:
		return fmt.Errorf("Not enough arguments to expect. Args passed: %v", args)
	default:
		return _expect(true, args...)
	}
}

func expectNot(args ...interface{}) error {
	switch len(args) {
	case 0, 1:
		return fmt.Errorf("Not enough arguments to expect. Args passed: %v", args)
	default:
		return _expect(false, args...)
	}
}

func _expect(boolWanted bool, args ...interface{}) error {

	result := args[0]
	expected := args[1]

	if boolReceived := _areEqual(result, expected); boolReceived != boolWanted {
		return _errExpected(boolWanted, args...)
	}
	return nil
}

func _areEqual(result interface{}, expected interface{}) bool {
	return reflect.DeepEqual(result, expected)
}

func _errExpected(boolWanted bool, args ...interface{}) error {

	condition := "to equal"
	if boolWanted == false {
		condition = "not to equal"
	}

	result := args[0]
	expected := args[1]

	switch len(args) {
	case 2:
		return fmt.Errorf("Expected '%v' %v '%v'", result, condition, expected)
	default:
		return fmt.Errorf("%v: Expected '%v' %v '%v'", args[2], result, condition, expected)
	}
}
