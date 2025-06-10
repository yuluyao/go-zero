// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.0

package logic

import (
	"log/slog"

	"hello/internal/svc"
	"hello/internal/types"

	"github.com/gin-gonic/gin"
)

type HealthLogic struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
	logger *slog.Logger
}

func NewHealthLogic(c *gin.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		c:      c,
		svcCtx: svcCtx,
		logger: slog.With("m", "HealthLogic"),
	}
}

func (l *HealthLogic) Health() (resp *types.HealthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
