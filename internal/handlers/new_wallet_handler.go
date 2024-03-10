package handlers

import (
	"log"
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
		log.Println(err)
		errorResponse := model.ErrorResponseDTO{
			Message: "cannot create wallet",
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := model.NewWalletResponseDTO{
		ID:      wallet.ID,
		Balance: wallet.Balance.GetAsStringWithDefaultPrecision(),
	}

	c.JSON(http.StatusCreated, response)
}
