package service

import (
	"ais/internal/app/repository"
	"ais/internal/domain"
	"errors"
)

type MeetingService struct {
	repo *repository.MeetingRepository
}

func NewMeetingService(repo *repository.MeetingRepository) *MeetingService {
	return &MeetingService{repo: repo}
}

func (s *MeetingService) CreateMeeting(meeting *domain.Meeting) error {
	if meeting == nil {
		return errors.New("встреча пустая")
	}

	if err := meeting.CheckMeetingNameNotWhitespace(); err != nil {
		return err
	}

	if err := meeting.CheckMeetingPlaceNotWhitespace(); err != nil {
		return err
	}

	if err := meeting.CheckCommentNotWhitespace(); err != nil {
		return err
	}

	if err := meeting.CheckApplicantIDNotWhitespace(); err != nil {
		return err
	}

	if err := meeting.CheckStartDateNotEmpty(); err != nil {
		return err
	}

	if err := meeting.CheckEndDateNotEmpty(); err != nil {
		return err
	}

	if err := meeting.CheckStartDateMinTime(); err != nil {
		return err
	}

	if err := meeting.CheckEndDateMaxTime(); err != nil {
		return err
	}

	return s.repo.Create(meeting)
}

func (s *MeetingService) UpdateMeeting(meeting *domain.Meeting) error {
	if meeting.ID == "" {
		return errors.New("ID не заполнен")
	}
	return s.repo.Update(meeting)
}

func (s *MeetingService) DeleteMeeting(meetingID string) error {
	if meetingID == "" {
		return errors.New("ID не заполнен")
	}
	return s.repo.Delete(meetingID)
}
