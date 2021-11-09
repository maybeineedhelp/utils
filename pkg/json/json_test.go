package jsonutils_test

import (
	"testing"
	jsonutils "utils/pkg/json"

	"github.com/go-playground/assert"
)

type InfoTest struct {
	Name   string      `json:"name"`
	Age    int         `json:"age"`
	Parent []*InfoTest `json:"parent"`
}

func TestMarshalJSONOrDie(t *testing.T) {
	info := &InfoTest{
		Name: "胡图图",
		Age:  3,
		Parent: []*InfoTest{
			{
				Name: "图图爸",
				Age:  26,
			},
			{
				Name: "图图妈",
				Age:  24,
			},
		},
	}
	assert.Equal(t, jsonutils.MarshalJSONOrDie(info), `{"name":"胡图图","age":3,"parent":[{"name":"图图爸","age":26,"parent":null},{"name":"图图妈","age":24,"parent":null}]}`)
}
