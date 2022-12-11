package repository

import (
	"github.com/jadhamwi21/mobilecasebuilder/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) error
	GetUserByEmail(string) (*entity.User, error)
}
