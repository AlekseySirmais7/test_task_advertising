package models

type AdsPost struct {
	Id          uint64   `json:"id"`
	Title       string   `json:"title" valid:"length(3|200), required"`
	Description string   `json:"description" valid:"length(3|1000), required"`
	Photos      []string `json:"photos" valid:"length(1|3), required"`
	Price       uint64   `json:"price" valid:"required"`
	Date        string   `json:"date"`
}

type AdsPostId struct {
	Id uint64 `json:"id"`
}

type AdsPostRequest struct {
	Id     uint64   `json:"id" valid:"required"`
	Fields []string `json:"fields" valid:"required"`
}

type AdsPostArrRequest struct {
	Start uint64 `json:"start"`
	Count uint64 `json:"count"`
	Sort  string `json:"sort"`
	Desc  bool   `json:"desc"`
}

type AdsPostArrItem struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Photo string `json:"photo"`
	Price uint64 `json:"price"`
}
