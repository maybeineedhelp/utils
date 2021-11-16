package strsutils

func InSlice(arr []string, str string) bool {
	for _, ele := range arr {
		if ele == str {
			return true
		}
	}
	return false
}

func CompareStrs(a []string, b []string) (aself []string, bself []string, both []string) {
	aMap := make(map[string]struct{})
	for _, v := range a {
		aMap[v] = struct{}{}
	}
	for _, v := range b {
		_, ok := aMap[v]
		if !ok {
			bself = append(bself, v)
			continue
		}
		both = append(both, v)
		delete(aMap, v)
	}
	for v := range aMap {
		aself = append(aself, v)
	}
	return
}

func Reverse(arr []string) (new []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		new = append(new, arr[i])
	}
	return new
}
