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

func (h *MeetingHandler) MeetingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateMeetingHandler(w, r) // Обработка создания встречи
	case http.MethodPut:
		h.UpdateMeetingHandler(w, r) // Обработка обновления встречи
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
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

func (h *MeetingHandler) UpdateMeetingHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateMeetingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	// Преобразуем DTO в доменную сущность
	meeting := domain.Meeting{
		ID:             req.ID, // Предполагается, что ID передается в запросе
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

	// Вызов метода сервиса для обновления встречи
	if err := h.service.UpdateMeeting(&meeting); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting) // Отправляем обновленную встречу как ответ
}
