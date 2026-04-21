package services

import (
	"database/sql"
	"drinklytics/internal/helper"
	"drinklytics/internal/models"
	"fmt"
	"log"
	"strings"
	"time"
)

type finService struct {
	db *sql.DB
}

type FinService interface {
	CreateTransactionTypes(request *models.TransactionTypes) (models.TransactionTypes, error)
	GetAllTransactionTypes() ([]models.TransactionTypes, error)
	UpdateTransactionTypes(request *models.TransactionTypes) (models.TransactionTypes, error)
	DeleteTransactionTypes(transactionTypeId *int64) error

	CreateTransaction(request *models.TransactionRecords) (models.TransactionRecords, error)
	GetTransaction(request *models.TransactionRequests, filters map[string]string) ([]models.TransactionRecords, error)
	UpdateTransaction(request *models.TransactionRecords) (models.TransactionRecords, error)
	DeleteTransaction(transactionTypeId *int64) error

	FinancialSummary(request *models.FinancialSummaryRequests, filters map[string]string) (models.FinancialSummaryResponse, error)
	CategorySummary(request *models.FinancialSummaryRequests, filters map[string]string) ([]models.FinancialSummaryResponse, error)
	RecentActivity() ([]models.FinancialSummaryResponse, error)
	Trends(request *models.FinancialSummaryRequests, trends string) (models.FinancialSummaryResponse, error)
}

func NewFinService(db *sql.DB) FinService {
	return &finService{db}
}

func (s *finService) CreateTransactionTypes(request *models.TransactionTypes) (models.TransactionTypes, error) {
	var transactionType models.TransactionTypes

	query := `INSERT INTO transaction_type (name,is_active) VALUES ($1, $2) RETURNING *;`
	err := s.db.QueryRow(query, request.Name, true).Scan(&transactionType.ID, &transactionType.Name, &transactionType.IsActive)
	if err != nil {
		log.Println("Error Executing query: ", err)
		return models.TransactionTypes{}, err
	}
	return transactionType, nil
}

func (s *finService) GetAllTransactionTypes() ([]models.TransactionTypes, error) {
	var transactionTypes []models.TransactionTypes
	query := `SELECT * FROM transaction_type WHERE is_active = true`

	rows, err := s.db.Query(query)
	defer rows.Close()

	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.TransactionTypes{}, fmt.Errorf("Error Executing query: %w", err)
	}
	err = helper.TransactionTypesDtoMapper(rows, &transactionTypes)
	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.TransactionTypes{}, err
	}
	return transactionTypes, nil
}

func (s *finService) UpdateTransactionTypes(request *models.TransactionTypes) (models.TransactionTypes, error) {
	var transactionType models.TransactionTypes

	query := `UPDATE transaction_type SET name = $1, is_active = $2 WHERE id = $3 AND (name, is_active) IS DISTINCT FROM ($1,$2) RETURNING *`
	err := s.db.QueryRow(query, request.Name, request.IsActive, request.ID).Scan(&transactionType.ID, &transactionType.Name, &transactionType.IsActive)

	if err != nil {
		log.Println("Error Executing query:", err)
		return models.TransactionTypes{}, err
	}

	return transactionType, nil
}

func (s *finService) DeleteTransactionTypes(transactionTypeId *int64) error {
	query := `UPDATE transaction_type SET is_active = false WHERE id = $1`
	_, err := s.db.Exec(query, transactionTypeId)
	if err != nil {
		log.Println("Error Executing query:", err)
	}
	return err
}

func (s *finService) CreateTransaction(request *models.TransactionRecords) (models.TransactionRecords, error) {
	var trxRecord models.TransactionRecords

	fmt.Println("request.TransactionDate ", request.TransactionDate)
	query := `INSERT INTO transactions (amount,transaction_type_id,category,note,transaction_date,created_at,updated_at,is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;`
	err := s.db.QueryRow(query, request.Amount, request.TransactionTypeId, request.Category, request.Note, request.TransactionDate, time.Now(), time.Now(), true).
		Scan(&trxRecord.ID, &trxRecord.Amount, &trxRecord.TransactionTypeId, &trxRecord.Category, &trxRecord.Note, &trxRecord.TransactionDate, &trxRecord.CreatedAt, &trxRecord.UpdatedAt, &trxRecord.IsActive)

	if err != nil {
		log.Println("Error Executing query: ", err)
		return models.TransactionRecords{}, err
	}
	return trxRecord, nil
}

