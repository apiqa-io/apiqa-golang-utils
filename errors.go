package golang_helpers

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RestError struct {
	Code    string     `json:"code"`
	Message FieldError `json:"message"`
}

type RestErrors struct {
	Code    string       `json:"code"`
	Message []FieldError `json:"message"`
}

type UnknownRestError struct {
	Code    string `json:"code"`
	Message error  `json:"message,omitempty"`
}
