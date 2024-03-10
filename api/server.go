package api

import (
	"context"
	"fmt"
	"main/config"
	"main/internal/adapters"
	"main/internal/app"
	"main/internal/handlers"
	"main/pkg"
	"main/pkg/postgres/conn"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Start() error
}

type GinHttpServer struct {
	serverConfig *config.ServerConfig
	router       *gin.Engine
}

type GinServerDeps struct {
	Config *config.Config
}

func NewGinHttpServer(deps *GinServerDeps) (*GinHttpServer, error) {
	if deps == nil {
		return nil, fmt.Errorf("deps is nil")
	}
	ctx := context.Background()
	dbPool, err := conn.GetDbPool(ctx, deps.Config.DbConfig)

	router := gin.Default()

	router.GET("/api/v1/health", healthHandler)

	if err != nil {
		return nil, err
	}

	walletRepo, err := adapters.NewWalletPostgresRepo(&adapters.WalletPostgresRepoDeps{ConnPool: dbPool})
	if err != nil {
		return nil, err
	}
	walletService := app.NewWalletService(walletRepo, pkg.NewUUIDGenerator())
	walletFacade := handlers.NewWalletHandlersFacade(walletService)

	intializeWalletHandlers(router, walletFacade)

	return &GinHttpServer{
		serverConfig: deps.Config.ServerConfig,
		router:       router,
	}, nil
}

func (s *GinHttpServer) Start() error {
	return s.router.Run(fmt.Sprintf(":%s", s.serverConfig.Port))
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func intializeWalletHandlers(router *gin.Engine, facade *handlers.WalletHandlersFacade) {
	walletGroup := router.Group("/api/v1/wallets")
	walletGroup.POST("", facade.NewWalletHandler())
	walletGroup.PUT("/:id/balance/deposit", facade.DepositHandler())
	walletGroup.PUT("/:id/balance/withdraw", facade.WithdrawHandler())
	walletGroup.GET("/:id", facade.GetWalletHandler())
}
