package userError

type ErrorCode string

// Errors
const (
	ErrorNotFound   = ErrorCode(`not_found`)
	ErrorInternal   = ErrorCode(`internal`)
	ErrorBadRequest = ErrorCode(`bad request`)
	ErrorForbidden  = ErrorCode(`forbidden`)
)

// Error implements the error interface.
func (e ErrorCode) Error() string {
	return string(e)
}
