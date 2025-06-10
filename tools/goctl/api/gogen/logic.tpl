// Code scaffolded by goctl. Safe to edit.
// goctl {{.version}}

package {{.pkgName}}

import (
	"log/slog"

	{{.imports}}
	
    "github.com/gin-gonic/gin"
)

type {{.logic}} struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
    logger *slog.Logger
}

{{if .hasDoc}}{{.doc}}{{end}}
func New{{.logic}}(c *gin.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		c:      c,
		svcCtx: svcCtx,
		logger: slog.With("m", "{{.logic}}"),
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
