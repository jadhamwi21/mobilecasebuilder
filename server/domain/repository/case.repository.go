package repository

import (
	"github.com/jadhamwi21/mobilecasebuilder/domain/entity"
	"github.com/jadhamwi21/mobilecasebuilder/models"
)

type CaseRepository interface {
	SaveCase(*entity.Case) (map[string]string, error)
	SetCaseStatus(string, models.CaseStatus) error
	GetCaseById(string) (*entity.Case, error)
	GetCasesByOrderId(string) ([]*entity.Case, error)
}
