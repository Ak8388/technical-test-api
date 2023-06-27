package entity

type Data struct {
	Uid   int    `json:"userId"`
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type DataApi struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
