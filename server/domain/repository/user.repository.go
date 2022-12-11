package repository

import (
	"github.com/jadhamwi21/mobilecasebuilder/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (map[string]string, error)
	GetUserByEmail(string) (*entity.User, error)
}
