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

type BrandHandler struct {
	brandFactory *factories.BrandFactory
}

func NewBrandHandler(factory *factories.BrandFactory) *BrandHandler {
	return &BrandHandler{
		brandFactory: factory,
	}
}

// @Summary Create a new brand
// @Description Registers a new brand in the system
// @Tags Items
// @Accept json
// @Produce json
// @Param request body usecases.Brand true "Brand data"
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /items/brands [post]
func (h *BrandHandler) CreateBrand(c *gin.Context) {
	ctx := c.Request.Context()

	userID, problem := GetAuthenticatedUserID(ctx, c)
	if len(problem) > 0 {
		c.AbortWithStatusJSON(problem[0].Status, gin.H{"error": problem})
		return
	}

	var brand usecases.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "BrandHandlerCreateBrand",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.CreateBrandInputDTO{
		UserID: userID,
		Brand:  brand,
	}

	output, errs := h.brandFactory.CreateBrand.Execute(ctx, input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
