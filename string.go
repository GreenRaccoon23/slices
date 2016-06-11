package slices

import "bytes"

// Return true if any element in a slice matches a string.
func Contains(slc []string, s string) bool {
	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		if slc[i] == s {
			return true
		}
	}
	return false
}

// Return true if a slice contains 0 elements
// or if all elements in a slice have lengths of 0.
func IsEmpty(slc []string) bool {
	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		if slc[i] != "" {
			return false
		}
	}
	return true
}

// Return true if the elements of both slices have equal values.
// In order for elements to equal each other,
// they do NOT need to point to the same memory location;
// they only need to hold an equal value.
func Equals(slc1 []string, slc2 []string) bool {

	lenSlc1 := len(slc1)
	lenSlc2 := len(slc2)

	if lenSlc1 != lenSlc2 {
		return false
	}

	iMax := lenSlc1
	for i := 0; i < iMax; i++ {
		if slc1[i] != slc2[i] {
			return false
		}
	}
	return true
}

// Return true if both slices point to the same array.
func IsSameArray(slc1 []string, slc2 []string) bool {
	return &slc1 == &slc2
}

// Concatenate/Join all elements of a slice into a single string.
func Concat(slc []string) string {

	var b bytes.Buffer
	defer b.Reset()

	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		b.WriteString(slc[i])
	}

	return b.String()
}

// Concatenate/Join all elements of a slice into a single string.
// Separate each element by a provided string.
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

// Get a slice of a slice.
// Get the elements of a slice from index 'start' to index 'stop'.
// 'start' is inclusive (include the element at that index).
// 'stop' is exclusive (do not include the element at that index).
func Cut(slc []string, start int, stop int) []string {

	if copyRequested := (start == 0 && stop == -1); copyRequested {
		return Copy(slc)
	}

	if goToEnd := (stop == -1); goToEnd {
		return slc[start:]
	}

	return slc[start:stop]
}

// Return a copy of a slice (different underlying array).
func Copy(slc []string) []string {
	newSlc := make([]string, len(slc))
	copy(newSlc, slc)
	return newSlc
}

// Return a copy of a slice with any empty strings removed.
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

// Remove elements from a slice.
// Return a copy of a slice with unwanted strings removed.
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

// Get elements from a slice.
// Return a new slice of the elements pulled from the original.
// Only get elements if the slice contains them.
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

// Add elements to a slice. Return the modified slice.
func Push(slc []string, args ...string) []string {
	return append(slc, args...)
}

// Remove the last element of a slice.
// Return the removed element along with the modified slice.
func Pop(slc []string) (string, []string) {

	iEnd := len(slc) - 1

	popped := slc[iEnd]
	cut := slc[:iEnd]

	return popped, cut
}

// Insert an element at the beginning of a slice,
// and move all the rest of the elements up an index.
// Return the modified slice.
func Unshift(slc []string, s string) []string {
	return append([]string{s}, slc...)
}

// Remove the first element of a slice.
// Return the removed element along with the modified slice.
func Shift(slc []string) (string, []string) {
	return slc[0], slc[1:]
}
