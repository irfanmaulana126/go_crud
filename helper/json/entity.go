package jsonHelper

import paginationHelper "belajar/helper/pagination"

type response struct {
	Level   string      `json:"level,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type responseV2 struct {
	Status     bool                     `json:"status"`
	StatusCode int                      `json:"status_code"`
	Message    interface{}              `json:"message"`
	Code       string                   `json:"code"`
	Pagination *paginationHelper.PageV2 `json:"pagination,omitempty"`
	Data       interface{}              `json:"data,omitempty"`
}
type responseV3 struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status     bool                   `json:"-"`
	Code       string                 `json:"code"`
	Message    interface{}            `json:"message"`
	Pagination *paginationHelper.Page `json:"pagination,omitempty"`
}
