package handlers

import (
	"main/internal/app"
	"main/internal/handlers/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getWalletEntriesHandler struct {
	walletService app.WalletService
}

func NewGetWalletEntriesHandler(walletService app.WalletService) *getWalletEntriesHandler {
	return &getWalletEntriesHandler{walletService: walletService}
}

func (h *getWalletEntriesHandler) ServeHTTP(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		errorResponse := model.ErrorResponseDTO{
			Message: "id is required",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	entries, err := h.walletService.GetWalletEntries(id)

	response := model.WalletEntriesResponseDTO{
		Entries: make([]*model.EntryResponseDTO, 0),
	}

	for _, entry := range entries {
		response.Entries = append(response.Entries, &model.EntryResponseDTO{
			ID:           entry.ID,
			WalletID:     entry.WalletID,
			Type:         entry.Type,
			Amount:       entry.Amount.GetAsStringWithDefaultPrecision(),
			BalanceAfter: entry.BalanceAfter.GetAsStringWithDefaultPrecision(),
			CreatedAt:    entry.CreatedAt.String(),
		})
	}

	if err != nil {
		errorResponse := model.ErrorResponseDTO{
			Message: "cannot get entries for wallet",
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	c.JSON(200, response)
}
