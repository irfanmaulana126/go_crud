package paginationHelper

type Page struct {
	Page       int `json:"page"`
	PerPage    int `json:"page_size"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
	First      int `json:"-"`
	Last       int `json:"-"`
}

type PageV2 struct {
	TotalRows int    `json:"total_rows"`
	TotalPage int    `json:"total_page"`
	Page      int    `json:"page"`
	Data      string `json:"data"`
	HasNext   bool   `json:"has_next"`
	HasPrev   bool   `json:"has_prev"`
	First     int    `json:"-"`
	Last      int    `json:"-"`
}
