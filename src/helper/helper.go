package helper

import "encoding/json"

func SToJson(str any)[]byte {
	res, _ := json.Marshal(str);
	return res
}