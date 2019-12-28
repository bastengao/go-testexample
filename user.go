package testexample

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	return db
}

func registerUser(db *sql.DB, email string, mailer Mailer) error {
	err := createUser(db, email)
	if err != nil {
		return err
	}

	return mailer.Send(email, "Registration", "welcome")
}

func createUser(db *sql.DB, email string) error {
	sql := `
		INSERT INTO users (email, created_at, updated_at) values(?, ?, ?)
	`
	_, err := db.Exec(sql, email, time.Now(), time.Now())
	return err
}

func queryUser(db *sql.DB, id int64) (*User, error) {
	sql := `
	  SELECT id, email, created_at, updated_at from users where id = ?
	`
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
