package honey

import (
	"honey/internal/config"

	"github.com/jmoiron/sqlx"
)

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
	db   *sqlx.DB
	conf *config.Config
}

//New ...
func New(db *sqlx.DB, c *config.Config) (HoneyService, error) {
	return service{db, c}, nil
}

func (s service) AddSeason(sn Season) error {
	return nil
}

func (s service) FindByID(ID int) *Season {
	return nil
}

func (s service) FindAll() []*Season {
	var list []*Season
	s.db.Select(&list, "SELECT * FROM seasons")
	return list
}
