package intsutils_test

import (
	"testing"

	intsutils "github.com/maybeineedhelp/utils/pkg/ints"

	"github.com/go-playground/assert/v2"
)

func TestInSlice(t *testing.T) {
	arr := []int{1, 34, 56, 2}
	assert.Equal(t, intsutils.InSlice(arr, 34), true)
	assert.Equal(t, intsutils.InSlice(arr, 79), false)
}
