package main

type ListResponse struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
