package utils

type Response struct {
	Message    any `json:"message,omitempty"`
	Data       any `json:"data,omitempty"`
	Validation any `json:"validation,omitempty"`
}

func BuildResponse(message any, data any, validation any) Response {
	res := Response{
		Message:    message,
		Data:       data,
		Validation: validation,
	}

	return res
}
