package iplibrary

import (
	"encoding/json"
	"net/http"

	"github.com/iwind/TeaGo/maps"
)

type BaseAction struct {
}

func (this *BaseAction) Close() error {
	return nil
}

// DoHTTP 处理HTTP请求
func (this *BaseAction) DoHTTP(req *http.Request, resp http.ResponseWriter) (goNext bool, err error) {
	return true, nil
}

func (this *BaseAction) convertParams(params maps.Map, ptr interface{}) error {
	data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, ptr)
	if err != nil {
		return err
	}
	return nil
}
