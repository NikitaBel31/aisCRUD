package repository

import (
	"ais/internal/domain"
	"database/sql"
	"errors"
	"fmt"
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

func (r *MeetingRepository) Update(meeting *domain.Meeting) error {
	query := `
		UPDATE tb_meeting 
		SET name = $1, place = $2, comment = $3, applicant_id = $4, start_date = $5, 
			end_date = $6, is_full_day = $7, is_online = $8, is_outlook_event = $9, 
			email_author = $10, is_view_meeting = $11, is_start_meeting = $12 
		WHERE id = $13
	`

	result, err := r.db.Exec(query, meeting.Name, meeting.Place, meeting.Comment,
		meeting.ApplicantID, meeting.StartDate, meeting.EndDate, meeting.IsFullDay,
		meeting.IsOnline, meeting.IsOutlookEvent, meeting.EmailAuthor,
		meeting.IsViewMeeting, meeting.IsStartMeeting, meeting.ID)
	if err != nil {
		return fmt.Errorf("ошибка изменения встречи: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка проверки результата изменения: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("встреча не найдена")
	}

	return nil
}
