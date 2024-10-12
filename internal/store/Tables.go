package store

import (
	"context"
	"database/sql"
)

type TablesStore struct {
	db *sql.DB
}

type Table struct {
	TableID     string `json:"id"`
	Floor       string `json:"floor"`
	Siting      int64  `json:"siting"`
	Curr_status string `json:"curr_status"`
	BranchID    string `json:"branchID"`
}

func (s *TablesStore) Create(ctx context.Context, table *Table) error {
	query := `
		INSERT INTO tables (id)
		VALUES ($1)
	`
	err := s.db.QueryRowContext(ctx, query, table.TableID).Scan(&table.TableID)

	return err
}

func (s *TablesStore) GetAllTables(ctx context.Context) ([]Table, error) {
	query := `
		SELECT *
		FROM tables
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		err := rows.Scan(&table.TableID, &table.Floor, &table.Siting, &table.Curr_status, &table.BranchID)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}
