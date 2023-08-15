package utils

func NewErrorResp(statusCode int32, statusMsg string) *map[string] interface{} {
	return &map[string] interface{}{
		"StatusCode": statusCode,
		"StatusMsg":  statusMsg,
	}
}