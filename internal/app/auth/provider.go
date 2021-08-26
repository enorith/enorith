package auth

import (
	"github.com/enorith/authenticate"
	"github.com/enorith/enorith/internal/app/models"
	"gorm.io/gorm"
)

type UserProvider struct {
	db *gorm.DB
}

func (up UserProvider) FindUserById(id authenticate.UserIdentifier) (authenticate.User, error) {
	var user models.User
	e := up.db.Where("id = ?", id.Int64()).Find(&user).Error
	return user, e
}

func NewUserProvider(db *gorm.DB) *UserProvider {
	return &UserProvider{db: db}
}
