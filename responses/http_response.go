package responses

type HttpResponse struct {
	StatusCode int `json:"status_code,omitempty"`
	Data       any `json:"data,omitempty"`
}

func NewHttpResponse(statusCode int, data any) *HttpResponse {
	return &HttpResponse{
		StatusCode: statusCode,
		Data: data,
	}
}

