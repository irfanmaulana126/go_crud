package masterHandler

import (
	"net/http"

	jsonHelper "belajar/helper/json"
	manager "belajar/package/manager"

	"github.com/rs/zerolog/log"

	permissionsDomainEntity "belajar/api/master/domain/entity"
	masterDomainInterface "belajar/api/master/domain/interface"
	masterUsecase "belajar/api/master/usecase"
)

type Permissions struct {
	Usecase masterDomainInterface.PermissionsUsecase
}

func NewPermissionsHandler(mgr manager.Manager) masterDomainInterface.PermissionsHandler {
	handler := new(Permissions)
	handler.Usecase = masterUsecase.NewPermissionsUsecase(mgr)

	return handler
}
func (h *Permissions) GetPermissions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := &permissionsDomainEntity.PermissionsReq{
			RoleCode: r.FormValue("role_code"),
		}

		results, err := h.Usecase.GetPermissions(ctx, req)
		if err != nil {
			code := "[Handler] GetPermissions-2"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponseV2(w, r, false, http.StatusInternalServerError, err.Error(), code)
			return
		}

		jsonHelper.SuccessResponseV2(w, r, true, "", results, nil)
	})
}
