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
// @Param request body usecases.VoteInputDTO true "Vote data"
// @Success 201 {object} usecases.VoteOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Router /votes [post]
func (h *VoteHandler) Vote(c *gin.Context) {
	var request usecases.VoteInputDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	output, errs := h.voteFactory.Vote.Execute(request)
	if len(errs) > 0 {
		handleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
