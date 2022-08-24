package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nicewok/mg/internal/product/dto"
	"github.com/nicewok/mg/internal/product/service"
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

	// e.POST("/api/v1", handler.InsertOne)
	e.POST("/api/v1/:db/:collection/insertone", handler.InsertOne)
}

// https://goplay.tools/snippet/epGWQSA2ZCx
func (h *ProductHandler) InsertOne(c echo.Context) error {
	log.Println("insert one")
	// databaseName := "db"
	databaseName := c.Param("db")
	// collectionName := "collection"
	collectionName := c.Param("collection")

	dtoReq := dto.ProductInsertReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}
	var product dto.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dtoReq.Product = product
	dtoResp := h.svc.InsertOne(dtoReq)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	return c.JSON(http.StatusOK, dtoResp)
}
