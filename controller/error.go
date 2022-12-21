package controller

// ErrorCode Use type alias to simulate enum
type ErrorCode string

const (
	MissingAuth  = "MISSING_AUTH"
	BadAuth      = "BAD_AUTH"
	JwtSignError = "SIGN_JWT_ERROR"
)

// HttpError The standard error format
type HttpError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}
