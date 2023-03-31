package user

import (
	"net/http"

	"boat/gateway/internal/logic/user"
	"boat/gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestErrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewTestErrLogic(r.Context(), svcCtx)
		resp, err := l.TestErr()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
