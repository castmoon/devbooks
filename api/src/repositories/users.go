package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func UserRepositoryFactory(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("insert into users (name, nickname, email, password) values(?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(insertedId), nil
}
