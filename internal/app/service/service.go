package service

import (
	"ais/internal/app/repository"
	"ais/internal/domain"
	"errors"
	"fmt"
)

type MeetingService struct {
	repo *repository.MeetingRepository
}

func NewMeetingService(repo *repository.MeetingRepository) *MeetingService {
	return &MeetingService{repo: repo}
}

func (s *MeetingService) CreateMeeting(meeting *domain.Meeting) error {
	fmt.Printf("StartDate: %v\n", meeting.StartDate)
	if meeting == nil {
		return errors.New("meeting can not be nil")
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
