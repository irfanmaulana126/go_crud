package masterUsecase

import (
	masterDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
	masterRepository "belajar/api/master/repository"
	"belajar/package/manager"
	"context"

	"github.com/rs/zerolog/log"
)

type Access struct {
	repo masterDomainInterface.AccessRepository
}

func NewAccessUsecase(mgr manager.Manager) masterDomainInterface.AccessUsecase {
	usecase := new(Access)
	usecase.repo = masterRepository.NewAccessRepository(mgr.GetGorm())
	return usecase
}

func (uc *Access) GetAccess(ctx context.Context) ([]*masterDomainEntity.Access, error) {
	results, err := uc.repo.GetAccess(ctx)
	if err != nil {
		code := "[Usecase] GetAccess-1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}
