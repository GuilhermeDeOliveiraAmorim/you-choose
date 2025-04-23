package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

type HandlerFactory struct {
	MovieHandler *MovieHandler
	ListHandler  *ListHandler
	VoteHandler  *VoteHandler
	UserHandler  *UserHandler
	BrandHandler *BrandHandler
}

func NewHandlerFactory(inputFactory util.ImputFactory) *HandlerFactory {
	movieFactory := factories.NewMovieFactory(inputFactory)
	listFactory := factories.NewListFactory(inputFactory)
	voteFactory := factories.NewVoteFactory(inputFactory)
	userFactory := factories.NewUserFactory(inputFactory)
	brandFactory := factories.NewBrandFactory(inputFactory)

	return &HandlerFactory{
		MovieHandler: NewMovieHandler(movieFactory),
		ListHandler:  NewListHandler(listFactory),
		VoteHandler:  NewVoteHandler(voteFactory),
		UserHandler:  NewUserHandler(userFactory),
		BrandHandler: NewBrandHandler(brandFactory),
	}
}

func GetAuthenticatedUserID(c *gin.Context) (string, *exceptions.ProblemDetails) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", &exceptions.ProblemDetails{
			Type:     "Unauthorized",
			Title:    "Missing User ID",
			Status:   http.StatusUnauthorized,
			Detail:   "User id is required",
			Instance: exceptions.RFC401,
		}
	}

	userIDStr, ok := userID.(string)
	if !ok || userIDStr == "" {
		return "", &exceptions.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Invalid User ID",
			Status:   http.StatusBadRequest,
			Detail:   "A valid user id is required",
			Instance: exceptions.RFC400,
		}
	}

	return userIDStr, nil
}
