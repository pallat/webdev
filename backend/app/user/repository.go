package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) User(id string) (UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStmt := `select name,age from users where id = ?`
	row := s.db.QueryRowContext(ctx, sqlStmt, id)
	var user User
	err := row.Scan(&user.Name, &user.Age)

	return UserEntity{User: &user}, err
}

func (s *Storage) NewUser(c User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStmt := `insert into users(name, age) values(?,?)`
	_, err := s.db.ExecContext(ctx, sqlStmt, c.Name, c.Age)
	if err != nil {
		return errors.WithMessagef(err, "new a course: %s", c.Name)
	}
	return nil
}
