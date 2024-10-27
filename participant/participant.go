package participant

type ID int

type Participant struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	IsPresent bool   `json:"isPresent"`
}

type Reader interface {
	Get(id ID) (*Participant, error)
	Search(query string) ([]*Participant, error)
	List() ([]*Participant, error)
	Update(id string) (string, error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	Update(id string) ([]*Participant, error)
	Search(query string) ([]*Participant, error)
	List() ([]*Participant, error)
}
