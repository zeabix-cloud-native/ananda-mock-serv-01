package main

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/zeabix-cloud-native/ananda-mock-serv-01/balance"
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

	log.Info("Connect to " + dsn)
	repo, err := profile.NewMySQLProfileRepository(dsn)

	if err != nil {
		panic(err)
	}

	s := profile.NewProfileService(repo)
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

	router.Run("0.0.0.0:8080")
}
