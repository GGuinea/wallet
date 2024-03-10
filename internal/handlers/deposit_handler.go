package handlers

import (
	"main/internal/app"
	"main/internal/handlers/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type depositHandler struct {
	walletService app.WalletService
}

func NewDepositHandler(walletService app.WalletService) *depositHandler {
	return &depositHandler{walletService: walletService}
}

func (h *depositHandler) ServeHTTP(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}
	body := c.Request.Body
	defer body.Close()
	var depositRequest model.DepositRequestDTO

	if err := c.BindJSON(&depositRequest); err != nil {
		errorResponse := model.ErrorResponseDTO{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := h.walletService.Deposit(id, depositRequest.Amount)
	if err != nil {
		errorResponse := model.ErrorResponseDTO{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
