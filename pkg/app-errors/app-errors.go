package apperrors

type AppError struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
}

func (err *AppError) Error() string {
	return err.Msg
}

func NewAppError(msg string) *AppError {
	return &AppError{
		Msg:     msg,
		Success: false,
	}
}
