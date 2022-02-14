package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Service - the struct for our users service
type Service struct {
	DB *gorm.DB
}

// User - defines the User model
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  bool      `json:"password" db:"password"`
	CreatedAt time.Time `json:"created" db:"created_at"`
	UpdatedAt time.Time `json:"updated" db:"updated_at"`
}

// UserService - the interface for our User service
type UserService interface {
	GetUser(ID uuid.UUID) (User, error)
	CreateUser(newUser User) (User, error)
	UpdateUser(ID uuid.UUID, user User) (User, error)
	DeleteUser(ID uuid.UUID) error
	GetAllUsers() ([]User, error)
}

// NewService - returns new users service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
