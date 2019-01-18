package changenow

// RequestError embeds in all responses
// Some responses from ChangeNow have an "error" instead.
type RequestError struct {
	Message string `json:"error"`
}

func (e *RequestError) Error() string {
	return e.Message
}

func (e *RequestError) IsError() bool {
	return e.Message != ""
}
