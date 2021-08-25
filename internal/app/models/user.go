package models

import (
	"database/sql/driver"

	"github.com/enorith/authenticate"
	"github.com/enorith/enorith/internal/pkg/database"
	"github.com/enorith/http/content"
	"gorm.io/gorm"
)

type Password []byte

// Value returns a driver Value.
// Value must not panic.
func (p Password) Value() (driver.Value, error) {
	return []byte(p), nil
}

type User struct {
	content.Request `gorm:"-" json:"-"`
	ID              int64             `gorm:"column:id;primaryKey" json:"id"`
	Name            string            `gorm:"column:name" json:"name" input:"name" validate:"required"`
	Username        string            `gorm:"column:username" json:"username" input:"username" validate:"required"`
	CreatedAt       database.Datetime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       database.Datetime `gorm:"column:updated_at" json:"updated_at"`
	Password        Password          `gorm:"column:password" json:"-" input:"password"`
}

func (u User) UserIdentifier() authenticate.UserIdentifier {
	return authenticate.Identifier(u.ID)
}

type UserProvider struct {
	db *gorm.DB
}

func (up UserProvider) FindUserById(id authenticate.UserIdentifier) (authenticate.User, error) {
	var user User
	e := up.db.Where("id = ?", id.Int64()).Find(&user).Error
	return user, e
}

func NewUserProvider(db *gorm.DB) *UserProvider {
	return &UserProvider{db: db}
}
