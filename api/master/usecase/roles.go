package masterUsecase

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
	masterRepository "belajar/api/master/repository"
	"belajar/package/manager"
	"context"

	"github.com/rs/zerolog/log"
)

type Roles struct {
	repo masterDomainInterface.RolesRepository
}

func NewRolesUsecase(mgr manager.Manager) masterDomainInterface.RolesUsecase {
	usecase := new(Roles)
	usecase.repo = masterRepository.NewRolesRepository(mgr.GetGorm())
	return usecase
}

func (uc *Roles) GetRoles(ctx context.Context) ([]*masterDomainEntity.Roles, error) {
	results, err := uc.repo.GetRoles(ctx)
	if err != nil {
		code := "[Usecase] GetRoles-1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}
