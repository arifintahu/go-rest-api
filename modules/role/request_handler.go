package role

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
	db         	*gorm.DB
	logger 		*log.Logger
	controller 	IController
}

func NewRequestHandler(
	db         *gorm.DB,
	logger *log.Logger,
) *RequestHandler {
	return &RequestHandler{
		db:       db,
		logger: logger,
	}
}

func (h *RequestHandler) Handle(router *gin.Engine) {
	roleRepo := repositories.NewRoleRepository(h.db)
	useCase := UseCase{
		role:      roleRepo,
	}
	h.controller = Controller{
		useCase: useCase,
	}
	r := router.Group(
		"/roles",
		middlewares.Authenticate(),
		middlewares.Authorize(constants.Admin),
	)
	r.POST("/", h.createRole)
	r.GET("/",h.getRoles)
}

func (h *RequestHandler) createRole(ctx *gin.Context) {
	res, err := h.controller.CreateRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) getRoles(ctx *gin.Context) {
	res, err := h.controller.GetRoles(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
