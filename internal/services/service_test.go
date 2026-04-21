package services

import (
	"drinklytics/internal/models"
	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestService_FinancialSummary(t *testing.T) {
	type args struct {
		request *models.FinancialSummaryRequests
		filters map[string]string
	}
	type want struct {
		response models.FinancialSummaryResponse
		err      error
	}

	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error Failed to create db tp mock: %v ", err)
	}
	defer db.Close()

	fc := &finService{db: db}

	query := `SELECT total_income,total_expence,(total_income-total_expence) AS net_balance from
				(SELECT SUM(CASE WHEN   transaction_type_id = 1 THEN amount ELSE 0 END) AS total_income,
				SUM(CASE WHEN   transaction_type_id = 2 THEN amount ELSE 0 END) AS total_expence
				FROM transactions `

	tests := []struct {
		name string
		args args
		want want
		mock func()
	}{
		{
			name: "Successful:  Total Income, Total Expenses, Net Balance without date filter",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "",
				},
				filters: map[string]string{},
			},

			mock: func() {

				// mockDb.ExpectQuery(query).WithArgs(sqlmock.AnyArg(), 440).WillReturnRows(
				// 	sqlmock.NewRows([]string{"id", "amount", "transaction_type_id", "category", "note", "transaction_date", "created_at", "updated_at", "is_active"}).
				// 		AddRow(1, 50000.00, 1, "Salary", "Monthly salary credited", "2026-04-01", "2026-04-04 17:43:17.821866", "2026-04-04 17:43:17.821866", true).
				// 		AddRow(2, 1500.00, 4, "Groceries", "Groceries from supermarket", "2026-04-02", "2026-04-04 17:44:42.467423", "2026-04-04 17:47:57.322166", true).
				// 		AddRow(3, 2000.00, 4, "Rent", "April house rent", "2026-04-02", "2026-04-04 18:01:09.201135", "2026-04-04 18:01:09.201135", true).
				// 		AddRow(4, 800.00, 4, "Fuel", "", "2026-04-02", "2026-04-04 18:01:41.663009", "2026-04-04 18:01:41.663009", true).
				// 		AddRow(5, 300.00, 4, "Entertainment", "Movie tickets", "2026-04-04", "2026-04-04 18:02:09.033384", "2026-04-04 18:03:24.435118", true).
				// 		AddRow(6, 450.00, 4, "Transport", "Taxi", "2026-04-04", "2026-04-04 18:04:23.246216", "2026-04-04 18:04:23.246216", true).
				// 		AddRow(7, 265.00, 4, "Transport", "Auto", "2026-04-05", "2026-04-04 18:11:56.25976", "2026-04-04 18:11:56.259761", true).
				// 		AddRow(8, 1250.00, 1, "Freelance", "Website project payment", "2026-04-04", "2026-04-04 18:14:51.693152", "2026-04-04 18:14:51.693152", true).
				// 		AddRow(9, 50.00, 4, "Freelance", "Website project payment", "2026-04-04", "2026-04-04 18:15:44.879426", "2026-04-04 18:15:44.879426", true).
				// 		AddRow(0, 560.00, 4, "Food", "Dinner at restaurant", "2026-04-05", "2026-04-04 18:17:46.481584", "2026-04-04 18:17:46.481584", true),
				// )
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"total_income", "total_expence", "net_balance"}).
						AddRow(51250, 5925, 45325),
				)
			},
			want: want{
				response: models.FinancialSummaryResponse{
					Category:     "",
					TotalIncome:  51250,
					TotalExpense: 5925,
					NetBalance:   45325,
				},
				err: nil,
			},
		},
		{
			name: "Successfl:  Total Income, Total Expenses, Net Balance with filter greater than or equal to startDate",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "",
				},
				filters: map[string]string{
					"start_date": "2026-04-03",
				},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"total_income", "total_expence", "net_balance"}).
						AddRow(1250, 1625, -375),
				)
			},
			want: want{
				response: models.FinancialSummaryResponse{
					Category:     "",
					TotalIncome:  1250,
					TotalExpense: 1625,
					NetBalance:   -375,
				},
				err: nil,
			},
		},
		{
			name: "Successful:  Total Income, Total Expenses, Net Balance with filter between startDate and endDate",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "",
				},
				filters: map[string]string{
					"start_date": "2026-04-03",
					"end_date":   "2026-04-05",
				},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"total_income", "total_expence", "net_balance"}).
						AddRow(51250, 5100, 46150),
				)
			},
			want: want{
				response: models.FinancialSummaryResponse{
					Category:     "",
					TotalIncome:  51250,
					TotalExpense: 5100,
					NetBalance:   46150,
				},
				err: nil,
			},
		},
		{
			name: "Successful:  Total Income, Total Expenses, Net Balance with on transactionDate",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "2026-04-04",
				},
				filters: map[string]string{},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"total_income", "total_expence", "net_balance"}).
						AddRow(1250, 800, 405),
				)
			},
			want: want{
				response: models.FinancialSummaryResponse{
					Category:     "",
					TotalIncome:  1250,
					TotalExpense: 800,
					NetBalance:   405,
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := fc.FinancialSummary(tt.args.request, tt.args.filters)

			if tt.want.err != nil {
				assert.EqualError(t, tt.want.err, err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want.response, result)
			}

		})
	}
}

