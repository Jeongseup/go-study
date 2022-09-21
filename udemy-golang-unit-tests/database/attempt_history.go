package database

import (
	"database/sql"
	"golang-unit-tests/entity"
)

type AttemptHistory interface {
	IncrementFailure(user entity.User) error
	CountFailures(user entity.User) (int, error)
}

type AttemptHistoryRepository struct {
	db *sql.DB
}

func NewAttemptHistoryRepository(db *sql.DB) *AttemptHistoryRepository {
	return &AttemptHistoryRepository{
		db: db,
	}
}

func (dc *AttemptHistoryRepository) IncrementFailure(user entity.User) error {
	sql := "INSERT INTO attempt_history (user_id) VALUES (?)"
	_, err := dc.db.Exec(sql, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (dc *AttemptHistoryRepository) CountFailures(user entity.User) (count int, err error) {

	return 0, nil
}
