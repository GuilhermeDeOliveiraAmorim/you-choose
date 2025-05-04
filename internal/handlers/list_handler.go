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
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists [post]
func (h *ListHandler) CreateList(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	var list usecases.List
	if err := c.ShouldBindJSON(&list); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "ListHandlerCreateList",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.CreateListInputDTO{
		List:   list,
		UserID: userID,
	}

	output, errs := h.listFactory.CreateList.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
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
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists/movies [post]
func (h *ListHandler) AddMoviesList(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	var movies usecases.Movies
	if err := c.ShouldBindJSON(&movies); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "ListHandlerAddMoviesList",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.AddMoviesListInputDTO{
		UserID: userID,
		Movies: movies,
	}

	output, errs := h.listFactory.AddMoviesList.Execute(c.Request.Context(), input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
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
// @Success 200 {object} usecases.GetListByUserIDOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists/users [get]
func (h *ListHandler) GetListByUserID(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	listID := c.Query("list_id")

	input := usecases.GetListByUserIDInputDTO{
		ListID: listID,
		UserID: userID,
	}

	output, errs := h.listFactory.GetListByUserID.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get List
// @Description Get a list of movies and numbers of votes
// @Tags Lists
// @Accept json
// @Produce json
// @Param list_id query string true "List id"
// @Success 200 {object} usecases.GetListByIDOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Router /lists [get]
func (h *ListHandler) GetListByID(c *gin.Context) {
	listID := c.Query("list_id")

	input := usecases.GetListByIDInputDTO{
		ListID: listID,
	}

	output, errs := h.listFactory.GetListByID.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get Lists
// @Description Get all lists
// @Tags Lists
// @Accept json
// @Produce json
// @Success 200 {object} usecases.GetListsOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Router /lists/all [get]
func (h *ListHandler) GetLists(c *gin.Context) {
	input := usecases.GetListsInputDTO{}

	output, errs := h.listFactory.GetLists.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Add brands to list
// @Description Add new brands to list
// @Tags Lists
// @Accept json
// @Produce json
// @Param request body usecases.Brands true "AddBrandsList data"
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /lists/brands [post]
func (h *ListHandler) AddBrandsList(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	var brands usecases.Brands
	if err := c.ShouldBindJSON(&brands); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "ListHandlerAddBrandsList",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.AddBrandsListInputDTO{
		UserID: userID,
		Brands: brands,
	}

	output, errs := h.listFactory.AddBrandsList.Execute(c.Request.Context(), input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// @Summary Sort items by type
// @Description List items sorted by number of votes
// @Tags Items
// @Accept json
// @Produce json
// @Param list_type query string true "List Type (MOVIE or BRAND)"
// @Success 200 {object} usecases.ShowsRankingItemsOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Router /items [get]
func (h *ListHandler) ShowsRankingItems(c *gin.Context) {
	listType := c.Query("list_type")

	input := usecases.ShowsRankingItemsInputDTO{
		ListType: listType,
	}

	output, errs := h.listFactory.ShowsRankingItems.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusOK, output)
}
