package repositories

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/steevehook/expenses-rest-api/logging"
	"github.com/steevehook/expenses-rest-api/models"
)

// BoltDB buckets
var (
	expensesBucket    = []byte("expenses")
	expensesIDsBucket = []byte("expenses_ids")
)

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
	expenses := make([]models.Expense, 0)
	idsLookup := make([][]byte, 0)

	err := d.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(expensesIDsBucket)
		for _, uid := range ids {
			id := bucket.Get([]byte(uid))
			if len(id) == 0 {
				logging.Logger.Debug(fmt.Sprintf("record with id: %s was not found in db", uid))
				continue
			}
			idsLookup = append(idsLookup, id)
		}
		return nil
	})
	if err != nil {
		logging.Logger.Error("could not fetch uid:id pairs", zap.Error(err))
		return []models.Expense{}, err
	}

	err = d.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(expensesBucket)
		for _, id := range idsLookup {
			var expense models.Expense
			bs := bucket.Get([]byte(id))
			if e := json.Unmarshal(bs, &expense); e != nil {
				logging.Logger.Error("could not unmarshal expense when fetching by ids", zap.Error(err))
				return err
			}
			expenses = append(expenses, expense)
		}
		return nil
	})
	if err != nil {
		logging.Logger.Error("could not fetch expenses by ids", zap.Error(err))
		return []models.Expense{}, err
	}
	return expenses, nil
}

// CreateExpense creates a brand new expense and saves it into BoltDB
func (d BoltDriver) CreateExpense(title, currency string, price float64) error {
	var idLookup, uidLookup []byte
	err := d.boltDB.Update(func(tx *bolt.Tx) error {
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
		idLookup = idData
		uidLookup = []byte(id.String())
		return nil
	})
	if err != nil {
		logging.Logger.Error("could not create expense in db", zap.Error(err))
		return err
	}
	return d.setExpenseID(idLookup, uidLookup)
}

// UpdateExpense updates an existing expense and updates the record in BoltDB
func (d BoltDriver) UpdateExpense(id, title, currency string, price float64) error {
	var lookupID []byte
	err := d.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(expensesIDsBucket)
		intID := bucket.Get([]byte(id))
		if len(intID) == 0 {
			return models.ResourceNotFoundError{}
		}
		lookupID = intID
		return nil
	})
	if err != nil {
		logging.Logger.Error("could not fetch uid:id for update expense", zap.Error(err))
		return err
	}

	return d.boltDB.Update(func(tx *bolt.Tx) error {
		var modified bool
		bucket := tx.Bucket(expensesBucket)
		bs := bucket.Get(lookupID)
		if len(bs) == 0 {
			return models.ResourceNotFoundError{}
		}
		var expense models.Expense
		err := json.Unmarshal(bs, &expense)
		if err != nil {
			logging.Logger.Error("could not unmarshal expense for update in db", zap.Error(err))
			return err
		}

		if title != "" && title != expense.Title {
			expense.Title = title
			modified = true
		}
		if price > 0 && price != expense.Price {
			expense.Price = price
			modified = true
		}
		if currency != "" && currency != expense.Currency {
			expense.Currency = currency
			modified = true
		}
		if modified {
			expense.ModifiedAt = time.Now().UTC()
		}

		bs, err = json.Marshal(expense)
		if err != nil {
			logging.Logger.Error("could not marshal expense for update in db", zap.Error(err))
			return err
		}

		err = bucket.Put(lookupID, bs)
		if err != nil {
			logging.Logger.Error("could not update expense in db", zap.Error(err))
			return err
		}
		return nil
	})
}

// DeleteExpense deletes a list of expenses from BoltDB given a list of IDs
func (d BoltDriver) DeleteExpense(id []string) error {
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
		logging.Logger.Error("could not count total count of expenses", zap.Error(err))
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

func (d BoltDriver) setExpenseID(id []byte, uid []byte) error {
	return d.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(expensesIDsBucket)
		if err != nil {
			logging.Logger.Error("could not create or open expenses ids bucket", zap.Error(err))
			return err
		}
		err = bucket.Put(uid, id)
		if err != nil {
			logging.Logger.Error("could not save uid:id record in boltdb", zap.Error(err))
			return err
		}
		return nil
	})
}
