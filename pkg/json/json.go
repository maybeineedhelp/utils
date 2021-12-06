package jsonutils

import "encoding/json"

func MarshalJSONOrDie(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Remarshal(from interface{}, to interface{}) error {
	encoded, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(encoded, to)
}
