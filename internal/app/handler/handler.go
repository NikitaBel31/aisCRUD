package handler

import (
	"ais/internal/app/service"
	"ais/internal/domain"
	"ais/internal/dto"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var req dto.CreateMeetingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	// Преобразуем DTO в доменную сущность
	meeting := domain.Meeting{
		Name:           req.Name,
		Place:          req.Place,
		Comment:        req.Comment,
		ApplicantID:    req.ApplicantID,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		IsFullDay:      req.IsFullDay,
		IsOnline:       req.IsOnline,
		IsOutlookEvent: req.IsOutlookEvent,
		EmailAuthor:    req.EmailAuthor,
		IsViewMeeting:  req.IsViewMeeting,
		IsStartMeeting: req.IsStartMeeting,
	}

	// Вызов метода сервиса для создания встречи
	if err := h.service.CreateMeeting(&meeting); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Подготовка ответа
	response := dto.MeetingResponse{
		Name:        meeting.Name,
		Place:       meeting.Place,
		StartDate:   meeting.StartDate,
		EndDate:     meeting.EndDate,
		IsFullDay:   meeting.IsFullDay,
		IsOnline:    meeting.IsOnline,
		EmailAuthor: meeting.EmailAuthor,
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response) // Отправляем созданную встречу как ответ
}
