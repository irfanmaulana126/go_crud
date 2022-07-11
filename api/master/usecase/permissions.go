package masterUsecase

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
	masterRepository "belajar/api/master/repository"
	"belajar/package/manager"
	"context"

	"github.com/rs/zerolog/log"
)

type Permissions struct {
	repo masterDomainInterface.PermissionsRepository
}

func NewPermissionsUsecase(mgr manager.Manager) masterDomainInterface.PermissionsUsecase {
	usecase := new(Permissions)
	usecase.repo = masterRepository.NewPermissionsRepository(mgr.GetGorm())
	return usecase
}

func (uc *Permissions) GetPermissions(ctx context.Context, req *masterDomainEntity.PermissionsReq) ([]*masterDomainEntity.Permissions, error) {
	results, err := uc.repo.GetPermissions(ctx, req)
	if err != nil {
		code := "[Usecase] GetPermissions-1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}
