package api

import (
	"api/entity"
	"net/http"
)

type UsecaseApi interface {
	GetApi(http.ResponseWriter, *http.Request) ([]entity.DataApi, error)
}

type RepoApi interface {
	GetApi(http.ResponseWriter, *http.Request) ([]entity.DataApi, error)
}
