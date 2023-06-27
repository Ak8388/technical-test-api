package repository

import (
	"api/entity"
	"encoding/json"
	"io"
	"net/http"
)

func NewRepoApi(r entity.Data) repo {
	return repo{r}
}

type repo struct {
	entity entity.Data
}

func (rep *repo) GetApi(w http.ResponseWriter, r *http.Request) ([]entity.DataApi, error) {
	var data []entity.Data
	var dat []entity.DataApi
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)

	if err != nil {
		return nil, err
	}

	res, err2 := client.Do(req)

	if err2 != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode == 200 {
		err := json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
	}

	for _, json := range data {
		dat = append(dat, entity.DataApi{Id: json.Id, Title: json.Title, Body: json.Body})
	}

	return dat, nil
}
