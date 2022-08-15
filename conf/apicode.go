package conf

const (
	// API error codes
	ApiCodeSuccess          int32 = 200
	ApiCodeParamErr         int32 = 401
	ApiCodeNoAuth           int32 = 403
	ApiCodeNotFound         int32 = 404
	ApiCodeMethodNotAllowed int32 = 405 // Method Not Allowed
	ApiCodeErrMsg           int32 = 500
)
