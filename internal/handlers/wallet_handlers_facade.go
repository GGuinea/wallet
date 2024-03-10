package handlers

import (
	"main/internal/app"

	"github.com/gin-gonic/gin"
)

type WalletHandlersFacade struct {
	depositHandler   *depositHandler
	withdrawHandler  *withdrawHandler
	getWalletBalanceHandler *getWalletBalanceHandler
	newWalletHandler *newWalletHandler
	getWalletHandler *getWalletHandler
}

func NewWalletHandlersFacade(walletService app.WalletService) *WalletHandlersFacade {
	return &WalletHandlersFacade{
		depositHandler:   NewDepositHandler(walletService),
		withdrawHandler:  NewWithdrawHandler(walletService),
		getWalletBalanceHandler: NewGetWalletBalanceHandler(walletService),
		newWalletHandler: NewNewWalletHandler(walletService),
		getWalletHandler: NewGetWalletHandler(walletService),
	}
}

func (f *WalletHandlersFacade) DepositHandler() func(c *gin.Context) {
	return f.depositHandler.ServeHTTP
}

func (f *WalletHandlersFacade) WithdrawHandler() func(c *gin.Context) {
	return f.withdrawHandler.ServeHTTP
}

func (f *WalletHandlersFacade) GetWalletBalanceHandler() func(c *gin.Context) {
	return f.getWalletBalanceHandler.ServeHTTP
}

func (f *WalletHandlersFacade) NewWalletHandler() func(c *gin.Context) {
	return f.newWalletHandler.ServeHTTP
}

func (f *WalletHandlersFacade) GetWalletHandler() func(c *gin.Context) {
	return f.getWalletHandler.ServeHTTP
}
