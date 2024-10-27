package participant

import (
	"fmt"
	"strings"
)

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Clean() {
	s.r.Clean()
}

func (s *Service) Update(id string) (string, error) {
	_, err := s.r.Update(id)
	if err != nil {
		return "", fmt.Errorf("erro lendo person do repositório: %w", err)
	}

	return "Success", nil
}

func (s *Service) Search(query string) ([]*Participant, error) {
	p, err := s.r.Search(query)
	if err != nil {
		return nil, fmt.Errorf("erro buscando person do repositório: %w", err)
	}

	treatedParticipant := RemoveDuplicateParticipants(p)
	return treatedParticipant, nil
}

func (s *Service) List() ([]*Participant, error) {
	p, err := s.r.List()
	if err != nil {
		return nil, fmt.Errorf("erro listando person do repositório: %w", err)
	}

	treatedParticipant := RemoveDuplicateParticipants(p)
	return treatedParticipant, nil
}

func RemoveDuplicateParticipants(participants []*Participant) []*Participant {
	seen := make(map[string]bool)
	uniqueParticipants := make([]*Participant, 0, len(participants))

	for _, participant := range participants {
		name := strings.ToLower(participant.Name)
		if !seen[name] {
			seen[name] = true
			uniqueParticipants = append(uniqueParticipants, participant)
		}
	}

	return uniqueParticipants
}
