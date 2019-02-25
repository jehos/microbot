package utils

import (
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

type Resp struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
	Error   RespError   `json:"error"`
}

type RespError struct {
	Message string `json:"message"`
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Resp{Result: data, Success: true})
}

func RenderErrorJson(w http.ResponseWriter, err error) {
	RenderJson(w, Resp{Success: false, Error: RespError{
		Message: err.Error(),
	}})
}

func Render(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderErrorJson(w, err)
		return
	}
	RenderDataJson(w, data)
}
