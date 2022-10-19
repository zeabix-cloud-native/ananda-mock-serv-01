package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/zeabix-cloud-native/ananda-mock-serv-01/balance"
	"github.com/zeabix-cloud-native/ananda-mock-serv-01/clients"
	"github.com/zeabix-cloud-native/ananda-mock-serv-01/health"
	"github.com/zeabix-cloud-native/ananda-mock-serv-01/profile"
)

func main() {

	log.Info("Starting service...")
	dsn := os.Getenv("MYSQL_CONNECTION_STRING")
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	if dsn == "" {
		log.Error("Unable to find $MYSQL_CONNECTION_STRING environment variable")
		panic(errors.New("MYSQL_CONNECTION_STRING not found"))
	}

	accApi := os.Getenv("ACCOUNT_API_ENDPOINT")
	if accApi == "" {
		log.Warn("ACCOUNT_API_ENDPOINT not presented, use http://localhost:8080")
		accApi = "http://localhost:8080"
	}

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		log.Warn("SERVICE_PORT not presented, use 8080 as default")
		port = "8080"
	}
	log.Info(fmt.Sprintf("Run service with port :%s", port))

	key := os.Getenv("SUBSCRIPTION_KEY")
	if key == "" {
		log.Warn("SUBSCRIPTION_KEY not presented, use empty")
		key = ""
	}

	log.Info("Connect to DB")
	repo, err := profile.NewMySQLProfileRepository(dsn)

	if err != nil {
		panic(err)
	}

	accClient := clients.NewAccountService(accApi, key)
	s := profile.NewProfileService(repo, accClient)
	handlers := profile.NewHandlers(s)

	router := gin.Default()
	router.POST("/api/profiles", handlers.CreateProfile)
	router.GET("/api/profiles/:id", handlers.GetProfile)

	// Balance Account

	accRepo, err := balance.NewMySQLBalanceAccountRepository(dsn)
	if err != nil {
		panic(err)
	}

	accService := balance.NewBalanceAccountService(accRepo)
	accHandlers := balance.NewHandler(accService)

	router.POST("/balance/accounts", accHandlers.CreateAccountBalance)
	router.GET("/balance/accounts/:id", accHandlers.GetAccountBalance)
	router.PATCH("/balance/accounts/:id/debit", accHandlers.DebitAccountBalance)

	router.GET("/health", health.Health)

	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