func (s *finService) GetTransaction(request *models.TransactionRequests, filters map[string]string) ([]models.TransactionRecords, error) {
	if request.Limit == 0 {
		request.Limit = 5
	}
	if request.Page == 0 {
		request.Page = 1
	}
	limit := request.Limit
	offset := 0
	if request.Page > 0 {
		offset = (limit * (request.Page - 1))
	}
	var transactionRecords []models.TransactionRecords
	query := `SELECT * FROM transactions `

	conditions := []string{}
	args := []interface{}{}
	i := 1

	conditions = append(conditions, "is_active= true ")

	if (filters["start_date"]) != "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date >= $%d", i))
		args = append(args, filters["start_date"])
		i++
	}

	if filters["end_date"] != "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date <= $%d", i))
		args = append(args, filters["end_date"])
		i++
	}

	if filters["trx_type_id"] != "" {
		conditions = append(conditions, fmt.Sprintf("transaction_type_id = $%d", i))
		args = append(args, filters["trx_type_id"])
		i++
	}

	if filters["category"] != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", i))
		args = append(args, filters["category"])
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY transaction_date DESC, id DESC LIMIT $%d OFFSET $%d", i, i+1)

	args = append(args, limit, offset)
	rows, err := s.db.Query(query, args...)

	defer rows.Close()

	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.TransactionRecords{}, fmt.Errorf("Error Executing query: %w", err)
	}
	err = helper.TransactionRecordsDtoMapper(rows, &transactionRecords)
	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.TransactionRecords{}, err
	}
	return transactionRecords, nil
}

func (s *finService) UpdateTransaction(request *models.TransactionRecords) (models.TransactionRecords, error) {
	var trxRecord models.TransactionRecords

	query := `UPDATE transactions SET amount = $1,transaction_type_id  =$2,category = $3,note = $4,transaction_date = $5, updated_at = $6  WHERE id = $7 AND is_active = true
			  AND (amount,transaction_type_id,category,note,transaction_date) IS DISTINCT FROM ($1, $2, $3, $4, $5) RETURNING *`
	err := s.db.QueryRow(query, request.Amount, request.TransactionTypeId, request.Category, request.Note, request.TransactionDate, time.Now(), request.ID).
		Scan(&trxRecord.ID, &trxRecord.Amount, &trxRecord.TransactionTypeId, &trxRecord.Category, &trxRecord.Note, &trxRecord.TransactionDate, &trxRecord.CreatedAt, &trxRecord.UpdatedAt, &trxRecord.IsActive)

	if err != nil {
		log.Println("Error Executing query:", err)
		return models.TransactionRecords{}, err
	}

	return trxRecord, nil
}

func (s *finService) DeleteTransaction(trxRecordId *int64) error {
	query := `UPDATE transactions SET is_active = false WHERE id = $1 and is_active = true`
	_, err := s.db.Exec(query, trxRecordId)
	if err != nil {
		log.Println("Error Executing query:", err)
	}
	return err
}

// Dashboard
func (s *finService) FinancialSummary(request *models.FinancialSummaryRequests, filters map[string]string) (models.FinancialSummaryResponse, error) {
	var financialSummary models.FinancialSummaryResponse
	query := `SELECT total_income,total_expence,(total_income-total_expence) AS net_balance from
				(SELECT SUM(CASE WHEN   transaction_type_id = 1 THEN amount ELSE 0 END) AS total_income,
				SUM(CASE WHEN   transaction_type_id = 2 THEN amount ELSE 0 END) AS total_expence
				FROM transactions `

	conditions := []string{}
	args := []interface{}{}
	i := 1

	conditions = append(conditions, "is_active= true ")

	if request.TransactionDate != "" && filters["start_date"] == "" && filters["end_date"] == "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date = $%d", i))
		args = append(args, request.TransactionDate)
		i++
	} else if filters["start_date"] != "" && filters["end_date"] == "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date >= $%d", i))
		args = append(args, filters["start_date"])
		i++
	} else if filters["start_date"] != "" && filters["end_date"] != "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date between $%d and $%d", i, i+1))
		args = append(args, filters["start_date"], filters["end_date"])
		i = i + 2
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ") + " ) "
	}

	err := s.db.QueryRow(query, args...).
		Scan(&financialSummary.TotalIncome, &financialSummary.TotalExpense, &financialSummary.NetBalance)

	if err != nil {
		log.Println("Error Executing query:", err)
		return models.FinancialSummaryResponse{}, err
	}
	return financialSummary, nil
}

