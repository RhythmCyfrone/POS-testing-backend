package store

import (
	"context"
	"database/sql"
)

type OrdersStore struct {
	db *sql.DB
}

type Order struct {
	OrderId        string `json:"orderId"`
	TableId        string `json:"tableId"`
	CustomerId     string `json:"customerId"`
	CustomerName   string `json:"customerName"`
	OrderTakerId   string `json:"orderTakerId"`
	OrderTakerName string `json:"orderTakerName"`
	OrderStatusId  int    `json:"orderStatusId"`
	BillID         int    `json:"billId"`
}

func (s *OrdersStore) GetOrderByTableId(ctx context.Context, tableId string) (Order, error) {
	query := `
		SELECT *
		FROM orders
		WHERE TableId = ?
	`
	var order Order
	err := s.db.QueryRowContext(ctx, query, tableId).Scan(&order.OrderId, &order.TableId, &order.CustomerId, &order.CustomerName, &order.OrderTakerId, &order.OrderTakerName, &order.OrderStatusId, &order.BillID)
	if err != nil {
		return Order{}, err
	}

	return order, err
}
