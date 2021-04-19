package controller

import (
	"encoding/json"
	"net/http"
	"toko/customerror"
)

func WriteReponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(customerror.GetHTTPResponseCode(err))
	if err != nil {
		resp := Response{Message: err.Error()}
		byteMsg, _ := json.Marshal(resp)
		w.Write(byteMsg)
		return
	}

	resp := Response{Message: "request succeeded"}
	byteMsg, _ := json.Marshal(resp)
	w.Write(byteMsg)
}
