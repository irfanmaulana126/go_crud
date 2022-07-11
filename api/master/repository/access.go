package masterRepository

import (
	"context"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	masterDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
)

type access struct {
	DB *gorm.DB
}

func NewAccessRepository(database *gorm.DB) masterDomainInterface.AccessRepository {
	repo := new(access)
	repo.DB = database

	return repo
}

func (repo *access) GetAccess(ctx context.Context) ([]*masterDomainEntity.Access, error) {
	results := []*masterDomainEntity.Access{}

	err := repo.DB.Where("is_active = ?", true).
		Order("order_item").
		Select("id", "code", "name", "description", "order_item").
		Find(&results).Error
	if err != nil {
		code := "[Repository] GetAccess-1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}
