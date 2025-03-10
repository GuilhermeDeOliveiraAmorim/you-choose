package util

import (
	"errors"
	"net/http"
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

func NewProblemDetails(t string, title string, status int, detail string, instance string) (*ProblemDetails, error) {
	pd := ProblemDetails{
		Type:     t,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}

	if err := pd.Validate(); err != nil {
		return nil, err
	}

	return &pd, nil
}

func (pd *ProblemDetails) Validate() error {
	if pd.Type == "" || len(pd.Type) > 100 {
		NewLoggerError(
			http.StatusBadRequest,
			"The type must be non-empty and have a maximum of 100 characters",
			"NewProblemDetails",
			"Entities",
			"Error",
		)

		return errors.New("type must be non-empty and have a maximum of 100 characters")
	}

	if pd.Title == "" || len(pd.Title) > 100 {
		NewLoggerError(
			http.StatusBadRequest,
			"The title must be non-empty and have a maximum of 100 characters",
			"NewProblemDetails",
			"Entities",
			"Error",
		)

		return errors.New("title must be non-empty and have a maximum of 100 characters")
	}

	if pd.Status < 100 || pd.Status >= 600 {
		NewLoggerError(
			http.StatusBadRequest,
			"Status must be a valid HTTP code",
			"NewProblemDetails",
			"Entities",
			"Error",
		)

		return errors.New("status must be a valid HTTP code")
	}

	if pd.Detail == "" || len(pd.Detail) > 255 {
		NewLoggerError(
			http.StatusBadRequest,
			"The detail must be non-empty and have a maximum of 255 characters",
			"NewProblemDetails",
			"Entities",
			"Error",
		)

		return errors.New("detail must be non-empty and have a maximum of 255 characters")
	}

	if len(pd.Instance) > 255 {
		NewLoggerError(
			http.StatusBadRequest,
			"The instance must not be longer than 255 characters",
			"NewProblemDetails",
			"Entities",
			"Error",
		)

		return errors.New("instance must not be longer than 255 characters")
	}

	return nil
}
