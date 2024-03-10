package handlers

import (
	"main/internal/app"
	"main/internal/handlers/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getWalletHandler struct {
	walletService app.WalletService
}

func NewGetWalletHandler(walletService app.WalletService) *getWalletHandler {
	return &getWalletHandler{walletService: walletService}
}

func (h *getWalletHandler) ServeHTTP(c *gin.Context) {
	walletID := c.Param("id")
	wallet, err := h.walletService.GetWalletByID(walletID)
	if err != nil {
		errorResponse := model.ErrorResponseDTO{
			Message: "cannot get wallet",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := model.WalletResponseDTO{
		ID:      wallet.ID,
		Balance: wallet.Balance.GetAsStringWithDefaultPrecision(),
		CreatedAt: wallet.CreatedAt.String(),
		UpdatedAt: wallet.UpdatedAt.String(),

	}

	c.JSON(http.StatusOK, response)
}
