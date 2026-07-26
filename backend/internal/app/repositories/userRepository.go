package repositories

import (
	"backend/internal/app/database"
	"backend/internal/app/entities"
	"database/sql"
)

type User struct {
	database.BaseSQLRepository[entities.User]
}

func NewUserRepository(db *sql.DB) *User {
	return &User{
		BaseSQLRepository: database.BaseSQLRepository[entities.User]{DB: db},
	}
}

func mapUser(rows *sql.Row, u *entities.User) error {
	return rows.Scan(&u.ID, &u.Email, &u.Name, &u.PhoneNumber)
}

func mapUsers(rows *sql.Rows, u *entities.User) error {
	return rows.Scan(&u.ID, &u.Email, &u.Name, &u.PhoneNumber)
}

func mapUserWithPassword(rows *sql.Row, u *entities.User) error {
	return rows.Scan(&u.ID, &u.Email, &u.Name, &u.PhoneNumber, &u.Password)
}

func (r *User) FindByEmail(email string) (*entities.User, error) {
	return r.SelectSingle(
		mapUser,
		"SELECT u.id, u.email, u.name, u.phone_number FROM users u WHERE u.email = $1",
		email,
	)
}

func (r *User) FindByEmailWithPassword(email string) (*entities.User, error) {
	return r.SelectSingle(
		mapUserWithPassword,
		"SELECT u.id, u.email, u.name, u.phone_number, u.password FROM users u WHERE u.email = $1",
		email,
	)
}

func (r *User) FindByPhoneNumber(phoneNumber string) (*entities.User, error) {
	return r.SelectSingle(
		mapUser,
		"SELECT u.id, u.email, u.name, u.phone_number FROM users u WHERE u.phone_number = $1",
		phoneNumber,
	)
}

func (r *User) FindByID(id int) (*entities.User, error) {
	return r.SelectSingle(
		mapUser,
		"SELECT u.id, u.email, u.name, u.phone_number FROM users u WHERE u.id = $1",
		id,
	)
}

func (r *User) FindByIDWithPassword(id int) (*entities.User, error) {
	return r.SelectSingle(
		mapUserWithPassword,
		"SELECT u.id, u.email, u.name, u.phone_number, u.password FROM users u WHERE u.id = $1",
		id,
	)
}

func (r *User) GetAllUsers() ([]*entities.User, error) {
	return r.SelectMultiple(
		mapUsers,
		"SELECT u.id, u.email, u.name, u.phone_number FROM users u",
	)
}

func (r *User) CreateUser(user *entities.User) error {
	id, err := r.Insert(
		"INSERT INTO users (email, name, password, phone_number) VALUES ($1, $2, $3, $4)",
		user.Email, user.Name, user.Password, user.PhoneNumber,
	)
	user.ID = id
	return err
}

func (r *User) UpdateUser(user *entities.User) error {
	_, err := r.ExecuteQuery(
		"UPDATE users SET email = $1, name = $2, password = $3, phone_number = $4 WHERE id = $5",
		user.Email, user.Name, user.Password, user.PhoneNumber, user.ID,
	)

	return err
}

func (r *User) DeleteUser(id int) error {
	_, err := r.ExecuteQuery("DELETE FROM users WHERE id = $1", id)
	return err
}
