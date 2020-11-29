package honey

import (
	"honey/internal/config"

	"github.com/jmoiron/sqlx"
)

//Season ...
type Season struct {
	ID    int64
	Name  string
	Hives int64
	Drums int64
}

//HoneyService ...
type HoneyService interface {
	AddSeason(Season)
	UpdateSeason(Season, int) error
	FindByID(int) (Season, error)
	FindAll() ([]*Season, error)
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
	_, err := s.db.Exec("UPDATE seasons SET name = ?, hives = ?, drums = ? WHERE id = ?", sn.Name, sn.Hives, sn.Drums, ID)
	return err
}

func (s service) AddSeason(sn Season) {
	insertSeason := `INSERT INTO seasons (name, hives, drums) VALUES (?, ?, ?)`
	s.db.MustExec(insertSeason, sn.Name, sn.Hives, sn.Drums)
}

func (s service) FindByID(ID int) (Season, error) {
	var sn Season
	err := s.db.Get(&sn, "SELECT * FROM seasons WHERE id = ?;", ID)
	return sn, err
}

func (s service) FindAll() ([]*Season, error) {
	var list []*Season
	err := s.db.Select(&list, "SELECT * FROM seasons")
	return list, err
}
