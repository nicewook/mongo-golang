package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nicewook/mg/internal/product/dto"
	"github.com/nicewook/mg/internal/product/service"
)

type ProductHandler struct {
	svc service.ProductService
}

func NewProductHandler(e *echo.Echo, svc service.ProductService) {
	handler := &ProductHandler{
		svc: svc,
	}

	e.GET("/", func(c echo.Context) error {
		log.Println("hello")
		return c.String(http.StatusOK, "Hello, World!")
	})

	vOne := e.Group("/v1")
	{
		vOne.GET("", func(c echo.Context) error {
			log.Println("hello. v1")
			return c.String(http.StatusOK, "Hello, World! (API: v1)")
		})
		// GET /

		// POST database/collection - one or many
		// GET /findone - query could be many, get on
		// GET /find - query could be many, get all
		// GET /count - query could be many

		// PUT /product-id/review
		// PUT /product-id/tag

		// DELETE / - query could be many

		vOne.POST("/api/v1/:db/:collection/insertone", handler.InsertOne)
		vOne.GET("/api/v1/:db/:collection/findone", handler.FindOne)
		vOne.GET("/api/v1/:db/:collection/find", handler.FindMany)
	}
	vTwo := e.Group("/v2")
	{
		vTwo.GET("", func(c echo.Context) error {
			log.Println("hello. v2")
			return c.String(http.StatusOK, "Hello, World! (API: v2)")
		})
	}

}

// https://goplay.tools/snippet/epGWQSA2ZCx
func (h *ProductHandler) InsertOne(c echo.Context) error {
	log.Println("insert one")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")

	dtoReq := dto.ProductInsertOneReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}
	var product dto.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dtoReq.Product = product
	dtoResp, err := h.svc.InsertOne(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoResp)
}

func (h *ProductHandler) FindOne(c echo.Context) error {
	log.Println("find one")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")
	productType := c.QueryParam("type")

	dtoReq := dto.ProductFindOneReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		Type:           productType,
	}
	log.Println("dtoReq", dtoReq)

	dtoResp, err := h.svc.FindOne(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoResp)
}

func (h *ProductHandler) FindMany(c echo.Context) error {
	log.Println("find many")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")
	productType := c.QueryParam("type")

	dtoReq := dto.ProductFindManyReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		Type:           productType,
	}
	log.Println("dtoReq", dtoReq)

	dtoResp, err := h.svc.FindMany(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoResp)
}
