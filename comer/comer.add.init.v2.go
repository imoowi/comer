/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/imoowi/comer/utils/format"
	"github.com/spf13/cobra"
)

func (c *Comer) initAppV2(cmd *cobra.Command, args []string) bool {
	moduleFile := `go.mod`
	_, gErr := os.Stat(moduleFile)
	if os.IsNotExist(gErr) {
		log.Println(`项目根目录下没有 go.mod 文件`)
		return false
	}
	file, err := os.OpenFile(moduleFile, os.O_RDWR, 0544)
	if err != nil {
		fmt.Printf("File open failed! err: %v\n", err)
		return false
	}
	reader := bufio.NewReader(file)
	_moduleName := ``
	for {
		line, err := reader.ReadString('\n') // 依次读一行
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("File raed failed! err: %v\n", err)
			return false
		}
		if strings.Contains(line, `module`) {
			_moduleName = line
			break
		}
	}
	file.Close()
	// fmt.Printf(`module=%v`, _moduleName)
	ModuleName := strings.Replace(_moduleName, "module ", "", -1)
	ModuleName = strings.Replace(ModuleName, "\r", "", -1)
	ModuleName = strings.Replace(ModuleName, "\n", "", -1)
	ModuleName = strings.Replace(ModuleName, "\n\r", "", -1)
	ModuleName = strings.Replace(ModuleName, "\r\n", "", -1)
	// fmt.Printf(`ModuleName=%v`, ModuleName)

	controllerName, _ := cmd.Flags().GetString(`controller`)
	if controllerName == `` {
		fmt.Println(`pls input controller name, e.g. -c=student (请输入控制器名,例如 -c=student)`)
		return false
	}
	SwaggerTags, _ := cmd.Flags().GetString(`swaggerTags`)
	if SwaggerTags == `` {
		SwaggerTags = controllerName
	}
	// fmt.Println(`SwaggerTags=`, SwaggerTags)

	// fmt.Println(`controllerName=`, controllerName)
	ServiceName, _ := cmd.Flags().GetString(`service`)
	if ServiceName == `` {
		ServiceName = controllerName
	}
	// fmt.Println(`ServiceName=`, ServiceName)
	ModelName, _ := cmd.Flags().GetString(`model`)
	if ModelName == `` {
		ModelName = ServiceName
	}
	// fmt.Println(`ModelName=`, ModelName)
	tplUri := ``
	c.App = &App{
		dirs: []string{
			`./internal`,
			`./internal` + `/controllers`,
			`./internal` + `/migrates`,
			`./internal` + `/models`,
			`./internal` + `/repos`,
			`./internal` + `/router`,
			`./internal` + `/services`,
		},
		files: map[string]string{
			`./internal` + `/router/` + format.Camel2Snake(controllerName) + `.router.go`:          tplUri + `templates/v2/internal/apps/router.tmpl`,
			`./internal` + `/controllers/` + format.Camel2Snake(controllerName) + `.controller.go`: tplUri + `templates/v2/internal/apps/controller.tmpl`,
			`./internal` + `/migrates/` + format.Camel2Snake(ModelName) + `.migrate.go`:            tplUri + `templates/v2/internal/apps/migrate.tmpl`,
			`./internal` + `/models/` + format.Camel2Snake(ModelName) + `.model.go`:                tplUri + `templates/v2/internal/apps/model.tmpl`,
			`./internal` + `/models/` + format.Camel2Snake(ModelName) + `.filter.go`:               tplUri + `templates/v2/internal/apps/filter.tmpl`,
			`./internal` + `/repos/` + format.Camel2Snake(ModelName) + `.repo.go`:                  tplUri + `templates/v2/internal/apps/repo.tmpl`,
			`./internal` + `/services/` + format.Camel2Snake(ServiceName) + `.service.go`:          tplUri + `templates/v2/internal/apps/service.tmpl`,
		},
	}
	c.tplAppData = map[string]any{
		`ModuleName`:           ModuleName,
		`moduleName`:           strings.ToLower(ModuleName),
		`AppName`:              format.FirstUpper(controllerName),
		`appName`:              strings.ToLower(controllerName),
		`ControllerName`:       format.FirstUpper(controllerName),
		`lControllerName`:      format.FirstLower(controllerName),
		`controllerName`:       strings.ToLower(controllerName),
		`controller_name`:      format.Camel2Snake(controllerName),
		`controller-name`:      format.Camel2Dash(controllerName),
		`controllerName2Dash`:  format.Camel2Dash(controllerName),
		`controllerName2Snake`: format.Camel2Snake(controllerName),
		`ServiceName`:          format.FirstUpper(ServiceName),
		`serviceName`:          strings.ToLower(ServiceName),
		`ModelName`:            format.FirstUpper(ModelName),
		`modelName`:            strings.ToLower(ModelName),
		`model_name`:           format.Camel2Snake(ModelName),
		`model-name`:           format.Camel2Dash(ModelName),
		`SwaggerTags`:          SwaggerTags,
	}
	return true
}