func (s *finService) CategorySummary(request *models.FinancialSummaryRequests, filters map[string]string) ([]models.FinancialSummaryResponse, error) {
	var financialSummary []models.FinancialSummaryResponse
	query := `SELECT category,
    SUM(CASE WHEN transaction_type_id = 1 THEN amount ELSE 0 END) AS total_income,
    SUM(CASE WHEN transaction_type_id = 2 THEN amount ELSE 0 END) AS total_expenses
    FROM transactions`

	conditions := []string{}
	args := []interface{}{}
	i := 1

	conditions = append(conditions, "is_active= true ")

	if request.TransactionDate != "" && filters["start_date"] == "" && filters["end_date"] == "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date = $%d", i))
		args = append(args, request.TransactionDate)
		i++
	} else if filters["start_date"] != "" && filters["end_date"] == "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date >= $%d", i))
		args = append(args, filters["start_date"])
		i++
	} else if filters["start_date"] != "" && filters["end_date"] != "" {
		conditions = append(conditions, fmt.Sprintf("transaction_date between $%d and $%d", i, i+1))
		args = append(args, filters["start_date"], filters["end_date"])
		i = i + 2
	}

	if filters["category"] != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", i))
		args = append(args, filters["category"])
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ") + " GROUP BY category ORDER BY category "
	}

	rows, err := s.db.Query(query, args...)

	defer rows.Close()

	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.FinancialSummaryResponse{}, fmt.Errorf("Error Executing query: %w", err)
	}
	err = helper.FinancialSummaryDtoMapper(rows, &financialSummary)
	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.FinancialSummaryResponse{}, err
	}
	return financialSummary, nil
}

func (s *finService) RecentActivity() ([]models.FinancialSummaryResponse, error) {
	var financialSummary []models.FinancialSummaryResponse
	query := `  SELECT category, total_income, total_expence  from
				(SELECT category, DATE_TRUNC('week', transaction_date) AS period,
				SUM(CASE WHEN   transaction_type_id = 1 THEN amount ELSE 0 END) AS total_income,
				SUM(CASE WHEN   transaction_type_id = 2 THEN amount ELSE 0 END) AS total_expence
				FROM transactions WHERE is_active = true GROUP BY category, period ORDER BY category,period)`

	rows, err := s.db.Query(query)

	defer rows.Close()

	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.FinancialSummaryResponse{}, fmt.Errorf("Error Executing query: %w", err)
	}
	err = helper.FinancialSummaryDtoMapper(rows, &financialSummary)
	if err != nil {
		log.Println("Error Executing query:", err)
		return []models.FinancialSummaryResponse{}, err
	}
	return financialSummary, nil
}

func (s *finService) Trends(request *models.FinancialSummaryRequests, trends string) (models.FinancialSummaryResponse, error) {
	var financialSummary models.FinancialSummaryResponse
	query := ` SELECT total_income, total_expence, (total_income-total_expence) AS net_balance from
				(SELECT DATE_TRUNC($1, transaction_date) AS period,SUM(CASE WHEN   transaction_type_id = 1 THEN amount ELSE 0 END) AS total_income,
				SUM(CASE WHEN transaction_type_id = 2 THEN amount ELSE 0 END) AS total_expence
				FROM transactions WHERE is_active= true GROUP BY period ORDER BY period)`

	err := s.db.QueryRow(query, trends).
		Scan(&financialSummary.TotalIncome, &financialSummary.TotalExpense, &financialSummary.NetBalance)

	if err != nil {
		log.Println("Error Executing query:", err)
		return models.FinancialSummaryResponse{}, fmt.Errorf("Error Executing query: %w", err)
	}

	if err != nil {
		log.Println("Error Executing query:", err)
		return models.FinancialSummaryResponse{}, err
	}
	return financialSummary, nil
}
