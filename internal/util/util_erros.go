package util

func NewBadRequestError(title, detail string) ProblemDetails {
	return ProblemDetails{
		Type:     "Bad Request",
		Title:    title,
		Status:   400,
		Detail:   detail,
		Instance: RFC400,
	}
}

func NewUnauthorizedError(title, detail string) ProblemDetails {
	return ProblemDetails{
		Type:     "Unauthorized",
		Title:    title,
		Status:   401,
		Detail:   detail,
		Instance: RFC401,
	}
}

func NewForbiddenError(title, detail string) ProblemDetails {
	return ProblemDetails{
		Type:     "Forbidden",
		Title:    title,
		Status:   403,
		Detail:   detail,
		Instance: RFC403,
	}
}

func NewNotFoundError(title, detail string) ProblemDetails {
	return ProblemDetails{
		Type:     "Not Found",
		Title:    title,
		Status:   404,
		Detail:   detail,
		Instance: RFC404,
	}
}

func NewInternalServerError(title, detail string) ProblemDetails {
	return ProblemDetails{
		Type:     "Internal Server Error",
		Title:    title,
		Status:   500,
		Detail:   detail,
		Instance: RFC500,
	}
}
