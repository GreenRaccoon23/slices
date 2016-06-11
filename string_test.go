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

	if err := expect(Cut([]string{""}, 0, -1), []string{""}); err != nil {
		t.Error(err)
	}
}

func TestIsSameArray(t *testing.T) {
	slc1 := []string{"same elements,", "different arrays"}
	if slc2 := Cut(slc1, 0, -1); IsSameArray(slc2, slc1) {
		t.Error(errExpected(false, slc2, slc1))
	}
}

func TestDuplicate(t *testing.T) {
	slc1 := []string{"same elements,", "different arrays"}
	if slc2 := Copy(slc1); IsSameArray(slc2, slc1) {
		t.Error(errExpected(false, slc2, slc1))
	}
}

func TestEquals(t *testing.T) {

	slc1 := []string{"same", "elements"}
	slc2 := []string{"same", "elements"}

	if err := expect(Equals(slc1, slc2)); err != nil {
		t.Error(err)
	}
}

func TestPop(t *testing.T) {

	slcBefore := []string{"just enough", "too much"}
	popped, slcAfter := Pop(slcBefore)

	if err := expect(popped, "too much"); err != nil {
		t.Error(err)
	}

	if err := expect(slcAfter, []string{"just enough"}); err != nil {
		t.Error(err)
	}
}

func TestShift(t *testing.T) {

	slcBefore := []string{"next", "enqueued"}
	shifted, slcAfter := Shift(slcBefore)

	if err := expect(shifted, "next"); err != nil {
		t.Error(err)
	}

	if err := expect(slcAfter, []string{"enqueued"}); err != nil {
		t.Error(err)
	}
}

func TestUnshift(t *testing.T) {
	if err := expect(Unshift([]string{"queue"}, "enqueued"), []string{"enqueued", "queue"}); err != nil {
		t.Error(err)
	}
}

// Ridiculously faster
func BenchmarkQueuePushShift(b *testing.B) {

	slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}

	for i := 0; i < b.N; i++ {
		slc = append(slc, "next")
		slc = slc[1:]
	}
}

// Ridiculously slower
func BenchmarkQueueUnshiftPop(b *testing.B) {

	slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}
	iEnd := len(slc) - 1

	for i := 0; i < b.N; i++ {
		slc = append([]string{"next"}, slc...)
		slc = slc[:iEnd]
	}
}

func expect(args ...interface{}) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("Not enough arguments to expect. Args passed: %v", args)
	case 1:
		args = append(args, []interface{}{true}...) // this is nuts, Go
		fallthrough
	default:
		return _expect(true, args...)
	}
}

func expectNot(args ...interface{}) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("Not enough arguments to expect. Args passed: %v", args)
	case 1:
		args = append(args, []interface{}{false}...) // this is nuts, Go
		fallthrough
	default:
		return _expect(false, args...)
	}
}

func _expect(boolWanted bool, args ...interface{}) error {

	result := args[0]
	expected := args[1]

	if boolReceived := _areEqual(result, expected); boolReceived != boolWanted {
		return errExpected(boolWanted, args...)
	}
	return nil
}

func _areEqual(result interface{}, expected interface{}) bool {
	return reflect.DeepEqual(result, expected)
}

func errExpected(boolWanted bool, args ...interface{}) error {

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
