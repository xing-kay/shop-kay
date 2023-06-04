package user

import (
	"net/http"

	"shop-kay/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"

	"shop-kay/apps/app/api/internal/logic/user"
	"shop-kay/apps/app/api/internal/svc"
	"shop-kay/apps/app/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
