package common

import "encoding/json"

func SwapTo(request interface{},product interface{}) error {
	req, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(req, product)
}
