package repository

import (
	"ais/internal/domain"
	"database/sql"
)

type MeetingRepository struct {
	db *sql.DB
}

// NewMeetingRepository создает новый MeetingRepository.
func NewMeetingRepository(db *sql.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

func (r *MeetingRepository) Create(meeting *domain.Meeting) error {
	query := `
		INSERT INTO tb_meeting (name, place, comment, applicant_id, start_date, end_date, is_full_day, is_online, is_outlook_event, email_author, is_view_meeting, is_start_meeting)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	_, err := r.db.Exec(query,
		meeting.Name,
		meeting.Place,
		meeting.Comment,
		meeting.ApplicantID,
		meeting.StartDate,
		meeting.EndDate,
		meeting.IsFullDay,
		meeting.IsOnline,
		meeting.IsOutlookEvent,
		meeting.EmailAuthor,
		meeting.IsViewMeeting,
		meeting.IsStartMeeting,
	)

	return err
}
