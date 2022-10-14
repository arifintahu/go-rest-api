package book

import (
	"log"
	"net/http"

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
	bookRepo := repositories.NewBookRepository(h.db)
	useCase := UseCase{
		book: bookRepo,
	}
	h.controller = Controller{
		useCase: useCase,
	}
	r := router.Group(
		"/books",
		middlewares.Authenticate(),
	)
	r.POST("/", h.createBook)
	r.GET("/", h.getBooks)
	r.GET("/:id", h.getBookDetail)
	r.PUT("/:id", h.updateBook)
	r.DELETE("/:id", h.deleteBook)
}

func (h *RequestHandler) createBook(ctx *gin.Context) {
	res, err := h.controller.CreateBook(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) getBooks(ctx *gin.Context) {
	res, err := h.controller.GetBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) getBookDetail(ctx *gin.Context) {
	res, err := h.controller.GetBookDetail(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) updateBook(ctx *gin.Context) {
	res, err := h.controller.UpdateBook(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) deleteBook(ctx *gin.Context) {
	res, err := h.controller.DeleteBook(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
