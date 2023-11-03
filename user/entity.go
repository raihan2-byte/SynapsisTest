package user

import (
	"time"
)

type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      int
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
	// Transaction transaction.TransactionHistory
}

