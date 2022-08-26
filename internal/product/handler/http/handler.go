package handler

import (
	"bytes"
	"io/ioutil"
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

		vOne.POST("/:db/:collection", handler.Insert)         // insert one or more documents
		vOne.GET("/:db/:collection/findone", handler.FindOne) // fine one with one or more query params
		vOne.GET("/:db/:collection/find", handler.FindMany)
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
func (h *ProductHandler) Insert(c echo.Context) error {
	log.Println("insert")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")

	dtoReq := dto.ProductInsertReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}

	var product dto.Product

	// reuse request body: https://www.slll.info/archives/2625.html
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(b))

	// bind product or products
	if err := c.Bind(&product); err == nil {
		log.Println("insert one product")
		dtoReq.Products = append(dtoReq.Products, product)
	} else {
		log.Println("insert many products:", err)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(b))
		if err = c.Bind(&dtoReq.Products); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	dtoResp, err := h.svc.Insert(dtoReq)
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
	queryParams := c.QueryParams() // if no query param - it will not come to this endpoint

	dtoReq := dto.ProductFindOneReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		QueryParams:    queryParams,
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
	queryParams := c.QueryParams() // if no query param - it will not come to this endpoint

	dtoReq := dto.ProductFindManyReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		QueryParams:    queryParams,
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
