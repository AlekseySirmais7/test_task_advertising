package uniqueSlice

func IsUniqueStrings(checkingSlice []string) bool {
	uniqueMap := make(map[string]bool)
	for _, item := range checkingSlice {
		if _, ok := uniqueMap[item]; !ok {
			uniqueMap[item] = true
		} else {
			return false
		}
	}
	return true
}
