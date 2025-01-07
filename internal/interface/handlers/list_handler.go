package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	listFactory *factories.ListFactory
}

func NewListHandler(factory *factories.ListFactory) *ListHandler {
	return &ListHandler{
		listFactory: factory,
	}
}

// @Summary Create a new list
// @Description Registers a new list in the system
// @Tags Lists
// @Accept json
// @Produce json
// @Param request body usecases.CreateListInputDTO true "Movie data"
// @Success 201 {object} usecases.CreateListOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Router /lists [post]
func (h *ListHandler) CreateList(c *gin.Context) {
	var request usecases.CreateListInputDTO
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

	output, errs := h.listFactory.CreateList.Execute(request)
	if len(errs) > 0 {
		handleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// @Summary Add movies to list
// @Description Add new movies to list
// @Tags Lists
// @Accept json
// @Produce json
// @Param request body usecases.AddMoviesListInputDTO true "AddMoviesList data"
// @Success 201 {object} usecases.AddMoviesListOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Router /lists/movies [post]
func (h *ListHandler) AddMoviesList(c *gin.Context) {
	var request usecases.AddMoviesListInputDTO
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

	output, errs := h.listFactory.AddMoviesList.Execute(request)
	if len(errs) > 0 {
		handleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// @Summary Get List
// @Description Get a list of movies and votes
// @Tags Lists
// @Accept json
// @Produce json
// @Param list_id query string true "List id"
// @Success 201 {object} usecases.GetListByIDOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists [get]
func (h *ListHandler) GetListByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, gin.H{"error": err})
		return
	}

	listID := c.Query("list_id")

	input := usecases.GetListByIDInputDTO{
		ListID: listID,
		UserID: userID,
	}

	output, errs := h.listFactory.GetListByID.Execute(input)
	if len(errs) > 0 {
		handleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
