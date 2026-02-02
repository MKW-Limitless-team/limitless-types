package json_response

type Status string

const (
	Failure Status = "failure"
	Success Status = "success"
)

type JsonResponse struct {
	Status Status `json:"status,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func FailureResponse(data string) *JsonResponse {
	return &JsonResponse{Status: Failure, Data: data}
}

func SuccessResponse(data any) *JsonResponse {
	return &JsonResponse{Status: Failure, Data: data}
}
