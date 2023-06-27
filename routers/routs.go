package routers

import (
	"api/api/handler"
	"api/api/repository"
	"api/api/usecase"
	"api/entity"
)

type Routers struct {
	Entity entity.Data
}

func (r Routers) Routs() {
	repo := repository.NewRepoApi(r.Entity)
	usecase := usecase.NewUsecaseApi(&repo)
	handler.NewHandlerApi(usecase)
}
