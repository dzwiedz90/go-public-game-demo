package utils

// RemoveElement takes a slice of strings, removes element from the given argument and returns original slice without this element
func RemoveElement(slice []string, element string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}
