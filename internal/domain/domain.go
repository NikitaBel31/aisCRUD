package domain

import (
	"errors"
	"strings"
	"time"
)

type Meeting struct {
	//ID             string    // Идентификатор встречи
	Name           string    // Название встречи
	Place          string    // Место проведения
	Comment        string    // Комментарий
	ApplicantID    string    // Соискатель
	StartDate      time.Time // Дата начала
	StartTime      time.Time // Время начала
	EndDate        time.Time // Дата окончания
	EndTime        time.Time // Время окончания
	IsFullDay      bool      // Весь день
	IsOnline       bool      // Онлайн встреча
	IsOutlookEvent bool      // Создать событие в Outlook
	EmailAuthor    string    // Email организатора
	IsViewMeeting  bool      // Признак просмотра за 15 минут
	IsStartMeeting bool      // Признак начала встречи
}

func (m *Meeting) CheckMeetingNameNotWhitespace() error {
	if strings.TrimSpace(m.Name) == "" {
		return errors.New(`введено недопустимое значение поля "Название встречи"`)
	}
	return nil
}

func (m *Meeting) CheckMeetingPlaceNotWhitespace() error {
	if strings.TrimSpace(m.Place) == "" {
		return errors.New(` "неверное значение, ссылка должна начинаться с http"`)
	}
	return nil
}

func (m *Meeting) CheckCommentNotWhitespace() error {
	if strings.TrimSpace(m.Place) == "" {
		return errors.New(`введено недопустимое значение поля "Комментарий"`)
	}
	return nil
}

func (m *Meeting) CheckApplicantIDNotWhitespace() error {
	if strings.TrimSpace(m.Place) == "" {
		return errors.New(`введено недопустимое значение поля "Соискатель"`)
	}
	return nil
}

func (m *Meeting) CheckStartDateNotEmpty() error {
	if m.StartDate.IsZero() {
		return errors.New("дата начала обязательна для заполнения")
	}
	return nil
}

func (m *Meeting) CheckEndDateNotEmpty() error {
	if m.EndDate.IsZero() {
		return errors.New("дата конца обязательна для заполнения")
	}
	return nil
}

func (m *Meeting) CheckStartDateMinTime() error {
	minTime := time.Date(m.StartDate.Year(), m.StartDate.Month(), m.StartDate.Day(), 8, 0, 0, 0, m.StartDate.Location())
	if m.StartDate.Before(minTime) {
		return errors.New("минимальное время начала встречи 08:00")
	}
	return nil
}

func (m *Meeting) CheckEndDateMaxTime() error {
	maxTime := time.Date(m.EndDate.Year(), m.EndDate.Month(), m.EndDate.Day(), 19, 0, 0, 0, m.EndDate.Location())
	if m.EndDate.After(maxTime) {
		return errors.New("максимальное время окончания встречи 19:00")
	}
	return nil
}
