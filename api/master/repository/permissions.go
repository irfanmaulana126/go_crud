package masterRepository

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
	"context"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type permissions struct {
	DB *gorm.DB
}

func NewPermissionsRepository(database *gorm.DB) masterDomainInterface.PermissionsRepository {
	repo := new(permissions)
	repo.DB = database

	return repo
}

func (repo *permissions) GetPermissions(ctx context.Context, req *masterDomainEntity.PermissionsReq) ([]*masterDomainEntity.Permissions, error) {
	results := []*masterDomainEntity.Permissions{}
	var err error
	if req.RoleCode == "" {
		err = repo.DB.Where("is_active = ?", true).
			Order("order_item").
			Select("code", "name", "description", "order_item").
			Joins("wsapp_role_permission").
			Group("code").
			Find(&results).Error
	} else {
		err = repo.DB.Raw(`
			SELECT
				c.id,
				c.code,
				c.name,
				c.description
			FROM wsapp_user_roles a
			INNER JOIN wsapp_role_permission b ON a.id = b.role_id
			INNER JOIN wsapp_accesses c ON b.permission_id = c.id
			WHERE a.code = ? AND c.is_active = 1
			ORDER BY c.order_item
		`, req.RoleCode).
			Find(&results).
			Error
	}

	if err != nil {
		code := "[Repository] GetPermissions-1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}
