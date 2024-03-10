package handlers

import (
	"main/internal/app"
	"main/internal/handlers/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newWalletHandler struct {
	walletService app.WalletService
}

func NewNewWalletHandler(walletService app.WalletService) *newWalletHandler {
	return &newWalletHandler{walletService: walletService}
}

func (h *newWalletHandler) ServeHTTP(c *gin.Context) {
	wallet, err := h.walletService.NewWallet()
	if err != nil {
		errorResponse := model.ErrorResponseDTO{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := model.NewWalletResponseDTO{
		ID:      wallet.ID,
		Balance: wallet.Balance.GetAsStringWithDefaultPrecision(),
	}

	c.JSON(http.StatusOK, response)
}
