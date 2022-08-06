package repository

type UserRepository interface {
	GetUserByID(userID uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
