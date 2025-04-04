package util

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

func NewProblemDetails(t string, title string, status int, detail string, instance string) *ProblemDetails {
	return &ProblemDetails{
		Type:     t,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}
