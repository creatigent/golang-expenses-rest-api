package repositories

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/steevehook/expenses-rest-api/logging"
	"github.com/steevehook/expenses-rest-api/models"
)

var expensesBucket = []byte("expenses")

// BoltDriver represents BoltDB repository driver
type BoltDriver struct {
	boltDB *bolt.DB
}

// NewBoltDriver creates a new instance of Bolt file database
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

// GetAllExpenses fetches all expenses with pagination possibilities from BoltDB
func (d BoltDriver) GetAllExpenses(page, pageSize int) ([]models.Expense, error) {
	expenses := make([]models.Expense, 0)
	err := d.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(expensesBucket)

		c := bucket.Cursor()
		for i := 1; i <= pageSize; i++ {
			keyLookup := []byte(strconv.Itoa(page*pageSize - pageSize + i))
			k, v := c.Seek(keyLookup)
			if string(k) == "" {
				break
			}

			var expense models.Expense
			err := json.Unmarshal(v, &expense)
			if err != nil {
				logging.Logger.Error("could not unmarshal expense when fetching all expenses", zap.Error(err))
				return err
			}
			if bytes.Equal(keyLookup, k) {
				expenses = append(expenses, expense)
			}
		}
		return nil
	})
	if err != nil {
		logging.Logger.Error("could not fetch all expenses from db", zap.Error(err))
		return []models.Expense{}, err
	}
	return expenses, nil
}

// GetExpensesByIDs fetches a list of expenses by a given list of IDs from BoldDB
func (d BoltDriver) GetExpensesByIDs(ids []string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

// CreateExpense creates a brand new expense and saves it into BoltDB
func (d BoltDriver) CreateExpense(title, currency string, price float64) error {
	return d.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(expensesBucket)
		if err != nil {
			logging.Logger.Error("could not create bucket", zap.Error(err))
			return err
		}

		next, err := bucket.NextSequence()
		if err != nil {
			logging.Logger.Error("could not get bucket next sequence", zap.Error(err))
			return err
		}
		idData := []byte(strconv.Itoa(int(next)))
		id := uuid.NewHash(md5.New(), uuid.NameSpaceURL, idData, 3)

		expense := models.Expense{
			ID:         id.String(),
			Title:      title,
			Currency:   currency,
			Price:      price,
			CreatedAt:  time.Now().UTC(),
			ModifiedAt: time.Now().UTC(),
		}

		bs, err := json.Marshal(expense)
		if err != nil {
			logging.Logger.Error("could not marshal json when creating expense")
			return err
		}
		err = bucket.Put(idData, bs)
		if err != nil {
			logging.Logger.Error("could not save expense in db")
			return err
		}
		logging.Logger.Info("successfully saved expense in db")
		return nil
	})
}

// UpdateExpense updates an existing expense and updates the record in BoltDB
func (d BoltDriver) UpdateExpense(title, currency string, price float64) error {
	return nil
}

// DeleteExpenses deletes a list of expenses from BoltDB given a list of IDs
func (d BoltDriver) DeleteExpenses(ids []string) error {
	return nil
}

// Count fetches the total count from expenses bucket from BoltDB
func (d BoltDriver) Count() (int, error) {
	var count int
	err := d.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(expensesBucket)
		count = bucket.Stats().KeyN
		return nil
	})
	if err != nil {
		logging.Logger.Error("could not count total count of expenses")
		return 0, err
	}
	return count, nil
}

// Close closes the BoltDB database
func (d BoltDriver) Close() error {
	logging.Logger.Info("stopping boltdb file database server")
	err := d.boltDB.Close()
	if err != nil {
		return err
	}

	logging.Logger.Info("file db server successfully stopped")
	return nil
}
