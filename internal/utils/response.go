package utils

type Response struct {
	Message     string `json:"message"`
	Data        any    `json:"data,omitempty"`
	Validations any    `json:"validations,omitempty"`
}

func BuildResponse(message string, data any, validations any) Response {
	res := Response{
		Message:     message,
		Data:        data,
		Validations: validations,
	}

	return res
}
