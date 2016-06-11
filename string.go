// Package slices is a collection of handy methods for string slices.
package slices

import "bytes"

// Contains tests whether any element in a slice matches a string.
func Contains(slc []string, s string) bool {
	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		if slc[i] == s {
			return true
		}
	}
	return false
}

// IsEmpty tests whether a slice has 0 elements
// or is full of empty strings.
func IsEmpty(slc []string) bool {
	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		if slc[i] != "" {
			return false
		}
	}
	return true
}

// Equals tests whether all the elements of two slices are equal.
// These elements do NOT need to point to the same memory location,
// only to hold an equal value.
func Equals(slc1 []string, slc2 []string) bool {

	lenSlc1 := len(slc1)
	lenSlc2 := len(slc2)

	if lenSlc1 != lenSlc2 {
		return false
	}

	iMax := lenSlc1
	for i := 0; i < iMax; i++ {
		if slc1[i] != slc2[i] { //todo: goroutine for unsorted slices?
			return false
		}
	}
	return true
}

// IsSameArray tests whether two slices point to the same array.
func IsSameArray(slc1 []string, slc2 []string) bool {
	return &slc1 == &slc2
}

// Concat concatenates/joins all elements of a slice into a single string.
func Concat(slc []string) string {

	var b bytes.Buffer
	defer b.Reset()

	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		b.WriteString(slc[i])
	}

	return b.String()
}

// Join concatenates/joins all elements of a slice into a single string
// and inserts a common string between each joined element.
func Join(slc []string, by string) string {

	lenSlc := len(slc)

	if isEmpty := (lenSlc == 0); isEmpty {
		return ""
	}

	var b bytes.Buffer
	defer b.Reset()

	iEnd := lenSlc - 1
	for i := 0; i < iEnd; i++ {
		b.WriteString(slc[i])
		b.WriteString(by)
	}
	b.WriteString(slc[iEnd])

	return b.String()
}

// Cut gets a slice of a slice.
// It gets the elements of a slice from index 'start' to index 'stop'.
// 'start' is inclusive (will include the element at that index).
// 'stop' is exclusive (will not include the element at that index).
// If 'start' is 0 and 'stop' is -1, it generates a new copy of the slice.
func Cut(slc []string, start int, stop int) []string {

	if copyRequested := (start == 0 && stop == -1); copyRequested {
		return Copy(slc)
	}

	if goToEnd := (stop == -1); goToEnd {
		return slc[start:]
	}

	return slc[start:stop]
}

// Copy generates a full copy of a slice,
// i.e., one which points to a different underlying array.
func Copy(slc []string) []string {
	newSlc := make([]string, len(slc))
	copy(newSlc, slc)
	return newSlc
}

// Compact generates a copy of a slice with any empty strings removed.
// The slice is not modified in place; the original will be unchanged.
func Compact(bloated []string) (compacted []string) {
	lenBloated := len(bloated)
	for i := 0; i < lenBloated; i++ {
		if s := bloated[i]; s != "" {
			compacted = append(compacted, s)
		}
	}
	return
}

// Filter removes elements from a slice.
// It returns a copy of a slice with unwanted strings removed.
// The slice is not modified in place; the original will be unchanged.
func Filter(unfiltered []string, unwanted ...string) (filtered []string) {

	lenUnfiltered := len(unfiltered)
	for i := 0; i < lenUnfiltered; i++ {
		s := unfiltered[i]

		if isUnwanted := Contains(unwanted, s); isUnwanted {
			continue
		}

		filtered = append(filtered, s)
	}

	return
}

// Extract gets elements from a slice.
// It return a new slice of the elements pulled from the original.
// The new slice contains only the 'wanted' elements
// which the original 'excess' slice contains.
func Extract(excess []string, wanted ...string) (extracted []string) {

	lenExcess := len(excess)
	for i := 0; i < lenExcess; i++ {
		s := excess[i]

		if isWanted := Contains(wanted, s); isWanted {
			extracted = append(extracted, s)
		}
	}

	return
}

// Push adds elements to a slice and returns the modified slice.
// It is a direct call to the built-in 'append()' func.
// It is meant to be clear, readable method for stack implementations.
func Push(slc []string, args ...string) []string {
	return append(slc, args...)
}

// Pop removes the last element of a slice.
// It return the removed element along with the modified slice.
func Pop(slc []string) (string, []string) {

	iEnd := len(slc) - 1

	popped := slc[iEnd]
	cut := slc[:iEnd]

	return popped, cut
}

// Unshift inserts an element at the beginning of a slice,
// and moves the rest of the elements up an index.
// It does not overwrite the first element.
// It returns the modified slice.
func Unshift(slc []string, s string) []string {
	return append([]string{s}, slc...)
}

// Shift removes the first element of a slice.
// It returns the removed element along with the modified slice.
func Shift(slc []string) (string, []string) {
	return slc[0], slc[1:]
}
