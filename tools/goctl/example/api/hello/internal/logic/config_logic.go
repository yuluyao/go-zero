// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.0

package logic

import (
	"log/slog"

	"hello/internal/svc"
	"hello/internal/types"

	"github.com/gin-gonic/gin"
)

type ConfigLogic struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
	logger *slog.Logger
}

func NewConfigLogic(c *gin.Context, svcCtx *svc.ServiceContext) *ConfigLogic {
	return &ConfigLogic{
		c:      c,
		svcCtx: svcCtx,
		logger: slog.With("m", "ConfigLogic"),
	}
}

func (l *ConfigLogic) Config() (resp *types.ConfigResp, err error) {
	// todo: add your logic here and delete this line

	return
}
