package honey

import (
	"honey/internal/config"

	"github.com/jmoiron/sqlx"
)

//Season ...
type Season struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//HoneyService ...
type HoneyService interface {
	AddSeason(Season)
	UpdateSeason(Season, int) error
	FindByID(int) Season
	FindAll() []*Season
	DeleteSeason(int) error
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New ...
func New(db *sqlx.DB, c *config.Config) (HoneyService, error) {
	return service{db, c}, nil
}

func (s service) DeleteSeason(ID int) error {
	_, err := s.db.Exec("DELETE FROM seasons WHERE id=?", ID)
	return err
}

func (s service) UpdateSeason(sn Season, ID int) error {
	_, err := s.db.Exec("UPDATE seasons SET name = ? WHERE id = ?", sn.Name, ID)
	return err
}

func (s service) AddSeason(sn Season) {
	insertSeason := `INSERT INTO seasons (name) VALUES (?)`
	s.db.MustExec(insertSeason, sn.Name)
}

func (s service) FindByID(ID int) Season {
	var sn Season
	s.db.Get(&sn, "SELECT * FROM seasons WHERE id = ?;", ID)
	return sn
}

func (s service) FindAll() []*Season {
	var list []*Season
	s.db.Select(&list, "SELECT * FROM seasons")
	return list
}
