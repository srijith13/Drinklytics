package models

import (
	"time"
)

// Authentication

// Users
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetails struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Finance
type TransactionTypes struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type TransactionRecords struct {
	ID                int64     `json:"id"`
	Amount            float64   `json:"amount"`
	TransactionTypeId int64     `json:"trx_type_id"`
	TransactionType   string    `json:"trx_type"`
	Category          string    `json:"category"`
	Note              string    `json:"note"`
	TransactionDate   string    `json:"trx_date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsActive          bool      `json:"is_active"`
}

type TransactionRequests struct {
	ID                int64   `json:"id"`
	Amount            float64 `json:"amount"`
	TransactionTypeId int64   `json:"trx_type_id"`
	TransactionType   string  `json:"trx_type"`
	Category          string  `json:"category"`
	Note              string  `json:"note"`
	TransactionDate   string  `json:"trx_date"`
	Page              int     `json:"page"`
	Limit             int     `json:"limit"`
}

type FinancialSummaryRequests struct {
	TransactionDate string `json:"trx_date"`
	Page            int    `json:"page"`
	Limit           int    `json:"limit"`
}

type FinancialSummaryResponse struct {
	Category     string  `json:"category"`
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	NetBalance   float64 `json:"net_balance"`
}
