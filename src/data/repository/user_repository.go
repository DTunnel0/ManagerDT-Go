package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/DTunnel0/ManagerDT-Go/src/domain/contracts"
	"github.com/DTunnel0/ManagerDT-Go/src/domain/entity"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func isRoot() bool {
	return os.Geteuid() == 0
}

func dbURI() string {
	uri := `./db.sqlite3`
	if isRoot() {
		uri = `/root/db.sqlite3`
	}

	db, err := filepath.Abs(uri)
	if err != nil {
		return uri
	}
	return db
}

type userSQLiteRepository struct {
	db *sql.DB
}

func NewUserSQLiteRepository() contracts.UserRepository {
	db, err := sql.Open("sqlite3", dbURI())
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		uuid TEXT DEFAULT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		connection_limit INTEGER DEFAULT 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP
	);`)
	if err != nil {
		log.Fatal(err)
	}

	return &userSQLiteRepository{db: db}
}

func (r *userSQLiteRepository) Save(ctx context.Context, user *entity.User) error {
	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO users (id, uuid, username, password, connection_limit, created_at, expires_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.UUID, user.Username, user.Password, user.Limit, user.CreatedAt, user.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *userSQLiteRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entity.User, 0)
	for rows.Next() {
		user := &entity.User{}
		err = rows.Scan(&user.ID, &user.UUID, &user.Username, &user.Password, &user.Limit, &user.CreatedAt, &user.ExpiresAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userSQLiteRepository) Delete(ctx context.Context, user ...*entity.User) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "DELETE FROM users WHERE id =?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, u := range user {
		_, err = stmt.ExecContext(ctx, u.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *userSQLiteRepository) ChangePassword(ctx context.Context, user *entity.User) error {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE users SET password =? WHERE id =?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}
