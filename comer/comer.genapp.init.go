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

func (c *Comer) initApp(cmd *cobra.Command, args []string) bool {

	appName, err := cmd.Flags().GetString(`app`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if appName == `` {
		fmt.Println(`pls input app, e.g. -a=student (请输入app,例如 -a=student)`)
		return false
	}
	moduleFile := `go.mod`
	_, gErr := os.Stat(moduleFile)
	if os.IsNotExist(gErr) {
		log.Println(`项目根目录下没有 go.mod 文件`)
		return false
	}
	/*
		data, err := ioutil.ReadFile(moduleFile)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		lines := strings.Split(string(data), "\n")
		ModuleName := strings.Replace(lines[0], "module ", "", -1)
		//*/
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
	SwaggerTags, _ := cmd.Flags().GetString(`swaggerTags`)
	if SwaggerTags == `` {
		SwaggerTags = appName
	}
	// fmt.Println(`SwaggerTags=`, SwaggerTags)
	HandlerName, _ := cmd.Flags().GetString(`handler`)
	if HandlerName == `` {
		HandlerName = appName
	}
	// fmt.Println(`HandlerName=`, HandlerName)
	ServiceName, _ := cmd.Flags().GetString(`service`)
	if ServiceName == `` {
		ServiceName = HandlerName
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
			`./apps`,
			`./apps/` + strings.ToLower(appName) + `/handlers`,
			`./apps/` + strings.ToLower(appName) + `/migrates`,
			`./apps/` + strings.ToLower(appName) + `/models`,
			`./apps/` + strings.ToLower(appName) + `/repos`,
			`./apps/` + strings.ToLower(appName) + `/services`,
		},
		files: map[string]string{
			`./apps/apps.go`: tplUri + `templates/v1/apps/apps.tpl`,
			`./apps/` + strings.ToLower(appName) + `/router.go`:                                                   tplUri + `templates/v1/apps/genapp/router.tpl`,
			`./apps/` + strings.ToLower(appName) + `/handlers/` + format.Camel2Snake(HandlerName) + `.handler.go`: tplUri + `templates/v1/apps/genapp/handler.tpl`,
			`./apps/` + strings.ToLower(appName) + `/migrates/` + format.Camel2Snake(ModelName) + `.migrate.go`:   tplUri + `templates/v1/apps/genapp/migrate.tpl`,
			`./apps/` + strings.ToLower(appName) + `/models/` + format.Camel2Snake(ModelName) + `.model.go`:       tplUri + `templates/v1/apps/genapp/model.tpl`,
			`./apps/` + strings.ToLower(appName) + `/repos/` + format.Camel2Snake(ModelName) + `.repo.go`:         tplUri + `templates/v1/apps/genapp/repo.tpl`,
			`./apps/` + strings.ToLower(appName) + `/services/` + format.Camel2Snake(ServiceName) + `.service.go`: tplUri + `templates/v1/apps/genapp/service.tpl`,
		},
	}
	c.tplAppData = map[string]any{
		`ModuleName`:        ModuleName,
		`moduleName`:        strings.ToLower(ModuleName),
		`AppName`:           format.FirstUpper(appName),
		`appName`:           strings.ToLower(appName),
		`HandlerName`:       format.FirstUpper(HandlerName),
		`lHandlerName`:      format.FirstLower(HandlerName),
		`handlerName`:       strings.ToLower(HandlerName),
		`handler_name`:      format.Camel2Snake(HandlerName),
		`handler-name`:      format.Camel2Dash(HandlerName),
		`handlerName2Dash`:  format.Camel2Dash(HandlerName),
		`handlerName2Snake`: format.Camel2Snake(HandlerName),
		`ServiceName`:       format.FirstUpper(ServiceName),
		`serviceName`:       strings.ToLower(ServiceName),
		`ModelName`:         format.FirstUpper(ModelName),
		`modelName`:         strings.ToLower(ModelName),
		`model_name`:        format.Camel2Snake(ModelName),
		`model-name`:        format.Camel2Dash(ModelName),
		`SwaggerTags`:       SwaggerTags,
	}
	return true
}
