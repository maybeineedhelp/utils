package fileutils

import (
	"encoding/json"
	"io/ioutil"
)

func Parse(file string, to interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}
