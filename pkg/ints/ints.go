package intsutils

func InSlice(arr []int, str int) bool {
	for _, ele := range arr {
		if ele == str {
			return true
		}
	}
	return false
}
