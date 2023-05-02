package dto

type ErrorResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
}
