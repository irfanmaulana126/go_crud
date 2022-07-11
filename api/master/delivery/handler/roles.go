package masterHandler

import (
	masterDomainInterface "belajar/api/master/domain/interface"
	jsonHelper "belajar/helper/json"
	"net/http"

	masterUsecase "belajar/api/master/usecase"
	errorHelper "belajar/helper/error"
	manager "belajar/package/manager"

	"github.com/rs/zerolog/log"
)

type Roles struct {
	Usecase    masterDomainInterface.RolesUsecase
	ErrMessage errorHelper.ErrorMessage
}

func NewRolesHandler(mgr manager.Manager) masterDomainInterface.RolesHandler {
	handler := new(Roles)
	handler.Usecase = masterUsecase.NewRolesUsecase(mgr)

	return handler
}

func (h *Roles) GetRoles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		results, err := h.Usecase.GetRoles(ctx)
		if err != nil {
			code := "[Handler] GetRoles-1"
			log.Error().Err(err).Msg(code)
			errMsg := h.ErrMessage.GetMessage("5101")
			jsonHelper.ErrorResponseV4(w, r, false, http.StatusBadRequest, errMsg)
			return
		}

		jsonHelper.SuccessResponseV4(w, r, true, jsonHelper.Success, results, nil)
	})
}
