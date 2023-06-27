package usecase

import (
	"api/api"
	"api/entity"
	"net/http"
)

func NewUsecaseApi(rep api.RepoApi) usecaseApi {
	return usecaseApi{rep}
}

type usecaseApi struct {
	repo api.RepoApi
}

func (use usecaseApi) GetApi(w http.ResponseWriter, r *http.Request) ([]entity.DataApi, error) {
	res, err := use.repo.GetApi(w, r)

	if err != nil {
		return nil, err
	}

	return res, nil
}
