// Code scaffolded by goctl. Safe to edit.
// goctl {{.version}}

package {{.PkgName}}

import (
    {{if .HasRequest}}"net/http"{{end}}

	{{.ImportPackages}}
    
    "github.com/gin-gonic/gin"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		{{if .HasRequest}}var req types.{{.RequestType}}
        if err := c.ShouldBind(&req); err != nil {
			response.ResponseInvalid(c, http.StatusBadRequest, err)
            return
        }
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(c, svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        {{if .HasResp}}response.Response(c, resp, err){{else}}response.Response(c, nil, err){{end}}
	}
}
