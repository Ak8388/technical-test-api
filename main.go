package main

import (
	"api/entity"
	"api/routers"
	"fmt"
	"net/http"
)

func main() {
	var data entity.Data
	eng := routers.Routers{
		Entity: data,
	}

	eng.Routs()
	fmt.Println("Start Server in Port :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Can't Runing Server")
	}
}
