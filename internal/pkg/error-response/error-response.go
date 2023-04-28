package erroresponse

import "fmt"

type ErrorResponse struct {
	ErrorName   string `json:"errorName"`
	ErrorDetail string `json:"errorDetail"`
}

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf("%s, %s", er.ErrorName, er.ErrorDetail)
}

func NewErrorResponse(errorName string) *ErrorResponse {
	errorResponse := &ErrorResponse{}

	if _, ok := ERRORS[errorName]; !ok {
		errorName = "UNPROCESSABLE_ENTITY"
	}

	err := ERRORS[errorName]

	errorResponse.ErrorName = errorName
	errorResponse.ErrorDetail = err["errorDetail"]

	return errorResponse

}
