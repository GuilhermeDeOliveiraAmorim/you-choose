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
// @Param request body usecases.List true "List data"
// @Success 201 {object} usecases.CreateListOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists [post]
func (h *ListHandler) CreateList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, gin.H{"error": err})
		return
	}

	var list usecases.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	input := usecases.CreateListInputDTO{
		List:   list,
		UserID: userID,
	}

	output, errs := h.listFactory.CreateList.Execute(input)
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
// @Param request body usecases.Movies true "AddMoviesList data"
// @Success 201 {object} usecases.AddMoviesListOutputDTO
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Failure 500 {object} util.ProblemDetails "Internal Server Error"
// @Failure 401 {object} util.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists/movies [post]
func (h *ListHandler) AddMoviesList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, gin.H{"error": err})
		return
	}

	var movies usecases.Movies
	if err := c.ShouldBindJSON(&movies); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	input := usecases.AddMoviesListInputDTO{
		UserID: userID,
		Movies: movies,
	}

	output, errs := h.listFactory.AddMoviesList.Execute(input)
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
