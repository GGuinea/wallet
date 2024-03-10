package handlers

import (
	"main/internal/app"

	"github.com/gin-gonic/gin"
)

type WalletHandlersFacade struct {
	depositHandler   *depositHandler
	withdrawHandler  *withdrawHandler
	getWalletHandler *getWalletHandler
	newWalletHandler *newWalletHandler
}

func NewWalletHandlersFacade(walletService app.WalletService) *WalletHandlersFacade {
	return &WalletHandlersFacade{
		depositHandler:   NewDepositHandler(walletService),
		withdrawHandler:  NewWithdrawHandler(walletService),
		getWalletHandler: NewGetWalletHandler(walletService),
		newWalletHandler: NewNewWalletHandler(walletService),
	}
}

func (f *WalletHandlersFacade) DepositHandler() func(c *gin.Context) {
	return f.depositHandler.ServeHTTP
}

func (f *WalletHandlersFacade) WithdrawHandler() func(c *gin.Context) {
	return f.withdrawHandler.ServeHTTP
}

func (f *WalletHandlersFacade) GetWalletHandler() func(c *gin.Context) {
	return f.getWalletHandler.ServeHTTP
}

func (f *WalletHandlersFacade) NewWalletHandler() func(c *gin.Context) {
	return f.newWalletHandler.ServeHTTP
}
