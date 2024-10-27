package programapocalipse

import (
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	client HTTPClient
	url    string
}

func NewService() *Service {
	return &Service{
		client: &http.Client{Timeout: time.Duration(5) * time.Second},
		url:    "https://encontrocomavida.tv/inscrito.php?id_inscrito=",
	}
}

func (s *Service) Get(participant_id string) (string, error) {
	url := fmt.Sprintf("%s%s", s.url, participant_id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "erro", err
	}

	resp, err := s.client.Do(request)
	if err != nil {
		return "erro", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "erro", fmt.Errorf("requisição falhou com status %d", resp.StatusCode)
	}

	fmt.Println(resp)

	return resp.Status, nil
}
