package repository

import (
	"ais/internal/domain"
	"database/sql"
	"errors"
)

type MeetingRepository struct {
	db *sql.DB
}

// NewMeetingRepository создает новый MeetingRepository.
func NewMeetingRepository(db *sql.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

// Create создает новое событие.
func (r *MeetingRepository) Create(meeting *domain.Meeting) error {
	if meeting == nil {
		return errors.New("meeting cannot be nil")
	}

	query := `
		INSERT INTO tb_meeting (name, place, comment, applicant_id, start_date, start_time, end_date, end_time, is_full_day, is_online, is_outlook_event, email_author, is_view_meeting, is_start_meeting) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := r.db.Exec(query, meeting.Name, meeting.Place, meeting.Comment, meeting.ApplicantID,
		meeting.StartDate, meeting.StartTime, meeting.EndDate, meeting.EndTime,
		meeting.IsFullDay, meeting.IsOnline, meeting.IsOutlookEvent,
		meeting.EmailAuthor, meeting.IsViewMeeting, meeting.IsStartMeeting)

	return err
}
