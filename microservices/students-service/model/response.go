package model

type BaseResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
