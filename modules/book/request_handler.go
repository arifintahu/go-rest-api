package book

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
	bookRepo := repositories.NewBookRepository(h.db)
	useCase := UseCase{
		book: bookRepo,
	}
	h.controller = Controller{
		useCase: useCase,
	}
	r := router.Group(
		"/books",
	)
	r.POST("/", h.addBook)
	r.GET("/", h.listBooks)
	r.GET("/:id", h.getBook)
	r.PUT("/:id", h.updateBook)
	r.DELETE("/:id", h.deleteBook)
}

func (h *RequestHandler) listBooks(ctx *gin.Context) {
	res, err := h.controller.ListBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) getBook(ctx *gin.Context) {
	res, err := h.controller.GetBook(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.BaseErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *RequestHandler) addBook(ctx *gin.Context) {
	res, err := h.controller.AddBook(ctx)
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
