package web

import (
	"github.com/gin-gonic/gin"
	useCase "github.com/isaias-dgr/ecommerce_service/src/internal/core/usecase"
	log "github.com/sirupsen/logrus"
)

type ECommerceHandler struct {
	uPayments  *useCase.Payments
	uProducts  *useCase.Products
	uCustomers *useCase.Customers
	uMerchants *useCase.Merchants
	uCarts     *useCase.Carts
	log        *log.Logger
}

func NewECommerceHandler(
	p *useCase.Payments,
	pr *useCase.Products,
	c *useCase.Customers,
	m *useCase.Merchants,
	cr *useCase.Carts,
	l *log.Logger) *ECommerceHandler {

	return &ECommerceHandler{
		uPayments:  p,
		uProducts:  pr,
		uCustomers: c,
		uMerchants: m,
		uCarts:     cr,
		log:        l,
	}
}

func (e ECommerceHandler) getPaymentsRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/products", e.GetAllProduct)
	router.POST("/products", e.CreateProduct)
	router.PUT("/products/:productID", e.UpdateProduct)

	router.GET("/carts", e.GetAllCart)
	router.POST("/carts", e.CreateCart)
	router.GET("/carts/:cartID", e.GetbyIDCart)
	router.PUT("/carts/:cartID/add_item", e.AddItemCart)

	router.GET("/customers", e.GetAllCustomer)
	router.POST("/customers", e.CreateCustomer)
	router.PUT("/customers/:customerID", e.UpdateCustomer)
	router.PUT("/customers/:customerID/add_payment_method", e.AddPaymentMethodCustomer)

	router.GET("/merchants", e.GetAllMerchant)
	router.POST("/merchants", e.CreateMerchant)
	router.POST("/merchants/:merchantsID", e.GetbyIDMerchant)
	router.PUT("/merchants/:merchantsID/add_payment_method", e.AddPaymentMethodMerchant)

	router.POST("/payments", e.CreatePayment)
	router.GET("/payments/:paymentID", e.GetByIDPayment)
	router.GET("/payments/:paymentID/refund", e.RefundPayment)
	return router
}

func (e ECommerceHandler) Execute() {
	e.log.Info("Service Payments running")
	routes := e.getPaymentsRoutes()
	routes.Run(":8080")
}
