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

func TestReverse(t *testing.T) {
	assert.Equal(t, strsutils.Reverse([]string{"a", "b", "c"}), []string{"c", "b", "a"})
	assert.Equal(t, strsutils.Reverse([]string{"a", "b", "c", "d"}), []string{"d", "c", "b", "a"})
}

func TestHasElementInCommon(t *testing.T) {
	assert.Equal(t, strsutils.HasElementInCommon([]string{"a", "b", "c"}, []string{"v", "c"}), true)
	assert.Equal(t, strsutils.HasElementInCommon([]string{"a", "d"}, []string{"v", "c"}), false)
}

func TestCountElementInCommon(t *testing.T) {
	assert.Equal(t, strsutils.CountElementInCommon([]string{"a", "b", "c"}, []string{"v", "c", "c"}), 2)
	assert.Equal(t, strsutils.CountElementInCommon([]string{"a", "d"}, []string{"v", "c"}), 0)
}
