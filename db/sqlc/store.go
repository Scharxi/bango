package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store SQLStore provides all functions to execute SQL queries and transactions
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return Store{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateBankAccountParams struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (store *Store) CreateBankAccountTx(ctx context.Context, args CreateBankAccountParams, accNumber int64) (int32, error) {
	var result int32
	err := store.execTx(ctx, func(queries *Queries) error {

		exists, err := store.DoesEmailExist(ctx, args.Email)
		if exists {
			return fmt.Errorf("email already exists")
		}

		exists, err = store.DoesPhoneNumberExist(ctx, args.PhoneNumber)
		if exists {
			return fmt.Errorf("phone number already exists")
		}

		holderId, err := store.CreateAccountHolder(ctx, CreateAccountHolderParams{
			FirstName: args.FirstName,
			LastName:  args.LastName,
			Email:     args.Email,
			Phone:     args.PhoneNumber,
			Address:   args.Address,
		})
		if err != nil {
			return err
		}

		result, err = store.CreateAccount(ctx, CreateAccountParams{
			AccountHolderID: holderId,
			AccountNumber:   accNumber,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

func (store *Store) OpenBankAccountTx(ctx context.Context, holderId int32, accNum int64) (int32, error) {
	var res int32
	err := store.execTx(ctx, func(queries *Queries) error {

		exists, err := store.DoesAccountHolderExist(ctx, holderId)
		if !exists {
			return fmt.Errorf("holder does not exists")
		}

		res, err = store.CreateAccount(ctx, CreateAccountParams{
			AccountHolderID: holderId,
			AccountNumber:   accNum,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return res, err
}
