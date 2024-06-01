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

func (r *repository) ToggleUserRole(email string) error {
	_, err := r.db.Exec("UPDATE users SET role = NOT role WHERE email = $1", email)
	if err != nil {
		return err
	}
	return nil
}
