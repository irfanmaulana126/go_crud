package masterDomainInterface

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	"context"
	"net/http"
)

type RolesHandler interface {
	GetRoles() http.Handler
}

type RolesUsecase interface {
	GetRoles(ctx context.Context) ([]*masterDomainEntity.Roles, error)
}
type RolesRepository interface {
	GetRoles(ctx context.Context) ([]*masterDomainEntity.Roles, error)
}
