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
// @Success 201 {object} presenters.SuccessOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Router /signup [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input usecases.CreateUserInputDto
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&input); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "UserHandlerCreateUser",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	output, problems := h.userFactory.CreateUser.Execute(ctx, input)
	if len(problems) > 0 {
		exceptions.HandleErrors(c, problems)
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
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var input usecases.LoginInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CommonErrors", "JsonBindingError"))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.INTERFACE_HANDLERS,
			Code:     exceptions.RFC500_CODE,
			From:     "UserHandlerLogin",
			Message:  "Failed to bind JSON",
			Error:    err,
			Problems: []exceptions.ProblemDetails{problem},
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	output, problems := h.userFactory.Login.Execute(ctx, input)
	if len(problems) > 0 {
		exceptions.HandleErrors(c, problems)
		return
	}

	c.JSON(http.StatusOK, output)
}
