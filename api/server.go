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
	GracefulStop(context context.Context) error
}

type GinHttpServer struct {
	serverConfig *config.ServerConfig
	server      *http.Server
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
	if err != nil {
		return nil, err
	}


	walletRepo, err := adapters.NewWalletPostgresRepo(&adapters.WalletPostgresRepoDeps{ConnPool: dbPool})
	if err != nil {
		return nil, err
	}


	walletService := app.NewWalletService(walletRepo, pkg.NewUUIDGenerator())
	walletFacade := handlers.NewWalletHandlersFacade(walletService)
	router := getRouter(walletFacade)

	intializeWalletHandlers(router, walletFacade)

	return &GinHttpServer{
		serverConfig: deps.Config.ServerConfig,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", deps.Config.ServerConfig.Port),
			Handler: router,
		},
	}, nil
}

func getRouter(walletFacade *handlers.WalletHandlersFacade) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/health", healthHandler)
	intializeWalletHandlers(router, walletFacade)
	return router
}

func (s *GinHttpServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *GinHttpServer) GracefulStop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
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
	walletGroup.GET("/:id/balance", facade.GetWalletBalanceHandler())
	walletGroup.GET("/:id", facade.GetWalletHandler())
}
