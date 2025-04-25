package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	voteFactory *factories.VoteFactory
}

func NewVoteHandler(factory *factories.VoteFactory) *VoteHandler {
	return &VoteHandler{
		voteFactory: factory,
	}
}

// @Summary Create a new vote
// @Description Registers a new vote in the system
// @Tags Votes
// @Accept json
// @Produce json
// @Param request body usecases.Vote true "Vote data"
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /votes [post]
func (h *VoteHandler) Vote(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	var vote usecases.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "VoteHandlerVote",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.VoteInputDTO{
		UserID: userID,
		Vote:   vote,
	}

	output, errs := h.voteFactory.Vote.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
