package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tables interface {
		Create(context.Context, *Table) error
		GetAllTables(context.Context) ([]Table, error)
	}
	Orders interface {
		GetOrderByTableId(context.Context, string) (Order, error)
	}
	Status interface {
		GetOrderStatusById(context.Context, int) (Status, error)
		UpdateOrderStatusByOrderId(context.Context, string, int) error
	}
	Takeaways interface {
		GetTakeawayOrders(context.Context) ([]Takeaway, error)
		CreateTakeawayOrder(context.Context, string, *string) (NewOrderDetails, error)
		GetTakeawayOrdersById(context.Context, string) (Takeaway, error)
	}
}

func NewSQLStorage(db *sql.DB) Storage {
	return Storage{
		Tables:    &TablesStore{db: db},
		Orders:    &OrdersStore{db: db},
		Status:    &StatusStore{db: db},
		Takeaways: &TakeawayStore{db: db},
	}
}
