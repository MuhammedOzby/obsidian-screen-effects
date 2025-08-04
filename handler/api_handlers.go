package handler

import (
	"encoding/json"
	"net/http"
	"obs-go-backend/youtube"
)

// APIHandler api endpointleri için bağımlılıkları tutar
type APIHandler struct {
	YTClient *youtube.Client
}

// LatestSubscriberHandler /api/latest-subscriber endpoint'inin yöneticisidir
func (h *APIHandler) LatestSubscriberHandler(w http.ResponseWriter, r *http.Request) {
	subscriber, err := h.YTClient.GetLatestSubscriber()
	if err != nil {
		http.Error(w, "Could not fetch latest subscriber", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriber)
}
