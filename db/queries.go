package db

import "github.com/xarick/golang-kafka-example/models"

type UserDB struct {
	ID        int    `db:"id" json:"id"`
	FIO       string `db:"fio" json:"fio"`
	Username  string `db:"username" json:"username"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

func RegisterUser(fio, username, hashedPassword string) error {
	query := "INSERT INTO users (fio, username, password) VALUES ($1, $2, $3)"
	_, err := DB.Exec(query, fio, username, hashedPassword)
	return err
}

func GetUserByName(username string) (*models.User, error) {
	query := "SELECT fio, username, password FROM users WHERE username = $1"
	var user models.User
	err := DB.Get(&user, query, username)
	return &user, err
}

func GetUsers() ([]UserDB, error) {
	var users []UserDB
	query := "SELECT id, fio, username, created_at FROM users"
	err := DB.Select(&users, query)
	return users, err
}
