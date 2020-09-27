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

func (d MariaDBDriver) GetAllExpenses(page, size int) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (d MariaDBDriver) GetExpensesByIDs(ids []string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (d MariaDBDriver) CreateExpense(title, currency string, price float64) error {
	return nil
}

func (d MariaDBDriver) UpdateExpense(title, currency string, price float64) error {
	return nil
}

func (d MariaDBDriver) DeleteExpenses(ids []string) error {
	return nil
}

func (d MariaDBDriver) Stop() error {
	logging.Logger.Info("stopping mariadb server")
	err := d.mariaDB.Close()
	if err != nil {
		return err
	}

	logging.Logger.Info("mariadb server successfully stopped")
	return nil
}
