package utils

type Response struct {
	Message     any `json:"message,omitempty"`
	Data        any `json:"data,omitempty"`
	Validations any `json:"validations,omitempty"`
}

func BuildResponse(message any, data any, validations any) Response {
	res := Response{
		Message:     message,
		Data:        data,
		Validations: validations,
	}

	return res
}
