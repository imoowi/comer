/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"strings"

	"github.com/imoowi/comer/utils/format"
	"github.com/spf13/cobra"
)

func (c *Comer) initApp(cmd *cobra.Command, args []string) error {
	appName, err := cmd.Flags().GetString(`app`)
	if err != nil {
		return err
	}
	if appName == `` {
		return fmt.Errorf(`pls input app, e.g. -a=student (请输入app,例如 -a=student)`)
	}
	moduleName, err := readModuleName()
	if err != nil {
		return err
	}

	swaggerTags, _ := cmd.Flags().GetString(`swaggerTags`)
	if swaggerTags == `` {
		swaggerTags = appName
	}
	handlerName, _ := cmd.Flags().GetString(`controller`)
	if handlerName == `` {
		handlerName = appName
	}
	serviceName, _ := cmd.Flags().GetString(`service`)
	if serviceName == `` {
		serviceName = handlerName
	}
	modelName, _ := cmd.Flags().GetString(`model`)
	if modelName == `` {
		modelName = serviceName
	}

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
			`./apps/apps.go`: tplUri + `templates/v1/apps/apps.tmpl`,
			`./apps/` + strings.ToLower(appName) + `/router.go`:                                                   tplUri + `templates/v1/apps/genapp/router.tmpl`,
			`./apps/` + strings.ToLower(appName) + `/handlers/` + format.Camel2Snake(handlerName) + `.handler.go`: tplUri + `templates/v1/apps/genapp/handler.tmpl`,
			`./apps/` + strings.ToLower(appName) + `/migrates/` + format.Camel2Snake(modelName) + `.migrate.go`:   tplUri + `templates/v1/apps/genapp/migrate.tmpl`,
			`./apps/` + strings.ToLower(appName) + `/models/` + format.Camel2Snake(modelName) + `.model.go`:       tplUri + `templates/v1/apps/genapp/model.tmpl`,
			`./apps/` + strings.ToLower(appName) + `/repos/` + format.Camel2Snake(modelName) + `.repo.go`:         tplUri + `templates/v1/apps/genapp/repo.tmpl`,
			`./apps/` + strings.ToLower(appName) + `/services/` + format.Camel2Snake(serviceName) + `.service.go`: tplUri + `templates/v1/apps/genapp/service.tmpl`,
		},
	}
	c.tplAppData = buildAppTplData(moduleName, appName, handlerName, serviceName, modelName, swaggerTags)
	return nil
}
