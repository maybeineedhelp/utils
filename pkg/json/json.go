package jsonutils

import "encoding/json"

func MarshalJSONOrDie(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(b)
}
