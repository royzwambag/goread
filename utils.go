package goread

import "strconv"

// stringSliceToIntSlice takes in a slice of strings and returns a slice of ints
func stringSliceToIntSlice(strings []string) []int {
	ints := make([]int, len(strings))

	for i := range strings {
		ints[i], _ = strconv.Atoi(strings[i])
	}

	return ints
}
