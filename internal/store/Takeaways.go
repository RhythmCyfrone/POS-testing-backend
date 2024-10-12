package store

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type TakeawayStore struct {
	db *sql.DB
}

type Takeaway struct {
	OrderId       string `json:"orderId"`
	CustomerId    string `json:"customerId"`
	CustomerName  string `json:"customerName"`
	CustomerPhone string `json:"customerPhone"`
	OrderStatusId int    `json:"orderStatusId"`
	CurrentStatus string `json:"curr_status"`
	BillID        int    `json:"billId"`
	BranchID      string `json:"branchID"`
}

func FormatTakeawayCount(id int64) string {
	return fmt.Sprintf("TW%03d", id)
}

func ReverseFormatTakeawayCount(orderId string) (int64, error) {
	// Remove the "TW" prefix
	numericPart := strings.TrimPrefix(orderId, "TW")

	// Convert the remaining string to an int64
	id, err := strconv.ParseInt(numericPart, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid orderId format: %w", err)
	}

	return id, nil
}

func (s *TakeawayStore) GetTakeawayOrders(ctx context.Context) ([]Takeaway, error) {
	query := `
		SELECT *
		FROM takeaways
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var takeaways []Takeaway
	for rows.Next() {
		var takeaway Takeaway
		var id int64
		err := rows.Scan(&id, &takeaway.CustomerId, &takeaway.CustomerName,
			&takeaway.CustomerPhone, &takeaway.OrderStatusId, &takeaway.CurrentStatus, &takeaway.BillID, &takeaway.BranchID)

		if err != nil {
			return nil, err
		}
		takeaway.OrderId = FormatTakeawayCount(id)
		takeaways = append(takeaways, takeaway)
	}

	return takeaways, nil
}

type NewOrderDetails struct {
	OrderId string `json:"orderId"`
}

func (s *TakeawayStore) CreateTakeawayOrder(ctx context.Context, customerName string, customerPhone *string) (NewOrderDetails, error) {

	var phone sql.NullString
	if customerPhone != nil {
		phone = sql.NullString{String: *customerPhone, Valid: true}
	} else {
		phone = sql.NullString{Valid: false}
	}
	query := `
		INSERT INTO takeaways (CustomerId, CustomerName, CustomerPhone, OrderStatusId, CurrentStatus, BillId, BranchId)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.ExecContext(ctx, query, '1', customerName, phone, 1, "Ordered", 1234, 1)

	id, _ := result.LastInsertId()

	if err != nil {
		return NewOrderDetails{}, err
	}
	return NewOrderDetails{OrderId: FormatTakeawayCount(id)}, nil
}

func (s *TakeawayStore) GetTakeawayOrdersById(ctx context.Context, orderId string) (Takeaway, error) {
	query := `
		SELECT *
		FROM takeaways
		WHERE OrderId = ?
	`
	var takeaway Takeaway

	id, _ := ReverseFormatTakeawayCount(orderId)

	err := s.db.QueryRowContext(ctx, query, id).Scan(&takeaway.OrderId, &takeaway.CustomerId, &takeaway.CustomerName,
		&takeaway.CustomerPhone, &takeaway.OrderStatusId, &takeaway.CurrentStatus, &takeaway.BillID, &takeaway.BranchID)
	if err != nil {
		return Takeaway{}, err
	}
	takeaway.OrderId = FormatTakeawayCount(id)

	return takeaway, nil

}
