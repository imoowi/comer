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

type TplDir struct {
	Controller string `json:"controller"`
	Service    string `json:"service"`
	Repo       string `json:"repo"`
	Model      string `json:"model"`
	Migrate    string `json:"migrate"`
	Router     string `json:"router"`
}
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

func (c *Comer) GenAppWithTpl(cmd *cobra.Command, args []string) {
	c.showLogo()
	tpl, err := cmd.Flags().GetString(`tpl`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(`comer-templates dir is: `, tpl)
	setting := tpl + `/setting.json5`
	if !myfile.IsFileExist(setting) {
		fmt.Println(`Warning:[.comer-templates/setting.json5 dose not exist]`)
		return
	}
	data, err := os.ReadFile(setting)
	if err != nil {
		fmt.Println(`Warning:[Read .comer-templates/setting.json5 failed] `, err)
		return
	}
	var tplSetting []TplSetting
	err = json5.Unmarshal(data, &tplSetting)
	if err != nil {
		fmt.Println(`Warning:[Parse .comer-templates/setting.json5 failed] `, err.Error())
		return
	}
	// fmt.Println(tplSetting)
	for _, v := range tplSetting {

		dirs := make([]string, 0)
		for _, vv := range v.Controller {
			dirs = append(dirs, vv.Dir)
		}
		// dirs = append(dirs, v.Dir.Controller)
		for _, vv := range v.Service {
			dirs = append(dirs, vv.Dir)
		}
		// dirs = append(dirs, v.Dir.Service)
		for _, vv := range v.Model {
			dirs = append(dirs, vv.Dir)
		}
		// dirs = append(dirs, v.Dir.Model)
		for _, vv := range v.Repo {
			dirs = append(dirs, vv.Dir)
		}
		// dirs = append(dirs, v.Dir.Repo)
		for _, vv := range v.Migrate {
			dirs = append(dirs, vv.Dir)
		}
		// dirs = append(dirs, v.Dir.Migrate)
		for _, vv := range v.Router {
			dirs = append(dirs, vv.Dir)
		}
		// dirs = append(dirs, v.Dir.Router)
		c.generateDirs(dirs)

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
			// files := map[string]string{
			// 	`./` + strings.ToLower(v.Dir.Controller) + `/` + format.Camel2Snake(vv.ControllerName) + `.handler.go`: tpl + `/controller.tpl`,
			// 	`./` + strings.ToLower(v.Dir.Migrate) + `/` + format.Camel2Snake(vv.ModelName) + `.migrate.go`:         tpl + `/migrate.tpl`,
			// 	`./` + strings.ToLower(v.Dir.Model) + `/` + format.Camel2Snake(vv.ModelName) + `.model.go`:             tpl + `/model.tpl`,
			// 	`./` + strings.ToLower(v.Dir.Repo) + `/` + format.Camel2Snake(vv.ModelName) + `.repo.go`:               tpl + `/repo.tpl`,
			// 	`./` + strings.ToLower(v.Dir.Service) + `/` + format.Camel2Snake(vv.ServiceName) + `.service.go`:       tpl + `/service.tpl`,
			// 	`./` + strings.ToLower(v.Dir.Router) + `/` + format.Camel2Snake(vv.ControllerName) + `.router.go`:      tpl + `/router.tpl`,
			// }

			tplAppData := map[string]any{
				`ModuleName`:   vv.ModuleName,
				`moduleName`:   strings.ToLower(vv.ModuleName),
				`HandlerName`:  format.FirstUpper(vv.ControllerName),
				`lHandlerName`: format.FirstLower(vv.ControllerName),
				`handlerName`:  strings.ToLower(vv.ControllerName),
				// `handler_name`:      format.Camel2Dash(HandlerName),
				`handlerName2Dash`:  format.Camel2Dash(vv.ControllerName),
				`handlerName2Snake`: format.Camel2Snake(vv.ControllerName),
				`ServiceName`:       format.FirstUpper(vv.ServiceName),
				`serviceName`:       strings.ToLower(vv.ServiceName),
				`ModelName`:         format.FirstUpper(vv.ModelName),
				`modelName`:         strings.ToLower(vv.ModelName),
				`SwaggerTags`:       vv.SwaggerTags,
			}
			c.generateFiles(files, tplAppData)
		}

	}
}
