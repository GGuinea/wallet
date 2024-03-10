package handlers

import (
	"log"
	"main/internal/app"
	"main/internal/handlers/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type withdrawHandler struct {
	walletService app.WalletService
}

func NewWithdrawHandler(walletService app.WalletService) *withdrawHandler {
	return &withdrawHandler{walletService: walletService}
}

func (h *withdrawHandler) ServeHTTP(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		errorResponse := model.ErrorResponseDTO{
			Message: "id is required",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	body := c.Request.Body
	defer body.Close()
	var withdrawRequestDTO model.WithdrawRequestDTO
	if err := c.BindJSON(&withdrawRequestDTO); err != nil {
		log.Println(err)
		errorResponse := model.ErrorResponseDTO{
			Message: "invalid request body",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := h.walletService.Withdraw(id, withdrawRequestDTO.Amount)
	if err != nil {
		log.Println(err)
		errorResponse := model.ErrorResponseDTO{
			Message: "cannot withdraw",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
