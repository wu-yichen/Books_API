package model

type Book struct {
	Id    int    `json:id`
	Title string `json:title`
	Year  string `json:year`
}
