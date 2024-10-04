package domain

import "time"

type Meeting struct {
	ID             string    // Идентификатор встречи
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
