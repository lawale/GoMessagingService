package responses

type StatusResponse struct {
	ResponseCode ResponseCode `json:"-"`
	Message      string       `json:"message"`
}

type ObjectResponse[T any] struct {
	StatusResponse
	Data T
}

type ListResponse[T any] struct {
	StatusResponse
	Data []T
}

type ResponseCode int

const (
	Success = iota + 1
	BadRequest
	NotFound
	ServiceError
)
