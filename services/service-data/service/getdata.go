package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (service *DataService) Get(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("Meta service GET for id %s", mux.Vars(r)["id"])

	//id := uuid.MustParse(mux.Vars(r)["id"])

	rw.WriteHeader(http.StatusAccepted)
}
