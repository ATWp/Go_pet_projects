package store

import "github.com/ATWp/Go_pet_projects/REST_API/internal/app/model"

//UserRepository
type UserRepository struct {
	store *Store
}

//Create User
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// Find Email in DB
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err:= r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users where email = $1",
		email
	).Scan(
		&u.ID, 
		&u.Email, 
		&u.EncryptedPassword
	); err != nil{
		return nil, err
	}
	return u, nil
}
