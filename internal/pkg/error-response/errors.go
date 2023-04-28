package erroresponse

var ERRORS = map[string]map[string]string{
	"UNPROCESSABLE_ENTITY": {"errorDetail": "unprocessable untity"},
	"PARAM_ERROR":          {"errorDetail": "error on parse param"},
	"INTERNAL_ERROR":       {"errorDetail": "internal server error"},
	"RECORD_NOT_FOUND":     {"errorDetail": "record not found"},
}
