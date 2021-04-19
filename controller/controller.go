package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"toko/customerror"
	"toko/usecase"

	log "github.com/sirupsen/logrus"
)

type controllerHttp struct {
	uc usecase.Usecase
}

func NewControllerHttp(uc usecase.Usecase) *controllerHttp {
	return &controllerHttp{uc: uc}
}

type Response struct {
	Message string `json:"message"`
}

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

func (c *controllerHttp) Buy(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		WriteReponse(w, err)
		return
	}

	req := BuyRequest{}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		log.Error(err)
		WriteReponse(w, err)
		return
	}

	err = c.uc.Buy(req.SKU, req.Quantity)
	if err != nil {
		log.Error(err)
		WriteReponse(w, err)
		return
	}

	WriteReponse(w, nil)
}
