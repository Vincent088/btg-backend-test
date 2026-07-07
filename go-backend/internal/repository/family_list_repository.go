package repository

import (
	"database/sql"

	"github.com/Vincent088/btg-backend-test/go-backend/internal/entity"
)

type FamilyListRepository struct {
	db *sql.DB
}

func NewFamilyListRepository(db *sql.DB) *FamilyListRepository {
	return &FamilyListRepository{db: db}
}

func (r *FamilyListRepository) Create(tx *sql.Tx, f entity.FamilyList) error {
	_, err := tx.Exec(
		`INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob) VALUES ($1, $2, $3, $4)`,
		f.CstID, f.FlRelation, f.FlName, f.FlDob,
	)
	return err
}

func (r *FamilyListRepository) GetByCustomerID(cstID int) ([]entity.FamilyList, error) {
	rows, err := r.db.Query(
		`SELECT fl_id, cst_id, fl_relation, fl_name, fl_dob FROM family_list WHERE cst_id = $1`,
		cstID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []entity.FamilyList{}
	for rows.Next() {
		var f entity.FamilyList
		if err := rows.Scan(&f.FlID, &f.CstID, &f.FlRelation, &f.FlName, &f.FlDob); err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}

func (r *FamilyListRepository) DeleteByCustomerID(tx *sql.Tx, cstID int) error {
	_, err := tx.Exec(`DELETE FROM family_list WHERE cst_id = $1`, cstID)
	return err
}
