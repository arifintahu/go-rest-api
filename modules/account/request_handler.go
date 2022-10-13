package account

import (
	"log"
	"net/http"

	"github.com/arifintahu/go-rest-api/dto"
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
		"/account",
	)
	r.POST("/login", h.login)
}

func (h *RequestHandler) login(ctx *gin.Context) {
	res, err := h.controller.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
