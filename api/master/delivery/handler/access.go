package masterHandler

import (
	"net/http"

	jsonHelper "belajar/helper/json"
	manager "belajar/package/manager"

	"github.com/rs/zerolog/log"

	masterDomainInterface "belajar/api/master/domain/interface"
	masterUsecase "belajar/api/master/usecase"
)

type Access struct {
	Usecase masterDomainInterface.AccessUsecase
}

func NewAccessHandler(mgr manager.Manager) masterDomainInterface.AccessHandler {
	handler := new(Access)
	handler.Usecase = masterUsecase.NewAccessUsecase(mgr)

	return handler
}

func (h *Access) GetAccess() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		results, err := h.Usecase.GetAccess(ctx)
		if err != nil {
			code := "[Handler] GetAccess-1"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponseV2(w, r, false, http.StatusInternalServerError, err.Error(), code)
			return
		}

		jsonHelper.SuccessResponseV2(w, r, true, "", results, nil)
	})
}
