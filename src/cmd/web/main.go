package main

import (
	"database/sql"
	"fmt"
	"os"

	web "github.com/isaias-dgr/ecommerce_service/src/internal/adapters/input/web"
	repo "github.com/isaias-dgr/ecommerce_service/src/internal/adapters/output/repository/psl"
	useCase "github.com/isaias-dgr/ecommerce_service/src/internal/core/usecase"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func SetUpLog() *log.Logger {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func SetClinetDB() *log.Logger {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func SetUpRepository(logger *log.Logger) *sql.DB {
	logger.Info("💾 Set up Database.")
	logger.Info(os.Getenv("POSTGRES_USER"))
	connection := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	logger.Info(connection)
	dbConn, err := sql.Open(`postgres`, connection)
	if err != nil {
		logger.Error(err)
	}
	err = dbConn.Ping()
	if err != nil {
		logger.Error(err)
	}
	return dbConn
}

func main() {
	log := SetUpLog()
	msg := fmt.Sprintf(
		"🤓 SetUp cmd %s_%s:%s..",
		os.Getenv("PROJ_NAME"),
		os.Getenv("PROJ_ENV"),
		"8080")
	log.Info(msg)
	log.Info("🚀 API V1.")
	dbConn := SetUpRepository(log)
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	ucPayments := useCase.NewPayments(log, repo.NewPayments(dbConn, log))
	ucProducts := useCase.NewProducts(log, repo.NewProducts(dbConn, log))
	ucCustomers := useCase.NewCustomers(log, repo.NewCustomers(dbConn, log))
	usMerchants := useCase.NewMerchants(log, repo.NewMerchants(dbConn, log))
	usCarts := useCase.NewCarts(log, repo.NewCarts(dbConn, log))
	webHandler := web.NewECommerceHandler(
		ucPayments,
		ucProducts,
		ucCustomers,
		usMerchants,
		usCarts,
		log,
	)
	webHandler.Execute()
}
