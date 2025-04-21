package models

type ErrorResponse struct {
	Error string `json:"Error"`
}

type SuccessResponse struct {
	Message string `json:"Message"`
}

type WarningResponse struct {
	Warning string `json:"Warning"`
}
