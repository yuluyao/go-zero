package dartgen

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/api/util"
)

//go:embed api.tpl
var apiTemplate string

func genApi(dir string, api *spec.ApiSpec) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dir+util.ToSnakeCase(api.Service.Name)+".dart", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()
	t := template.New("apiTemplate")
	t = t.Funcs(funcMap)
	tpl := apiTemplate
	t, err = t.Parse(tpl)
	if err != nil {
		return err
	}

	return t.Execute(file, api)
}
