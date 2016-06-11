package slices

import (
	"fmt"
	"reflect"
	"testing"
)

// https://golang.org/cmd/go/#hdr-Description_of_testing_flags
// go test -v string*
// go test -v string* -bench '.' -benchtime 2s

var (
	DoFailAll bool = false
	// DoFailAll bool = true
)

func TestContains(t *testing.T) {
	theGriffinFamily := []string{"the", "Griffin", "family"}
	if err := expectNot(Contains(theGriffinFamily, "Meg")); err != nil {
		t.Error(err)
	}

	anyOldMatter := []string{"a", "snail", "on", "the", "tail", "of", "the", "frog", "on", "the", "bump", "on", "this", "log", "that", "I", "found", "in", "a", "hole", "in", "the", "bottom", "of", "the", "sea"}
	errPrefix := "Ehhh wha? Dr. Farnsworth"

	if err := expect(Contains(anyOldMatter, "snail"), true, errPrefix); err != nil {
		t.Error(err)
	}

	if err := expect(Contains(anyOldMatter, "the ultimate secret of the universe"), false, errPrefix); err != nil {
		t.Error(err)
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

func TestEquals(t *testing.T) {
	slc1 := []string{"same", "elements"}
	slc2 := []string{"same", "elements"}
	if err := expect(Equals(slc1, slc2)); err != nil {
		t.Error(err)
	}
}

func TestIsSameArray(t *testing.T) {
	slc1 := []string{"same elements,", "different arrays"}
	if slc2 := Cut(slc1, 0, -1); IsSameArray(slc2, slc1) {
		t.Error(errExpected(false, slc2, slc1))
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

func TestCopy(t *testing.T) {
	slc1 := []string{"same elements,", "different arrays"}
	if slc2 := Copy(slc1); IsSameArray(slc2, slc1) {
		t.Error(errExpected(false, slc2, slc1))
	}
}

func TestCompact(t *testing.T) {
	thsUS := []string{"Gooo", "", "", "", "", "", "", "", "", "", "", "d", "mor", "", "ning", "Vietnam!"}
	theUK := []string{"Gooo", "d", "mor", "ning", "Vietnam!"}
	errPrefix := "Robin Williams"
	if err := expect(Compact(thsUS), theUK, errPrefix); err != nil {
		t.Error(err)
	}
}

func TestFilter(t *testing.T) {

	legalFamilyMembers := []string{"Peter", "Lois", "Chris", "Stewie", "Brian", "Meg"}
	wantedFamilyMembers := []string{"Peter", "Lois", "Chris", "Stewie", "Brian"}

	errPrefix := "Meg, who let you back in the house?"

	if err := expect(Filter(legalFamilyMembers, "Meg"), wantedFamilyMembers, errPrefix); err != nil {
		t.Error(err)
	}
}

func TestExtract(t *testing.T) {

	theGriffins := []string{"Peter", "Lois", "Chris", "Stewie", "Brian"}
	if DoFailAll {
		theGriffins = []string{"Peter", "Lois", "Chris", "Stewie", "Brian", "Meg"}
	}

	theBestGriffins := Extract(theGriffins, "Peter", "Stewie", "Meg")
	expected := []string{"Peter", "Stewie"}

	if err := expect(theBestGriffins, expected); err != nil {
		t.Error(err)
	}
	if shutUpMeg := Contains(theBestGriffins, "Meg"); shutUpMeg {
		t.Error(errExpected(false, theBestGriffins, expected, "Shut up, Meg."))
	}
}

func TestPush(t *testing.T) {

	mamaBear := []string{"not enough"}
	babyBear := []string{"not enough", "just right"}

	goldilocks := Push(mamaBear, "just right")

	if err := expect(goldilocks, babyBear, "Goldilocks"); err != nil {
		t.Error(err)
	}
}

func TestPop(t *testing.T) {

	extraPorridgeHeat := "too much"

	papaBear := []string{"just right", extraPorridgeHeat}
	babyBear := []string{"just right"}

	porridgeReview, goldilocks := Pop(papaBear)

	if err := expect(porridgeReview, extraPorridgeHeat, "Goldilocks"); err != nil {
		t.Error(err)
	}
	if err := expect(goldilocks, babyBear, "Goldilocks"); err != nil {
		t.Error(err)
	}
}

func TestUnshift(t *testing.T) {
	if err := expect(Unshift([]string{"queue"}, "enqueued"), []string{"enqueued", "queue"}); err != nil {
		t.Error(err)
	}
}

func TestShift(t *testing.T) {

	dmvNow := []string{"hope", "eternity"}
	shifted, dmvTenYearsLater := Shift(dmvNow)

	if err := expect(shifted, "hope"); err != nil {
		t.Error(err)
	}
	if err := expect(dmvTenYearsLater, []string{"eternity"}, "Everyone at the DMV"); err != nil {
		t.Error(err)
	}
}

func expect(args ...interface{}) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("Not enough arguments to expect(). Args passed: %v", args)
	case 1:
		return expect(args[0], true)
	default:
		return _expect(true, args...)
	}
}

func expectNot(args ...interface{}) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("Not enough arguments to expect(). Args passed: %v", args)
	case 1:
		return expectNot(args[0], true)
	default:
		return _expect(false, args...)
	}
}

