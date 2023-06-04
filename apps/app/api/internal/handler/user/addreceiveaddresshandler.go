package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"shop-kay/apps/app/api/internal/logic/user"
	"shop-kay/apps/app/api/internal/svc"
	"shop-kay/apps/app/api/internal/types"
	"shop-kay/pkg/result"
)

func AddReceiveAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReceiveAddressAddReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewAddReceiveAddressLogic(r.Context(), svcCtx)
		resp, err := l.AddReceiveAddress(&req)
		result.HttpResult(r, w, resp, err)
	}
}
