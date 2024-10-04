package dto

import "time"

type CreateMeetingRequest struct {
	Name           string    `json:"name"`
	Place          string    `json:"place"`
	Comment        string    `json:"comment"`
	ApplicantID    string    `json:"applicant_id"`
	StartDate      time.Time `json:"start_date"`
	StartTime      time.Time `json:"start_time"`
	EndDate        time.Time `json:"end_date"`
	EndTime        time.Time `json:"end_time"`
	IsFullDay      bool      `json:"is_full_day"`
	IsOnline       bool      `json:"is_online"`
	IsOutlookEvent bool      `json:"is_outlook_event"`
	EmailAuthor    string    `json:"email_author"`
	IsViewMeeting  bool      `json:"is_view_meeting"`
	IsStartMeeting bool      `json:"is_start_meeting"`
}

type UpdateMeetingRequest struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Place     string    `json:"place"`
	Comment   string    `json:"comment"`
	StartDate time.Time `json:"start_date"`
	StartTime time.Time `json:"start_time"`
	EndDate   time.Time `json:"end_date"`
	EndTime   time.Time `json:"end_time"`
	IsFullDay bool      `json:"is_full_day"`
	IsOnline  bool      `json:"is_online"`
}

type MeetingResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Place       string    `json:"place"`
	StartDate   time.Time `json:"start_date"`
	StartTime   time.Time `json:"start_time"`
	EndDate     time.Time `json:"end_date"`
	EndTime     time.Time `json:"end_time"`
	IsFullDay   bool      `json:"is_full_day"`
	IsOnline    bool      `json:"is_online"`
	EmailAuthor string    `json:"email_author"`
}
