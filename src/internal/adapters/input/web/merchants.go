package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
)

func (e ECommerceHandler) GetAllMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) GetbyIDMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) AddPaymentMethodMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) CreateMerchant(c *gin.Context) {
	ctx := c.Request.Context()

	merchant := &domain.Merchant{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e.log.Error("Fail creatting merchant")
	}
	if err := json.Unmarshal(body, merchant); err != nil {
		e.log.Errorf("JSON unmarshaling failed: %s", err)
	}

	err = e.uMerchants.Create(ctx, merchant)
	if err != nil {
		e.log.Error("Fail creatting merchant")
	}
	c.JSON(http.StatusOK, merchant)
}

func (e ECommerceHandler) UpdateMerchant(c *gin.Context) {
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
