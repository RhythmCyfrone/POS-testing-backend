package store

import (
	"context"
	"database/sql"
)

type StatusStore struct {
	db *sql.DB
}
type Status struct {
	StatusId    int    `json:"statusId"`
	StatusName  string `json:"statusName"`
	StatusColor string `json:"statusColor"`
}

func (s *StatusStore) GetOrderStatusById(ctx context.Context, statusId int) (Status, error) {
	query := `
		SELECT *
		FROM orderStatus
		WHERE StatusId = ?
	`
	var status Status
	err := s.db.QueryRowContext(ctx, query, statusId).Scan(&status.StatusId, &status.StatusName, &status.StatusColor)
	if err != nil {
		return Status{}, err
	}

	return status, err
}

func (s *StatusStore) UpdateOrderStatusByOrderId(ctx context.Context, orderId string, orderStatusId int) error {
	query := `
		UPDATE orders SET OrderStatusId = ?
		WHERE OrderId = ?
	`
	_, err := s.db.ExecContext(ctx, query, orderStatusId, orderId)
	if err != nil {
		return err
	}

	return nil
}
