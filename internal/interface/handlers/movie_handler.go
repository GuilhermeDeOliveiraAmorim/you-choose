package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
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
// @Tags Movies
// @Accept json
// @Produce json
// @Param request body usecases.Movie true "Movie data"
// @Success 201 {object} usecases.CreateMovieOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, gin.H{"error": err})
		return
	}

	var movie usecases.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	input := usecases.CreateMovieInputDTO{
		UserID: userID,
		Movie:  movie,
	}

	output, errs := h.movieFactory.CreateMovie.Execute(input)
	if len(errs) > 0 {
		handleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
