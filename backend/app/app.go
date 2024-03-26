package app

const (
	storeErrorStutas = 450
)

type status string

const (
	Success status = "success"
	Fail    status = "error"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}

type Error Response

func (err *Error) Error() string {
	return err.Message
}
