package gogen

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const (
	defaultPort = 8888
	etcDir      = "etc"
)

//go:embed etc.tpl
var etcTemplate string

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	if err := genEtcOnce(dir, "dev", cfg, api); err != nil {
		return err
	}
	if err := genEtcOnce(dir, "test", cfg, api); err != nil {
		return err
	}
	if err := genEtcOnce(dir, "pro", cfg, api); err != nil {
		return err
	}
	return nil
}

func genEtcOnce(dir, mode string, cfg *config.Config, api *spec.ApiSpec) error {
	host := "0.0.0.0"
	port := strconv.Itoa(defaultPort)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        fmt.Sprintf("config-%s.yaml", mode),
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: etcTemplate,
		data: map[string]string{
			"host": host,
			"port": port,
		},
	})
}
