package masterDomainInterface

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	"context"
	"net/http"
)

type AccessHandler interface {
	GetAccess() http.Handler
}

type AccessUsecase interface {
	GetAccess(ctx context.Context) ([]*masterDomainEntity.Access, error)
}

type AccessRepository interface {
	GetAccess(ctx context.Context) ([]*masterDomainEntity.Access, error)
}
