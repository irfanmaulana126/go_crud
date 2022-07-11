package masterDomainInterface

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	"context"
	"net/http"
)

type PermissionsHandler interface {
	GetPermissions() http.Handler
}

type PermissionsUsecase interface {
	GetPermissions(ctx context.Context, req *masterDomainEntity.PermissionsReq) ([]*masterDomainEntity.Permissions, error)
}

type PermissionsRepository interface {
	GetPermissions(ctx context.Context, req *masterDomainEntity.PermissionsReq) ([]*masterDomainEntity.Permissions, error)
}
