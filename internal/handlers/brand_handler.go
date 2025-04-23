package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
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
	userID, err := GetAuthenticatedUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, gin.H{"error": err})
		return
	}

	var brand usecases.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": exceptions.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: exceptions.RFC400,
		}})
		return
	}

	input := usecases.CreateBrandInputDTO{
		UserID: userID,
		Brand:  brand,
	}

	output, errs := h.brandFactory.CreateBrand.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
