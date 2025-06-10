package gogen

import (
	_ "embed"
	"os"
	"path"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

//go:embed response.tpl
var responseTemplate string

func genResponse(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	responseFilename := "response" + ".go"
	filename := path.Join(dir, responseDir, responseFilename)
	os.Remove(filename)
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          responseDir,
		filename:        "response.go",
		templateName:    "responseTemplate",
		category:        category,
		templateFile:    responseTemplateFile,
		builtinTemplate: responseTemplate,
		data:            map[string]string{},
	})
}
