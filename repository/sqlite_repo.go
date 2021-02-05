package repository

import (
	"database/sql"

	"github.com/artfunder/structs"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteRepo ...
type SQLiteRepo struct {
	db *sql.DB
}

// NewSQLiteRepo ...
func NewSQLiteRepo() *SQLiteRepo {
	repo := new(SQLiteRepo)
	db, err := sql.Open("sqlite3", "../users.db")
	if err != nil {
		panic(err)
	}
	setupTable(db)
	repo.db = db
	return repo
}

// GetAll ...
func (SQLiteRepo) GetAll() ([]structs.User, error) {
	return nil, nil
}

// GetOne ...
func (repo SQLiteRepo) GetOne(id int) (structs.User, error) {
	var user structs.User

	query := `SELECT id, firstname, lastname, email, username, password FROM users WHERE id = $1`
	err := repo.db.QueryRow(query, id).Scan(
		&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Username, &user.Password,
	)

	if err != nil {
		return structs.User{}, err
	}

	return user, nil
}

// Create ...
func (repo SQLiteRepo) Create(userInfo structs.User) (structs.User, error) {
	query := `INSERT INTO users (firstname, lastname, email, username, password) VALUES ($1, $2, $3, $4, $5)`
	result, err := repo.db.Exec(query,
		userInfo.Firstname, userInfo.Lastname, userInfo.Email, userInfo.Username, userInfo.Password,
	)
	if err != nil {
		return structs.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.User{}, err
	}

	return repo.GetOne(int(id))
}

func setupTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id				INTEGER PRIMARY KEY NOT NULL,
		firstname	TEXT,
		lastname	TEXT,
		email			TEXT NOT NULL UNIQUE,
		username	TEXT NOT NULL UNIQUE,
		password	TEXT NOT NULL
	);`)

	if err != nil {
		panic(err)
	}
}
