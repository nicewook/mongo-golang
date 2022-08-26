package handler

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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

		vOne.POST("/:db/:collection", handler.Insert) // insert one or more documents

		vOne.GET("/:db/:collection/findone", handler.FindOne) // fine one with one or more query params
		vOne.GET("/:db/:collection/find", handler.FindMany)
		vOne.GET("/:db/:collection/count", handler.CountDocuments)

		vOne.PUT("/:db/:collection/:productName/review", handler.AddReview) // https://stackoverflow.com/q/54764101
		vOne.PUT("/:db/:collection/:productName/tag", handler.AddTag)

		vOne.DELETE("/:db/:collection", handler.DeleteDocuments)
	}
	vTwo := e.Group("/v2")
	{
		vTwo.GET("", func(c echo.Context) error {
			log.Println("hello. v2")
			return c.String(http.StatusOK, "Hello, World! (API: v2)")
		})
	}

}

// @Summary     Insert products
// @Description Insert products to the designated database and collection
// @Description You can insert just one product instead of array of product
// @Accept      json
// @Produce     json
// @Param       db          path     string true "Database Name"
// @Param       collection  path     string true "Collection Name"
// @Param       Products   body     []dto.Product true "Array of Product Body"
// @Success     200        {object} dto.ProductInsertResp
// @Failure     200         {object} dto.ErrorResp
// @Failure     400        {string} string "error messages"
// @Router      /v1/{db}/{collection} [post]
func (h *ProductHandler) Insert(c echo.Context) error { // https://goplay.tools/snippet/epGWQSA2ZCx
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

// @Summary     Find one product
// @Description Find one product with a filter from the designated database and collection
// @Description You should have at least one filter
// @Accept      json
// @Produce     json
// @Param       db         path     string true "Database Name"
// @Param       collection path     string true "Collection Name"
// @Param       q          query    string true "Any root field can be used as a filter"
// @Success     200        {object} dto.ProductFindOneResp
// @Failure     200         {object} dto.ErrorResp
// @Router      /v1/{db}/{collection}/findone [get]
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

// @Summary     Find many products
// @Description Find many products with a filter from the designated database and collection
// @Description If there's no filter, it will get all the documents
// @Accept      json
// @Produce     json
// @Param       db         path     string true  "Database Name"
// @Param       collection path     string true  "Collection Name"
// @Param       q          query    string false "Any root field can be used as a filter"
// @Success     200        {object} dto.ProductFindManyResp
// @Failure     200        {object} dto.ErrorResp
// @Router      /v1/{db}/{collection}/find [get]
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

// @Summary     Count documents with a filter
// @Description Count documents with a filter from the designated database and collection
// @Description If there's no filter, it will count all the documents
// @Accept      json
// @Produce     json
// @Param       db         path     string true  "Database Name"
// @Param       collection path     string true  "Collection Name"
// @Param       q          query    string false "Any root field can be used as a filter"
// @Success     200        {object} dto.ProductCountDocumentsResp
// @Failure     200        {object} dto.ErrorResp
// @Router      /v1/{db}/{collection}/count [get]
func (h *ProductHandler) CountDocuments(c echo.Context) error {
	log.Println("count documents")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")
	queryParams := c.QueryParams() // if no query param - it will count all documents

	dtoReq := dto.ProductCountDocumentsReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		QueryParams:    queryParams,
	}
	// log.Println("dtoReq", dtoReq)

	dtoResp, err := h.svc.CountDocuments(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoResp)
}

// @Summary     Add a review to the document
// @Description Add a review to the document only if it is already exist
// @Accept      json
// @Produce     json
// @Param       db          path     string true "Database Name"
// @Param       collection  path     string true "Collection Name"
// @Param       productName path     string true "Product Name"
// @Success     200         {object} dto.ProductAddReviewResp
// @Failure     200        {object} dto.ErrorResp
// @Router      /v1/{db}/{collection}/{productName}/review [put]
func (h *ProductHandler) AddReview(c echo.Context) error {
	log.Println("add review - no duplication")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")
	productName := c.Param("productName")

	var review dto.Review
	if err := c.Bind(&review); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dtoReq := dto.ProductAddReviewReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		ProductName:    productName,
		Review:         review,
	}

	dtoResp, err := h.svc.AddReview(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtoResp)
}

// @Summary     Add a tag to the document
// @Description Add a tag to the document only if it is already exist
// @Accept      json
// @Produce     json
// @Param       db         path     string        true "Database Name"
// @Param       collection path     string        true "Collection Name"
// @Param       productName path     string true "Product Name"
// @Success     200         {object} dto.ProductAddTagResp
// @Failure     200        {object} dto.ErrorResp
// @Router      /v1/{db}/{collection}/{productName}/tag [put]
func (h *ProductHandler) AddTag(c echo.Context) error {
	log.Println("add tag - no duplication")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")
	productName := c.Param("productName")

	tag := struct{ Tag string }{}
	if err := c.Bind(&tag); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dtoReq := dto.ProductAddTagReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		ProductName:    productName,
		Tag:            tag.Tag,
	}

	dtoResp, err := h.svc.AddTag(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtoResp)

}

// @Summary     Delete documents with a filter
// @Description Delete documents with a filter from the designated database and collection
// @Description If there's no filter, it will delete all the documents
// @Accept      json
// @Produce     json
// @Param       db         path     string true  "Database Name"
// @Param       collection path     string true  "Collection Name"
// @Param       q          query    string false "Any root field can be used as a filter"
// @Success     200        {object} dto.ProductDeleteDocumentsResp
// @Failure     200        {object} dto.ErrorResp
// @Router      /v1/{db}/{collection} [delete]
func (h *ProductHandler) DeleteDocuments(c echo.Context) error {
	log.Println("delete documents")
	databaseName := c.Param("db")
	collectionName := c.Param("collection")
	queryParams := c.QueryParams() // if no query param - it will delete all documents

	dtoReq := dto.ProductDeleteDocumentsReq{
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		QueryParams:    queryParams,
	}
	// log.Println("dtoReq", dtoReq)

	dtoResp, err := h.svc.DeleteDocuments(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResp{
			Code:    "E0001",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoResp)
}
