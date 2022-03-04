package utils

import (
	jsoniter "github.com/json-iterator/go"
)

var Json = newDefaultJson()

type defaultJson struct {
	json jsoniter.API
}

func newDefaultJson() *defaultJson {
	_json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		UseNumber:              true,
	}.Froze()

	return &defaultJson{json: _json}
}

func (j *defaultJson) Marshal(v interface{}) (Bytes, error) {
	return j.json.Marshal(v)
}

func (j *defaultJson) Unmarshal(data interface{}, v interface{}) error {
	var b []byte
	switch data.(type) {
	case string:
		b = []byte(data.(string))
		break
	case Bytes:
		b = []byte(data.(Bytes))
		break
	case []byte:
		b = data.([]byte)
		break

	default:
		panic("data must be []byte or string")
	}

	return j.json.Unmarshal(b, v)
}
