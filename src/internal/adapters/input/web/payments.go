package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	useCase "github.com/isaias-dgr/ecommerce_service/src/internal/core/usecase"
	log "github.com/sirupsen/logrus"
)

type PaymentsHandler struct {
	uPayments *useCase.Payments
	log       *log.Logger
}

func NewPaymentsHandler(u *useCase.Payments, l *log.Logger) *PaymentsHandler {
	return &PaymentsHandler{
		uPayments: u,
		log:       l,
	}
}

func (p PaymentsHandler) getPaymentsRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/payments", p.GetAll)
	router.POST("/payments", p.Create)
	return router
}

func (p PaymentsHandler) GetAll(c *gin.Context) {
	payment, err := p.uPayments.GetAll()
	if err != nil {
		p.log.Error("Fail getting all payments")
	}
	c.JSON(http.StatusOK, payment)
}

func (p PaymentsHandler) Create(c *gin.Context) {
	payment, err := p.uPayments.Create()
	if err != nil {
		p.log.Error("Fail creatting payment")
	}
	c.JSON(http.StatusOK, payment)
}

func (p PaymentsHandler) Execute() {
	p.log.Info("Service Payments running")
	routes := p.getPaymentsRoutes()
	routes.Run(":8080")
}
