package tool

import "encoding/json"

// struct to json
// Path: internal/tool/bytes.go

func ToBytes(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}
