package programapocalipse

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type UseCase interface {
	Get(participant_id string) (string, error)
}
