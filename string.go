package slices

import "bytes"

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
//   they do NOT need to point to the same memory location;
//   they only need to hold an equal value.
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
// Probably not useful.
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
// Get the elements of a slice from index 'iBeg' to index 'iEnd'.
// 'iBeg' is inclusive (include the element at that index).
// 'iEnd' is exclusive (do not include the element at that index).
func Cut(slc []string, iBeg int, iEnd int) []string {

	if copyRequested := (iBeg == 0 && iEnd == -1); copyRequested {
		return Copy(slc)
	}

	if goToEnd := (iEnd == -1); goToEnd {
		return slc[iBeg:]
	}

	return slc[iBeg:iEnd]
}

// func Cut(slc []string, iBeg int, iEnd int) []string {

// 	if copyAll := (iBeg == 0 && iEnd == -1); copyAll {
// 		return Copy(slc)
// 	}

// 	if goToEnd := (iEnd == -1); goToEnd {
// 		return slc[iBeg:]
// 	}

// 	return slc[iBeg:iEnd]
// }

func Copy(slc []string) []string {
	newSlc := make([]string, len(slc))
	copy(newSlc, slc)
	return newSlc
}

// Remove the last element of a slice.
// Return the removed element along with the modified slice.
func Pop(slc []string) (string, []string) {

	iEnd := len(slc) - 1

	popped := slc[iEnd]
	cut := slc[:iEnd]

	return popped, cut
}

// func shift(slc []string) (string, []string) {
// 	return slc[0], slc[1:]
// }

// func unshift(slc []string, s string) []string {
// 	return append([]string{s}, slc...)
// }

func slcContains(slc []string, s string) bool {
	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		if slc[i] == s {
			return true
		}
	}
	return false
}

func slcIsEmpty(slc []string) bool {
	lenSlc := len(slc)
	for i := 0; i < lenSlc; i++ {
		if slc[i] != "" {
			return false
		}
	}
	return true
}

func compact(args ...string) (compacted []string) {
	lenArgs := len(args)
	for i := 0; i < lenArgs; i++ {
		if s := args[i]; s != "" {
			compacted = append(compacted, s)
		}
	}
	return
}

func isEmpty(args ...string) bool {
	lenArgs := len(args)
	for i := 0; i < lenArgs; i++ {
		if notEmpty := args[i] != ""; notEmpty {
			return false
		}
	}
	return true
}

func filter(unfiltered []string, unwanted ...string) (filtered []string) {

	lenUnfiltered := len(unfiltered)
	for i := 0; i < lenUnfiltered; i++ {
		s := unfiltered[i]

		if isUnwanted := slcContains(unwanted, s); isUnwanted {
			continue
		}

		filtered = append(filtered, s)
	}

	return
}

func extract(excess []string, wanted ...string) (extracted []string) {

	lenExcess := len(excess)
	for i := 0; i < lenExcess; i++ {
		s := excess[i]

		if isWanted := slcContains(wanted, s); isWanted {
			extracted = append(extracted, s)
		}
	}

	return
}
