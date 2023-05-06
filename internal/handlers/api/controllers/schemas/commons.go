package schemas

type SuccessResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
}
