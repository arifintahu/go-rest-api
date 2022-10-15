package user

import (
	"log"
	"net/http"

	"github.com/arifintahu/go-rest-api/constants"
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/middlewares"
	"github.com/arifintahu/go-rest-api/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	db         *gorm.DB
	logger     *log.Logger
	controller IController
}

func NewRequestHandler(
	db *gorm.DB,
	logger *log.Logger,
) *RequestHandler {
	return &RequestHandler{
		db:     db,
		logger: logger,
	}
}

func (h *RequestHandler) Handle(router *gin.Engine) {
	userRepo := repositories.NewUserRepository(h.db)
	useCase := UseCase{
		user: userRepo,
	}
	h.controller = Controller{
		useCase: useCase,
	}
	r := router.Group(
		"/users",
		middlewares.Authenticate(),
		middlewares.Authorize(constants.Admin),
	)
	r.POST("/", h.createUser)
	r.GET("/", h.getUsers)
	r.GET("/:id", h.getUserDetail)
	r.PUT("/:id", h.updateUser)
	r.DELETE("/:id", h.deleteUser)
}

func (h *RequestHandler) createUser(ctx *gin.Context) {
	res, err := h.controller.CreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) getUsers(ctx *gin.Context) {
	res, err := h.controller.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) getUserDetail(ctx *gin.Context) {
	res, err := h.controller.GetUserDetail(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) updateUser(ctx *gin.Context) {
	res, err := h.controller.UpdateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) deleteUser(ctx *gin.Context) {
	res, err := h.controller.DeleteUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
