package db

import "github.com/xarick/golang-kafka-example/models"

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
