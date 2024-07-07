package handler

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func newSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{Message: message}
}

func newErrorResponse(err string) ErrorResponse {
	return ErrorResponse{Error: err}
}
