/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"

	"github.com/imoowi/comer/utils/format"
	"github.com/spf13/cobra"
)

func (c *Comer) initAppV2(cmd *cobra.Command, args []string) error {
	moduleName, err := readModuleName()
	if err != nil {
		return err
	}

	controllerName, err := cmd.Flags().GetString(`controller`)
	if err != nil {
		return err
	}
	if controllerName == `` {
		return fmt.Errorf(`pls input controller name, e.g. -c=student (请输入控制器名,例如 -c=student)`)
	}

	swaggerTags, _ := cmd.Flags().GetString(`swaggerTags`)
	if swaggerTags == `` {
		swaggerTags = controllerName
	}
	serviceName, _ := cmd.Flags().GetString(`service`)
	if serviceName == `` {
		serviceName = controllerName
	}
	modelName, _ := cmd.Flags().GetString(`model`)
	if modelName == `` {
		modelName = serviceName
	}

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
			`./internal` + `/migrates/` + format.Camel2Snake(modelName) + `.migrate.go`:            tplUri + `templates/v2/internal/apps/migrate.tmpl`,
			`./internal` + `/models/` + format.Camel2Snake(modelName) + `.model.go`:                tplUri + `templates/v2/internal/apps/model.tmpl`,
			`./internal` + `/models/` + format.Camel2Snake(modelName) + `.filter.go`:               tplUri + `templates/v2/internal/apps/filter.tmpl`,
			`./internal` + `/repos/` + format.Camel2Snake(modelName) + `.repo.go`:                  tplUri + `templates/v2/internal/apps/repo.tmpl`,
			`./internal` + `/services/` + format.Camel2Snake(serviceName) + `.service.go`:          tplUri + `templates/v2/internal/apps/service.tmpl`,
		},
	}
	c.tplAppData = buildAppTplData(moduleName, controllerName, controllerName, serviceName, modelName, swaggerTags)

	fields, err := resolveFields(cmd)
	if err != nil {
		return err
	}
	c.tplAppData[`Fields`] = fields
	c.tplAppData[`SearchColumn`] = firstStringColumn(fields)
	c.tplAppData[`HasTime`] = hasTimeField(fields)
	return nil
}
