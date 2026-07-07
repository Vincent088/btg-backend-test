package repository

import (
	"database/sql"
	"strings"

	"github.com/Vincent088/btg-backend-test/go-backend/internal/entity"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(tx *sql.Tx, c entity.Customer) (int, error) {
	var id int
	err := tx.QueryRow(
		`INSERT INTO customer (nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email)
		 VALUES ($1, $2, $3, $4, $5) RETURNING cst_id`,
		c.NationalityID, c.CstName, c.CstDob, c.CstPhoneNum, c.CstEmail,
	).Scan(&id)
	return id, err
}

func (r *CustomerRepository) GetAll() ([]entity.Customer, error) {
	rows, err := r.db.Query(`SELECT cst_id, nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email FROM customer`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entity.Customer
	for rows.Next() {
		var c entity.Customer
		if err := rows.Scan(&c.CstID, &c.NationalityID, &c.CstName, &c.CstDob, &c.CstPhoneNum, &c.CstEmail); err != nil {
			return nil, err
		}
		c.CstName = strings.TrimSpace(c.CstName)
		list = append(list, c)
	}
	return list, nil
}

func (r *CustomerRepository) GetByID(id int) (*entity.Customer, error) {
	var c entity.Customer
	err := r.db.QueryRow(
		`SELECT cst_id, nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email FROM customer WHERE cst_id = $1`,
		id,
	).Scan(&c.CstID, &c.NationalityID, &c.CstName, &c.CstDob, &c.CstPhoneNum, &c.CstEmail)
	if err != nil {
		return nil, err
	}
	c.CstName = strings.TrimSpace(c.CstName)
	return &c, nil
}

func (r *CustomerRepository) Update(c entity.Customer) error {
	_, err := r.db.Exec(
		`UPDATE customer SET nationality_id=$1, cst_name=$2, cst_dob=$3, cst_phoneNum=$4, cst_email=$5 WHERE cst_id=$6`,
		c.NationalityID, c.CstName, c.CstDob, c.CstPhoneNum, c.CstEmail, c.CstID,
	)
	return err
}

func (r *CustomerRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM customer WHERE cst_id = $1`, id)
	return err
}

func (r *CustomerRepository) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}
