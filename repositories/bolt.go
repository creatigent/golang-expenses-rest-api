package repositories

import (
	"github.com/boltdb/bolt"
	"go.uber.org/zap"

	"github.com/steevehook/expenses-rest-api/logging"
	"github.com/steevehook/expenses-rest-api/models"
)

// BoltDriver represents BoltDB repository driver
type BoltDriver struct {
	boltDB *bolt.DB
}

func NewBoltDriver(filename string) (*BoltDriver, error) {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		logging.Logger.Error("could not create file database", zap.Error(err))
		return nil, err
	}

	driver := &BoltDriver{
		boltDB: db,
	}
	return driver, nil
}

func (d BoltDriver) GetAllExpenses(page, size int) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (d BoltDriver) GetExpensesByIDs(ids []string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (d BoltDriver) CreateExpense(title, currency string, price float64) error {
	return nil
}

func (d BoltDriver) UpdateExpense(title, currency string, price float64) error {
	return nil
}

func (d BoltDriver) DeleteExpenses(ids []string) error {
	return nil
}

func (d BoltDriver) Stop() error {
	logging.Logger.Info("stopping file db server")
	err := d.boltDB.Close()
	if err != nil {
		return err
	}

	logging.Logger.Info("file db server successfully stopped")
	return nil
}