func _expect(boolExpected bool, args ...interface{}) error {

	result := args[0]
	expected := args[1]

	if boolReceived := _areEqual(result, expected); boolReceived != boolExpected || DoFailAll {
		return errExpected(boolExpected, args...)
	}
	return nil
}

func _areEqual(result interface{}, expected interface{}) bool {
	return reflect.DeepEqual(result, expected)
}

func errExpected(boolExpected bool, args ...interface{}) error {

	conditionExpected := "to equal"
	if boolExpected == false {
		conditionExpected = "not to equal"
	}

	result := args[0]
	expected := args[1]

	switch len(args) {
	case 2:
		return fmt.Errorf("Expected '%v' %v '%v'", result, conditionExpected, expected)
	default:
		return fmt.Errorf("%v expected '%v' %v '%v'", args[2], result, conditionExpected, expected)
	}
}

var (
	DoBenchmarkAppendPush bool = true
	// DoBenchmarkAppendPush bool = false
)

func BenchmarkAppend(b *testing.B) {
	if !DoBenchmarkAppendPush {
		return
	}
	for i := 0; i < b.N; i++ {
		slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}
		slc = append(slc, "next")
	}
}

func BenchmarkPush(b *testing.B) {
	if !DoBenchmarkAppendPush {
		return
	}
	for i := 0; i < b.N; i++ {
		slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}
		slc = Push(slc, "next")
	}
}

var (
	DoBenchmarkQueueStack bool = true
	// DoBenchmarkQueueStack bool = false
)

// Ridiculously faster
func BenchmarkQueuePushShift(b *testing.B) {
	if !DoBenchmarkQueueStack {
		return
	}
	slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}

	for i := 0; i < b.N; i++ {
		slc = Push(slc, "next")
		_, slc = Shift(slc)
	}
}

// Ridiculously slower
func BenchmarkQueueUnshiftPop(b *testing.B) {
	if !DoBenchmarkQueueStack {
		return
	}
	slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}

	for i := 0; i < b.N; i++ {
		slc = Unshift(slc, "next")
		_, slc = Pop(slc)
	}
}

// Ridiculously faster
func BenchmarkStackPushPop(b *testing.B) {
	if !DoBenchmarkQueueStack {
		return
	}
	slc := []string{"tail", "of", "the", "frog", "on", "the", "bump", "on", "this", "log", "that", "I", "found", "in", "a", "hole", "in", "the", "bottom", "of", "the", "sea"}

	for i := 0; i < b.N; i++ {
		slc = Push(slc, "snail on the")
		_, slc = Pop(slc)
	}
}

// Ridiculously slower
func BenchmarkStackUnshiftShift(b *testing.B) {
	if !DoBenchmarkQueueStack {
		return
	}
	slc := []string{"tail", "of", "the", "frog", "on", "the", "bump", "on", "this", "log", "that", "I", "found", "in", "a", "hole", "in", "the", "bottom", "of", "the", "sea"}

	for i := 0; i < b.N; i++ {
		slc = Unshift(slc, "snail on the")
		_, slc = Shift(slc)
	}
}

var (
	DoBenchmarkQueueStackNative bool = true
	// DoBenchmarkQueueStackNative bool = false
)

// Ridiculously faster
func BenchmarkQueuePushShiftNative(b *testing.B) {
	if !DoBenchmarkQueueStackNative {
		return
	}
	slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}
	for i := 0; i < b.N; i++ {
		slc = append(slc, "next")
		slc = slc[1:]
	}
}

// Ridiculously slower
func BenchmarkQueueUnshiftPopNative(b *testing.B) {
	if !DoBenchmarkQueueStackNative {
		return
	}
	slc := []string{"deeporange50", "deeporange500", "deeporange900", "deeporangeA100", "deeporangeA200", "deeporangeA400", "deeporangea700", "deeporange-50", "deeporange:500", "deeporange_900", "deeporange A100", "deeporange-A200", "deeporange:A400", "deeporange_A700", "deeporange000", "deeporange1000", "deeporangeA000", "deeporangeA300", "deeporangeA1000", "deeporange:000", "deeporange-1000", "deeporange_A000", "deeporange A300", "deeporange-A1000", "deeporange"}
	iEnd := len(slc) - 1
	for i := 0; i < b.N; i++ {
		slc = append([]string{"next"}, slc...)
		slc = slc[:iEnd]
	}
}

// Ridiculously faster
func BenchmarkStackPushPopNative(b *testing.B) {
	if !DoBenchmarkQueueStackNative {
		return
	}
	slc := []string{"tail", "of", "the", "frog", "on", "the", "bump", "on", "this", "log", "that", "I", "found", "in", "a", "hole", "in", "the", "bottom", "of", "the", "sea"}
	iEnd := len(slc) - 1
	for i := 0; i < b.N; i++ {
		slc = append(slc, "snail on the")
		slc = slc[:iEnd]
	}
}

// Ridiculously slower
func BenchmarkStackUnshiftShiftNative(b *testing.B) {
	if !DoBenchmarkQueueStackNative {
		return
	}
	slc := []string{"tail", "of", "the", "frog", "on", "the", "bump", "on", "this", "log", "that", "I", "found", "in", "a", "hole", "in", "the", "bottom", "of", "the", "sea"}
	for i := 0; i < b.N; i++ {
		slc = append([]string{"snail on the"}, slc...)
		slc = slc[1:]
	}
}
