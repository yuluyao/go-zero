// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.0

package handler

import (
	"hello/internal/logic"
	"hello/internal/response"
	"hello/internal/svc"

	"github.com/gin-gonic/gin"
)

func HealthHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logic.NewHealthLogic(c, svcCtx)
		resp, err := l.Health()
		response.Response(c, resp, err)
	}
}
