package domain

import (
	"errors"
	"strings"
	"time"
)

type Meeting struct {
	ID             string    // Идентификатор встречи
	Name           string    // Название встречи
	Place          string    // Место проведения
	Comment        string    // Комментарий
	ApplicantID    string    // Соискатель
	StartDate      time.Time // Дата начала
	EndDate        time.Time // Дата окончания
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
	if strings.TrimSpace(m.Comment) == "" {
		return errors.New(`введено недопустимое значение поля "Комментарий"`)
	}
	return nil
}

func (m *Meeting) CheckApplicantIDNotWhitespace() error {
	if strings.TrimSpace(m.ApplicantID) == "" {
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

// Проверка на наличие даты и времени конца встречи
func (m *Meeting) CheckEndDateNotEmpty() error {
	if m.EndDate.IsZero() {
		return errors.New("дата окончания обязательна для заполнения")
	}
	return nil
}

// Проверка минимального времени начала встречи (не раньше 8:00)
func (m *Meeting) CheckStartDateMinTime() error {
	if err := m.CheckStartDateNotEmpty(); err != nil {
		return err
	}

	if m.StartDate.Hour() < 8 {
		return errors.New("время начала встречи не может быть раньше 08:00")
	}
	return nil
}

// Проверка максимального времени окончания встречи (не позже 19:00)
func (m *Meeting) CheckEndDateMaxTime() error {
	if err := m.CheckEndDateNotEmpty(); err != nil {
		return err
	}

	if m.EndDate.Hour() > 19 {
		return errors.New("время окончания встречи не может быть позже 19:00")
	}
	return nil
}
