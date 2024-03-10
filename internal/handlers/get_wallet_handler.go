package handlers

import (
	"log"
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
	id := c.Param("id")
	if id == "" {
		errorResponse := model.ErrorResponseDTO{
			Message: "id is required",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	balance, err := h.walletService.GetBalance(id)
	if err != nil {
		log.Println(err)
		errorResponse := model.ErrorResponseDTO{
			Message: "cannot get balance",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
