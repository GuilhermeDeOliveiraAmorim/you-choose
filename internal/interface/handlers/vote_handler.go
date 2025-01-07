package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
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
// @Success 201 {object} usecases.VoteOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /votes [post]
func (h *VoteHandler) Vote(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, gin.H{"error": err})
		return
	}

	var vote usecases.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	input := usecases.VoteInputDTO{
		UserID: userID,
		Vote:   vote,
	}

	output, errs := h.voteFactory.Vote.Execute(input)
	if len(errs) > 0 {
		handleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
