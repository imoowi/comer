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
	dryRun, err := cmd.Flags().GetBool(`dry-run`)
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

	fields, err := resolveFields(cmd)
	if err != nil {
		return err
	}
	searchColumn, _ := cmd.Flags().GetString(`searchColumn`)
	if searchColumn == `` {
		searchColumn = firstStringColumn(fields)
	}

	for _, v := range tplSetting {
		if len(v.Vars) == 0 {
			return fmt.Errorf(`[setting.json5 每个配置项至少需要一个 var]`)
		}
		dirs, err := appTplDirs(v)
		if err != nil {
			return err
		}
		if err := validateTplFiles(tpl, v); err != nil {
			return err
		}

		if dryRun {
			for _, d := range dirs {
				fmt.Printf("would create dir %s\n", d)
			}
		} else if err = c.generateDirs(dirs); err != nil {
			return err
		}

		for _, vv := range v.Vars {
			if vv.ServiceName == `` {
				vv.ServiceName = vv.ControllerName
			}
			if vv.ModelName == `` {
				vv.ModelName = vv.ServiceName
			}
			files := appTplFiles(tpl, v, vv)
			tplAppData := buildAppTplData(vv.ModuleName, vv.ControllerName, vv.ControllerName, vv.ServiceName, vv.ModelName, vv.SwaggerTags)
			tplAppData[`Fields`] = fields
			tplAppData[`SearchColumn`] = searchColumn
			tplAppData[`HasTime`] = hasTimeField(fields)
			tplAppData[`HasJSON`] = hasJSONField(fields)

			if dryRun {
				for f := range files {
					fmt.Printf("would create file %s\n", f)
				}
				continue
			}
			if err = c.generateFiles(files, tplAppData); err != nil {
				return err
			}
		}
	}
	return nil
}

// allDirAndTpl 收集 TplSetting 下所有 DirAndTpl。
func allDirAndTpl(v TplSetting) []DirAndTpl {
	out := make([]DirAndTpl, 0, len(v.Controller)+len(v.Service)+len(v.Repo)+len(v.Model)+len(v.Migrate)+len(v.Router))
	out = append(out, v.Controller...)
	out = append(out, v.Service...)
	out = append(out, v.Repo...)
	out = append(out, v.Model...)
	out = append(out, v.Migrate...)
	out = append(out, v.Router...)
	return out
}

// appTplDirs 收集需要创建的目录（去空、去重）。
func appTplDirs(v TplSetting) ([]string, error) {
	seen := map[string]bool{}
	var dirs []string
	for _, d := range allDirAndTpl(v) {
		dir := strings.TrimSpace(d.Dir)
		if dir == `` {
			return nil, fmt.Errorf(`[setting.json5 中存在空 dir]`)
		}
		if !seen[dir] {
			seen[dir] = true
			dirs = append(dirs, dir)
		}
	}
	return dirs, nil
}

// validateTplFiles 校验所有引用的模板文件存在。
func validateTplFiles(tpl string, v TplSetting) error {
	var missing []string
	for _, d := range allDirAndTpl(v) {
		if d.Tpl != `` && !myfile.IsFileExist(tpl+`/`+d.Tpl) {
			missing = append(missing, d.Tpl)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf(`[模板文件不存在: %s]`, strings.Join(missing, `, `))
	}
	return nil
}

// appTplFiles 构建某个 var 对应的文件路径到模板路径映射。
func appTplFiles(tpl string, v TplSetting, vv TplVar) map[string]string {
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
	return files
}
