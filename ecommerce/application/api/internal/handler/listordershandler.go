package handler

import (
	"net/http"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/logic"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListOrdersRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListOrdersLogic(r.Context(), svcCtx)
		resp, err := l.ListOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
