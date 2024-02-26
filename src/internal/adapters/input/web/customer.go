package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaias-dgr/ecommerce_service/src/internal/core/domain"
)

func (e ECommerceHandler) GetAllCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	customers, err := e.uCustomers.GetAll(ctx)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, customers)
}

func (e ECommerceHandler) GetbyIDCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) AddPaymentMethodCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	payment, err := e.uPayments.GetAll(ctx, userID)
	if err != nil {
		e.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (e ECommerceHandler) CreateCustomer(c *gin.Context) {
	ctx := c.Request.Context()

	customer := &domain.Customer{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e.log.Error("Fail creatting customer")
	}
	if err := json.Unmarshal(body, customer); err != nil {
		e.log.Errorf("JSON unmarshaling failed: %s", err)
	}

	err = e.uCustomers.Create(ctx, customer)
	if err != nil {
		e.log.Error("Fail creatting customer")
	}
	c.JSON(http.StatusOK, customer)
}

func (e ECommerceHandler) UpdateCustomer(c *gin.Context) {
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
