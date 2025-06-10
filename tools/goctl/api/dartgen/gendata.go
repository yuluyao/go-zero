package dartgen

import (
	"bytes"
	_ "embed"
	"os"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

//go:embed entity.tpl
var entityTemplate string

type DartSpec struct {
	APISpec        *spec.ApiSpec
	InnerClassList []string
}

func genData(dir string, api *spec.ApiSpec) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dir+strings.ToLower("entity"+".dart"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.New("dataTemplate")
	t = t.Funcs(funcMap)
	tpl := entityTemplate
	t, err = t.Parse(tpl)
	if err != nil {
		return err
	}

	dartSpec, err := convertDataType(api)
	if err != nil {
		return err
	}

	return t.Execute(file, dartSpec)
}

func convertDataType(api *spec.ApiSpec) (*DartSpec, error) {
	var result DartSpec
	types := api.Types
	if len(types) == 0 {
		return &result, nil
	}

	for _, ty := range types {
		defineStruct, ok := ty.(spec.DefineStruct)
		if ok {
			for index, member := range defineStruct.Members {
				structMember, ok := member.Type.(spec.NestedStruct)
				if ok {
					defineStruct.Members[index].Type = spec.PrimitiveType{RawName: member.Name}
					t := template.New("dataTemplate")
					t = t.Funcs(funcMap)
					tpl := entityTemplate
					t, err := t.Parse(tpl)
					if err != nil {
						return nil, err
					}

					var innerClassSpec = &spec.ApiSpec{
						Types: []spec.Type{
							spec.DefineStruct{
								RawName: member.Name,
								Members: structMember.Members,
							},
						},
					}
					dartSpec, err := convertDataType(innerClassSpec)
					if err != nil {
						return nil, err
					}

					writer := bytes.NewBuffer(nil)
					err = t.Execute(writer, dartSpec)
					if err != nil {
						return nil, err
					}
					result.InnerClassList = append(result.InnerClassList, writer.String())
				} else {
					tp, err := specTypeToDart(member.Type)
					if err != nil {
						return nil, err
					}
					defineStruct.Members[index].Type = buildSpecType(member.Type, tp)
				}
			}
		}
	}
	result.APISpec = api

	return &result, nil
}
