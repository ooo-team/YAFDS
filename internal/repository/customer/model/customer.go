package model

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID        uint32
	Info      CustomerInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CustomerInfo struct {
	Phone   string
	Email   string
	Address string
}
