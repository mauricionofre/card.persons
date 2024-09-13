package repository

import (
	"database/sql"

	"github.com/mauricionofre/person-api/internal/model"
)

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(p *model.Person) error {
	query := `INSERT INTO persons (name, age) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, p.Name, p.Age).Scan(&p.ID)
}

func (r *PersonRepository) GetByID(id int) (*model.Person, error) {
	query := `SELECT id, name, age FROM persons WHERE id = $1`
	p := &model.Person{}
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Age)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PersonRepository) Update(p *model.Person) error {
	query := `UPDATE persons SET name = $1, age = $2 WHERE id = $3`
	_, err := r.db.Exec(query, p.Name, p.Age, p.ID)
	return err
}

func (r *PersonRepository) Delete(id int) error {
	query := `DELETE FROM persons WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
