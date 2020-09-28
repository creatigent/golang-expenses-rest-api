package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/steevehook/expenses-rest-api/logging"
	"github.com/steevehook/expenses-rest-api/models"
)

// MariaDBDriver represents MariaDB repository driver
type MariaDBDriver struct {
	mariaDB *sql.DB
}

// NewMariaDBDriver creates a new instance of MariaDB database
func NewMariaDBDriver(dbURL string) (*MariaDBDriver, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		logging.Logger.Error("could not open mariadb database", zap.Error(err))
		return nil, err
	}
	driver := &MariaDBDriver{
		mariaDB: db,
	}
	return driver, nil
}

// GetAllExpenses fetches all expenses with pagination possibilities from MariaDB
func (d MariaDBDriver) GetAllExpenses(page, size int) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

// GetExpensesByIDs fetches a list of expenses by a given list of IDs from MariaDB
func (d MariaDBDriver) GetExpensesByIDs(ids []string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

// CreateExpense creates a brand new expense and saves it into MariaDB
func (d MariaDBDriver) CreateExpense(title, currency string, price float64) error {
	return nil
}

// UpdateExpense updates an existing expense and updates the record in MariaDB
func (d MariaDBDriver) UpdateExpense(id, title, currency string, price float64) error {
	return nil
}

// DeleteExpense deletes a given expense from MariaDB given a list of IDs
func (d MariaDBDriver) DeleteExpense(id string) error {
	return nil
}

// Count fetches the total count from expenses table from MariaDB
func (d MariaDBDriver) Count() (int, error) {
	return 0, nil
}

// Close closes the MariaDB database
func (d MariaDBDriver) Close() error {
	logging.Logger.Info("stopping mariadb server")
	err := d.mariaDB.Close()
	if err != nil {
		return err
	}

	logging.Logger.Info("mariadb server successfully stopped")
	return nil
}
