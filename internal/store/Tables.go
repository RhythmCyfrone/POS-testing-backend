package store

import (
	"context"
	"database/sql"
)

type TablesStore struct {
	db *sql.DB
}

type Table struct {
	TableId                       int    `json:"tableId"`
	TableName                     string `json:"tableName"`
	TableMapperId                 int    `json:"tableMapperId"`
	TableMaxPax                   int    `json:"tableMaxPax"`
	DateTimeSinceLastStatusChange string `json:"dateTimeSinceLastStatusChange"`
	TableTrackingStatusId         int    `json:"tableTrackingStatusId"`
	TableTrackingStatusName       string `json:"tableTrackingStatusName"`
}

func (s *TablesStore) Create(ctx context.Context, table *Table) error {
	query := `
		INSERT INTO tables (id)
		VALUES ($1)
	`
	err := s.db.QueryRowContext(ctx, query, table.TableId).Scan(&table.TableId)

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
		err := rows.Scan(&table.TableId, &table.TableName, &table.TableMapperId, &table.TableMaxPax,
			&table.DateTimeSinceLastStatusChange, &table.TableTrackingStatusId, &table.TableTrackingStatusName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}
