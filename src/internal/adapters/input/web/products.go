package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
)

func (e ECommerceHandler) GetAllProduct(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()

	product := &domain.Product{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e.log.Error("Fail creatting Product")
	}
	if err := json.Unmarshal(body, product); err != nil {
		e.log.Errorf("JSON unmarshaling failed: %s", err)
	}

	err = e.uProducts.Create(ctx, product)
	if err != nil {
		e.log.Error("Fail creatting Product")
	}
	c.JSON(http.StatusOK, product)
}

func (e ECommerceHandler) UpdateProduct(c *gin.Context) {
	ctx := c.Request.Context()

	payment := &domain.Payment{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e.log.Error("Fail creatting payment")
	}
	if err := json.Unmarshal(body, payment); err != nil {
		e.log.Errorf("JSON unmarshaling failed: %s", err)
	}

	err = e.uPayments.Create(ctx, payment)
	if err != nil {
		e.log.Error("Fail creatting payment")
	}
	c.JSON(http.StatusOK, payment)
}
