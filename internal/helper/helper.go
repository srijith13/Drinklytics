package helper

import (
	"database/sql"
	"drinklytics/internal/models"
	"log"
)

// Mapping the db queried details
func UsersDetailsDtoMapper(rows *sql.Rows, users *[]models.UserDetails) error {
	log.Println("Scanning SQl rows to map")
	for rows.Next() {
		usersDto := models.UserDetails{}
		err := rows.Scan(&usersDto.ID,
			&usersDto.Name,
			&usersDto.Email,
			&usersDto.Role,
			&usersDto.IsActive,
			&usersDto.CreatedAt,
			&usersDto.UpdatedAt)
		*users = append(*users, usersDto)
		if err == sql.ErrNoRows {
			log.Println("No data found")
			return err
		} else if err != nil {
			log.Println("Error scanning row:", err)
			return err
		}
	}
	return nil
}

func UserDtoMapper(row *sql.Row, user *models.User) error {
	log.Println("Scanning SQl rows to map")
	err := row.Scan(&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		log.Println("Error scanning row:", err)
		return err

	}
	return nil
}

// Finace mapping
func TransactionTypesDtoMapper(rows *sql.Rows, trxTypes *[]models.TransactionTypes) error {
	log.Println("Scanning SQl rows to map")
	for rows.Next() {
		trxTypeDto := models.TransactionTypes{}
		err := rows.Scan(&trxTypeDto.ID,
			&trxTypeDto.Name,
			&trxTypeDto.IsActive)
		*trxTypes = append(*trxTypes, trxTypeDto)
		if err == sql.ErrNoRows {
			log.Println("No data found")
			return err
		} else if err != nil {
			log.Println("Error scanning row:", err)
			return err
		}
	}
	return nil
}

// TransactionRecordsDtoMapper
func TransactionRecordsDtoMapper(rows *sql.Rows, trxRecords *[]models.TransactionRecords) error {
	log.Println("Scanning SQl rows to map")
	for rows.Next() {
		trxRecordDto := models.TransactionRecords{}
		err := rows.Scan(&trxRecordDto.ID,
			&trxRecordDto.Amount,
			&trxRecordDto.TransactionTypeId,
			&trxRecordDto.Category,
			&trxRecordDto.Note,
			&trxRecordDto.TransactionDate,
			&trxRecordDto.CreatedAt,
			&trxRecordDto.UpdatedAt,
			&trxRecordDto.IsActive)
		*trxRecords = append(*trxRecords, trxRecordDto)
		if err == sql.ErrNoRows {
			log.Println("No data found")
			return err
		} else if err != nil {
			log.Println("Error scanning row:", err)
			return err
		}
	}
	return nil
}

func FinancialSummaryDtoMapper(rows *sql.Rows, finSummaries *[]models.FinancialSummaryResponse) error {
	log.Println("Scanning SQl rows to map")
	for rows.Next() {
		finSummaryDto := models.FinancialSummaryResponse{}
		err := rows.Scan(&finSummaryDto.Category,
			&finSummaryDto.TotalIncome,
			&finSummaryDto.TotalExpense)
		*finSummaries = append(*finSummaries, finSummaryDto)
		if err == sql.ErrNoRows {
			log.Println("No data found")
			return err
		} else if err != nil {
			log.Println("Error scanning row:", err)
			return err
		}
	}
	return nil
}
