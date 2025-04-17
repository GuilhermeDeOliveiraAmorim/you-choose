package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userFactory *factories.UserFactory
}

func NewUserHandler(factory *factories.UserFactory) *UserHandler {
	return &UserHandler{
		userFactory: factory,
	}
}

// @Summary Create a new user
// @Description Registers a new user in the system
// @Tags Authentication
// @Accept json
// @Produce json
// @Param CreateUserRequest body usecases.CreateUserInputDto true "User data"
// @Success 201 {object} usecases.CreateUserOutputDto
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Router /signup [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input usecases.CreateUserInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	output, errs := h.userFactory.CreateUser.Execute(input)
	if len(errs) > 0 {
		util.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// @Summary Login a user
// @Description Authenticates a user and returns a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param LoginRequest body usecases.LoginInputDto true "User credentials"
// @Success 200 {object} usecases.LoginOutputDto
// @Failure 400 {object} util.ProblemDetails "Bad Request"
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var input usecases.LoginInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Did not bind JSON",
			Status:   http.StatusBadRequest,
			Detail:   err.Error(),
			Instance: util.RFC400,
		}})
		return
	}

	output, errs := h.userFactory.Login.Execute(c.Request.Context(), input)
	if len(errs) > 0 {
		util.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusOK, output)
}
