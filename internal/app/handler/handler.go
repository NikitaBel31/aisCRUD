package handler

import (
	"ais/internal/app/service"
	"ais/internal/domain"
	"encoding/json"
	"net/http"
)

type MeetingHandler struct {
	service *service.MeetingService
}

func NewMeetingHandler(service *service.MeetingService) *MeetingHandler {
	return &MeetingHandler{service: service}
}

func (h *MeetingHandler) CreateMeetingHandler(w http.ResponseWriter, r *http.Request) {

	var meeting domain.Meeting
	err := json.NewDecoder(r.Body).Decode(&meeting)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.service.CreateMeeting(&meeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//После тестов убрать
	// Установка заголовка ответа (чтобы клиент знал, что это JSON)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meeting)
}
