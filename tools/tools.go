package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func OpenJson(path string) map[string]interface{} {
	archive, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var f interface{}
	json.Unmarshal(archive, &f)
	return f.(map[string]interface{})
}

func ConvertString(list_interface map[string]interface{}, key_value string) map[string]string {
	result := map[string]string{}
	interf := list_interface[key_value].(map[string]interface{})

	for k := range interf {
		if v, ok := interf[k].(string); ok {
			result[k] = v
		}
	}
	return result

}
