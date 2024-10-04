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
		return errors.New("meeting can not be nil")
	}

}
