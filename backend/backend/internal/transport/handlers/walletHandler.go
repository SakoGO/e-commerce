package handlers

import (
	"e-commerce/backend/internal/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type WalletService interface {
	WalletBalance(balance float64) (*model.Wallet, error)
}

func (h *Handler) WalletBalance(w http.ResponseWriter, r *http.Request) {
	var wallet model.Wallet
	balance, err := h.WalletService.WalletBalance(wallet.Balance)
	if err != nil {
		log.Error().Err(err).Msg("Error when user try to get balance info")
		http.Error(w, "Unable to get wallet balance", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
