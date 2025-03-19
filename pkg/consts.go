package pkg

var StatusCode = map[string]int{
	"SUCCESS":        200,
	"CREATED":        201,
	"NOT_FOUND":      404,
	"BAD_REQUEST":    400,
	"INTERNAL_ERROR": 500,
}

type ValidationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
