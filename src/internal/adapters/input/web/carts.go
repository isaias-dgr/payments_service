package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
)

func (e ECommerceHandler) GetAllCart(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) GetbyIDCart(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) AddItemCart(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) CreateCart(c *gin.Context) {
	ctx := c.Request.Context()

	carts := &domain.ShoppingCart{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e.log.Error("Fail creatting carts")
	}
	if err := json.Unmarshal(body, carts); err != nil {
		e.log.Errorf("JSON unmarshaling failed: %s", err)
	}

	err = e.uCarts.Create(ctx, carts)
	if err != nil {
		e.log.Error("Fail creatting carts")
	}
	c.JSON(http.StatusOK, carts)
}

func (e ECommerceHandler) UpdateCart(c *gin.Context) {
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
