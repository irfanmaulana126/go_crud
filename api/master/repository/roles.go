package masterRepository

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
	"context"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type roles struct {
	DB *gorm.DB
}

func NewRolesRepository(database *gorm.DB) masterDomainInterface.RolesRepository {
	repo := new(roles)
	repo.DB = database

	return repo
}

func (repo *roles) GetRoles(ctx context.Context) ([]*masterDomainEntity.Roles, error) {
	results := []*masterDomainEntity.Roles{}

	err := repo.DB.Where("is_active = ?", true).
		Order("order_item").
		Select("id", "code", "name", "description", "order_item").
		Find(&results).Error
	if err != nil {
		code := "[Repository] GetRoles-1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}
