package util

type ErrorType string

const (
	BadRequest          ErrorType = "Bad Request"
	Unauthorized        ErrorType = "Unauthorized"
	Forbidden           ErrorType = "Forbidden"
	NotFound            ErrorType = "Not Found"
	InternalServerError ErrorType = "Internal Server Error"
)

type ProblemDetailsOutputDTO struct {
	ProblemDetails []ProblemDetails `json:"problem_details"`
}

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance,omitempty"`
}

func NewProblemDetails(errorType ErrorType, msg ErrorMessage) ProblemDetails {
	var status int
	var instance string

	switch errorType {
	case BadRequest:
		status = 400
		instance = RFC400
	case Unauthorized:
		status = 401
		instance = RFC401
	case Forbidden:
		status = 403
		instance = RFC403
	case NotFound:
		status = 404
		instance = RFC404
	case InternalServerError:
		status = 500
		instance = RFC500
	default:
		status = 500
		errorType = InternalServerError
		instance = RFC500
	}

	return ProblemDetails{
		Type:     string(errorType),
		Title:    msg.Title,
		Status:   status,
		Detail:   msg.Detail,
		Instance: instance,
	}
}
