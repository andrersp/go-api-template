package erroresponse

import "fmt"

type ErrorResponse struct {
	ErrorName   string `json:"errorName"`
	ErrorDetail string `json:"errorDetail"`
}

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf("%s, %s", er.ErrorName, er.ErrorDetail)
}

func NewErrorResponse(errorName, errorDetail string) *ErrorResponse {
	errorResponse := &ErrorResponse{}

	if _, ok := ERRORS[errorName]; !ok {
		errorName = "UNPROCESSABLE_ENTITY"
	}

	err := ERRORS[errorName]

	errorResponse.ErrorName = errorName

	if errorDetail == "" {
		errorResponse.ErrorDetail = err["errorDetail"]
	} else {
		errorResponse.ErrorDetail = errorDetail
	}

	return errorResponse

}
