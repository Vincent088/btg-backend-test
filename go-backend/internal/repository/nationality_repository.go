package repository

import (
	"database/sql"

	"github.com/Vincent088/btg-backend-test/go-backend/internal/entity"
)

type NationalityRepository struct {
	db *sql.DB
}

func NewNationalityRepository(db *sql.DB) *NationalityRepository {
	return &NationalityRepository{db: db}
}

func (r *NationalityRepository) GetAll() ([]entity.Nationality, error) {
	rows, err := r.db.Query(`SELECT nationality_id, nationality_name, nationality_code FROM nationality`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entity.Nationality
	for rows.Next() {
		var n entity.Nationality
		if err := rows.Scan(&n.NationalityID, &n.NationalityName, &n.NationalityCode); err != nil {
			return nil, err
		}
		list = append(list, n)
	}
	return list, nil
}

func (r *NationalityRepository) GetByID(id int) (*entity.Nationality, error) {
	var n entity.Nationality
	err := r.db.QueryRow(
		`SELECT nationality_id, nationality_name, nationality_code FROM nationality WHERE nationality_id = $1`,
		id,
	).Scan(&n.NationalityID, &n.NationalityName, &n.NationalityCode)
	if err != nil {
		return nil, err
	}
	return &n, nil
}
