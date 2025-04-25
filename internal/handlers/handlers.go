package handlers

import (
	"context"
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/gin-gonic/gin"
)

type HandlerFactory struct {
	MovieHandler *MovieHandler
	ListHandler  *ListHandler
	VoteHandler  *VoteHandler
	UserHandler  *UserHandler
	BrandHandler *BrandHandler
}

func NewHandlerFactory(inputFactory database.StorageInput) *HandlerFactory {
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

func GetAuthenticatedUserID(ctx context.Context, c *gin.Context) (string, []exceptions.ProblemDetails) {
	problems := []exceptions.ProblemDetails{}

	userID, exists := c.Get("userID")
	if !exists {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("CommonErrors", "MissingUserID")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "GetAuthenticatedUserID",
			Message:  "Failed to get user ID from context",
			Error:    errors.New("user ID not found"),
			Problems: problems,
		})

		return "", problems
	}

	userIDStr, ok := userID.(string)
	if !ok || userIDStr == "" {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.BadRequest, language.GetErrorMessage("CommonErrors", "InvalidUserID")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "GetAuthenticatedUserID",
			Message:  "Failed to convert user ID to string",
			Error:    errors.New("user ID conversion error"),
			Problems: problems,
		})

		return "", problems
	}

	return userIDStr, []exceptions.ProblemDetails{}
}