func TestService_CategorySummary(t *testing.T) {
	type args struct {
		request *models.FinancialSummaryRequests
		filters map[string]string
	}
	type want struct {
		response []models.FinancialSummaryResponse
		err      error
	}

	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error Failed to create db tp mock: %v ", err)
	}
	defer db.Close()

	fc := &finService{db: db}

	query := `SELECT category,
    SUM(CASE WHEN transaction_type_id = 1 THEN amount ELSE 0 END) AS total_income,
    SUM(CASE WHEN transaction_type_id = 2 THEN amount ELSE 0 END) AS total_expenses
    FROM transactions`

	tests := []struct {
		name string
		args args
		want want
		mock func()
	}{
		{
			name: "Successful:  Total Income, Total Expenses, Net Balance without date filter",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "",
				},
				filters: map[string]string{},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"category", "total_income", "total_expence"}).
						AddRow("Entertainment", 0, 300).
						AddRow("Food", 0, 560).
						AddRow("Freelance", 1250, 50).
						AddRow("Fuel", 0, 800).
						AddRow("Groceries", 0, 1500).
						AddRow("Rent", 0, 9000).
						AddRow("Salary", 50000, 0),
				)
			},
			want: want{
				response: []models.FinancialSummaryResponse{
					{
						Category:     "Entertainment",
						TotalIncome:  0,
						TotalExpense: 300,
					},
					{
						Category:     "Food",
						TotalIncome:  0,
						TotalExpense: 560,
					},
					{
						Category:     "Freelance",
						TotalIncome:  1250,
						TotalExpense: 50,
					},
					{
						Category:     "Fuel",
						TotalIncome:  0,
						TotalExpense: 800,
					},
					{
						Category:     "Groceries",
						TotalIncome:  0,
						TotalExpense: 1500,
					},
					{
						Category:     "Rent",
						TotalIncome:  0,
						TotalExpense: 9000,
					},
					{
						Category:     "Salary",
						TotalIncome:  50000,
						TotalExpense: 0,
					},
				},
				err: nil,
			},
		},
		{
			name: "Successfl:  Total Income, Total Expenses, Net Balance with filter greater than or equal to startDate",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "",
				},
				filters: map[string]string{
					"start_date": "2026-04-03",
				},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"category", "total_income", "total_expence"}).
						AddRow("Entertainment", 0, 300).
						AddRow("Food", 0, 560).
						AddRow("Freelance", 1250, 50),
				)
			},
			want: want{
				response: []models.FinancialSummaryResponse{
					{
						Category:     "Entertainment",
						TotalIncome:  0,
						TotalExpense: 300,
					},
					{
						Category:     "Food",
						TotalIncome:  0,
						TotalExpense: 560,
					},
					{
						Category:     "Freelance",
						TotalIncome:  1250,
						TotalExpense: 50,
					},
				},
				err: nil,
			},
		},
		{
			name: "Successful:  Total Income, Total Expenses, Net Balance with filter between startDate and endDate",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "",
				},
				filters: map[string]string{
					"start_date": "2026-04-2",
					"end_date":   "2026-04-03",
				},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"category", "total_income", "total_expence"}).
						AddRow("Fuel", 0, 800).
						AddRow("Groceries", 0, 1500).
						AddRow("Rent", 0, 9000),
				)
			},
			want: want{
				response: []models.FinancialSummaryResponse{
					{
						Category:     "Fuel",
						TotalIncome:  0,
						TotalExpense: 800,
					},
					{
						Category:     "Groceries",
						TotalIncome:  0,
						TotalExpense: 1500,
					},
					{
						Category:     "Rent",
						TotalIncome:  0,
						TotalExpense: 9000,
					},
				},
				err: nil,
			},
		},
		{
			name: "Successful:  Total Income, Total Expenses, Net Balance with on transactionDate",
			args: args{
				request: &models.FinancialSummaryRequests{
					TransactionDate: "2026-04-04",
				},
				filters: map[string]string{},
			},

			mock: func() {
				mockDb.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					sqlmock.NewRows([]string{"category", "total_income", "total_expence"}).
						AddRow("Entertainment", 0, 300).
						AddRow("Freelance", 1250, 50),
				)
			},
			want: want{
				response: []models.FinancialSummaryResponse{
					{
						Category:     "Entertainment",
						TotalIncome:  0,
						TotalExpense: 300,
					},
					{
						Category:     "Freelance",
						TotalIncome:  1250,
						TotalExpense: 50,
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := fc.CategorySummary(tt.args.request, tt.args.filters)

			if tt.want.err != nil {
				assert.EqualError(t, tt.want.err, err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want.response, result)
			}

		})
	}
}
