package util

import "encoding/json"

func StructToJson(value interface{}) (res string, err error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func JsonToStruct(data string, value interface{}) error {
	return json.Unmarshal([]byte(data), value)
}

func ToJSON(a interface{}) string {
	if a == nil {
		return ""
	}
	b, _ := json.Marshal(a)
	return string(b)
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
