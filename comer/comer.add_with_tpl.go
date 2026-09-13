/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"os"
	"strings"

	"github.com/imoowi/comer/utils/format"
	"github.com/imoowi/comer/utils/myfile"
	"github.com/spf13/cobra"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

type TplVar struct {
	ModuleName     string `json:"module_name"`
	ControllerName string `json:"controller_name"`
	ServiceName    string `json:"service_name"`
	ModelName      string `json:"model_name"`
	SwaggerTags    string `json:"swagger_tags"`
}
type DirAndTpl struct {
	Dir string `json:"dir"`
	Tpl string `json:"tpl"`
}
type TplSetting struct {
	Vars       []TplVar    `json:"var"`
	Controller []DirAndTpl `json:"controller"`
	Service    []DirAndTpl `json:"service"`
	Repo       []DirAndTpl `json:"repo"`
	Model      []DirAndTpl `json:"model"`
	Migrate    []DirAndTpl `json:"migrate"`
	Router     []DirAndTpl `json:"router"`
}

func (c *Comer) GenAppWithTpl(cmd *cobra.Command, args []string) error {
	c.showLogo()
	tpl, err := cmd.Flags().GetString(`tpl`)
	if err != nil {
		return err
	}
	fmt.Println(`comer-templates dir is: `, tpl)
	setting := tpl + `/setting.json5`
	if !myfile.IsFileExist(setting) {
		return fmt.Errorf(`[.comer-templates/setting.json5 dose not exist]`)
	}
	data, err := os.ReadFile(setting)
	if err != nil {
		return fmt.Errorf(`[Read .comer-templates/setting.json5 failed] %w`, err)
	}
	var tplSetting []TplSetting
	if err = json5.Unmarshal(data, &tplSetting); err != nil {
		return fmt.Errorf(`[Parse .comer-templates/setting.json5 failed] %w`, err)
	}

	for _, v := range tplSetting {
		dirs := make([]string, 0)
		for _, vv := range v.Controller {
			dirs = append(dirs, vv.Dir)
		}
		for _, vv := range v.Service {
			dirs = append(dirs, vv.Dir)
		}
		for _, vv := range v.Model {
			dirs = append(dirs, vv.Dir)
		}
		for _, vv := range v.Repo {
			dirs = append(dirs, vv.Dir)
		}
		for _, vv := range v.Migrate {
			dirs = append(dirs, vv.Dir)
		}
		for _, vv := range v.Router {
			dirs = append(dirs, vv.Dir)
		}
		if err = c.generateDirs(dirs); err != nil {
			return err
		}

		for _, vv := range v.Vars {
			if vv.ServiceName == `` {
				vv.ServiceName = vv.ControllerName
			}
			if vv.ModelName == `` {
				vv.ModelName = vv.ServiceName
			}
			files := make(map[string]string)
			for _, vvv := range v.Controller {
				files[`./`+strings.ToLower(vvv.Dir)+`/`+format.Camel2Snake(vv.ControllerName)+`.controller.go`] = tpl + `/` + vvv.Tpl
			}
			for _, vvv := range v.Service {
				files[`./`+strings.ToLower(vvv.Dir)+`/`+format.Camel2Snake(vv.ServiceName)+`.service.go`] = tpl + `/` + vvv.Tpl
			}
			for _, vvv := range v.Migrate {
				files[`./`+strings.ToLower(vvv.Dir)+`/`+format.Camel2Snake(vv.ModelName)+`.migrate.go`] = tpl + `/` + vvv.Tpl
			}
			for _, vvv := range v.Repo {
				files[`./`+strings.ToLower(vvv.Dir)+`/`+format.Camel2Snake(vv.ModelName)+`.repo.go`] = tpl + `/` + vvv.Tpl
			}
			for _, vvv := range v.Model {
				files[`./`+strings.ToLower(vvv.Dir)+`/`+format.Camel2Snake(vv.ModelName)+`.model.go`] = tpl + `/` + vvv.Tpl
			}
			for _, vvv := range v.Router {
				files[`./`+strings.ToLower(vvv.Dir)+`/`+format.Camel2Snake(vv.ControllerName)+`.router.go`] = tpl + `/` + vvv.Tpl
			}

			tplAppData := buildAppTplData(vv.ModuleName, vv.ControllerName, vv.ControllerName, vv.ServiceName, vv.ModelName, vv.SwaggerTags)
			if err = c.generateFiles(files, tplAppData); err != nil {
				return err
			}
		}
	}
	return nil
}
