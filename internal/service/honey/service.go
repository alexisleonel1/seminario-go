package honey

import "honey/internal/config"

//Season ...
type Season struct {
	ID   int64
	Name string
}

//HoneyService ...
type HoneyService interface {
	AddSeason(Season) error
	FindByID(int) *Season
	FindAll() []*Season
}

type service struct {
	conf *config.Config
}

//New ...
func New(c *config.Config) (HoneyService, error) {
	return service{c}, nil
}

func (s service) AddSeason(sn Season) error {
	return nil
}

func (s service) FindByID(ID int) *Season {
	return nil
}

func (s service) FindAll() []*Season {
	var list []*Season
	list = append(list, &Season{0, "temporada 1"})
	return list
}
