package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"toko/usecase"

	log "github.com/sirupsen/logrus"
)

type controllerHttp struct {
	uc usecase.Usecase
}

func NewControllerHttp(uc usecase.Usecase) *controllerHttp {
	return &controllerHttp{uc: uc}
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
