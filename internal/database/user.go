package database

import (
	"fidelizou-go/internal/entities/models"
)

func (r *repository) GetUser(email string) (models.User, error) {
	var user models.User

	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *repository) UpdateUserRole(newRole int8, email string) error {

	_, err := r.db.Exec("UPDATE users SET role = $1 WHERE email = $2 ", newRole, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateUser(user models.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (email, role) VALUES (:Email, $2)", user)
	if err != nil {
		return err
	}
	return nil
}
