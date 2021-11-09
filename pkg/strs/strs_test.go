package strsutils_test

import (
	"testing"

	strsutils "github.com/maybeineedhelp/utils/pkg/strs"

	"github.com/go-playground/assert/v2"
)

func TestInSlice(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	assert.Equal(t, strsutils.InSlice(arr, "a"), true)
	assert.Equal(t, strsutils.InSlice(arr, "ae"), false)
}

func TestCompareStrs(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"a", "b", "c"}
	arr1self, arr2self, both := strsutils.CompareStrs(arr1, arr2)
	assert.Equal(t, arr1self, nil)
	assert.Equal(t, arr2self, nil)
	assert.Equal(t, both, []string{"a", "b", "c"})

	arr1 = []string{"a", "c", "e"}
	arr2 = []string{"a", "b", "c"}
	arr1self, arr2self, both = strsutils.CompareStrs(arr1, arr2)
	assert.Equal(t, arr1self, []string{"e"})
	assert.Equal(t, arr2self, []string{"b"})
	assert.Equal(t, both, []string{"a", "c"})
}
