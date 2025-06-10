// Code scaffolded by goctl. Safe to edit.
// goctl {{.version}}

package middleware

import (
    {{.importPackages}}

    "github.com/gin-gonic/gin"
)

type {{.name}} struct {
}

func New{{.name}}() *{{.name}} {
	return &{{.name}}{}
}

func (m *{{.name}})Handle(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		c.Next()
	}
}
