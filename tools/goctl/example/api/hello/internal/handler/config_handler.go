// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.0

package handler

import (
	"hello/internal/logic"
	"hello/internal/response"
	"hello/internal/svc"

	"github.com/gin-gonic/gin"
)

func ConfigHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logic.NewConfigLogic(c, svcCtx)
		resp, err := l.Config()
		response.Response(c, resp, err)
	}
}
