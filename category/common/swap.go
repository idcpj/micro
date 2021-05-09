package common

import "encoding/json"

// swapTo  通过json tag 进行结构体的赋值
func SwapTo(request,category interface{}) (err error){
	dataByte,err:=json.Marshal(request)
	if err!=nil {
		return err
	}
	return json.Unmarshal(dataByte,category)
}

