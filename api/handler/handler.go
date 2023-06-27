package handler

import (
	"api/api"
	"encoding/json"
	"net/http"
)

func NewHandlerApi(use api.UsecaseApi) {
	eng := handApi{use}

	http.HandleFunc("/get-api", eng.GetApi)
}

type handApi struct {
	usecase api.UsecaseApi
}

func (hand handApi) GetApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := hand.usecase.GetApi(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dataJson, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "apllication/json")
		w.Write(dataJson)
		// w.WriteHeader(http.StatusOK)
		return
	}
}
