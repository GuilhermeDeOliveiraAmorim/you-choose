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

type MovieHandler struct {
	movieFactory *factories.MovieFactory
}

func NewMovieHandler(factory *factories.MovieFactory) *MovieHandler {
	return &MovieHandler{
		movieFactory: factory,
	}
}

// @Summary Create a new movie
// @Description Registers a new movie in the system
// @Tags Items
// @Accept json
// @Produce json
// @Param request body usecases.Movie true "Movie data"
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /items/movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	var movie usecases.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "MovieHandlerCreateMovie",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.CreateMovieInputDTO{
		UserID: userID,
		Movie:  movie,
	}

	output, errs := h.movieFactory.CreateMovie.Execute(ctx, input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
