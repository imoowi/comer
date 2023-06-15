package comer

import "github.com/spf13/cobra"

func (c *Comer) init(cmd *cobra.Command, args []string) {
	// fmt.Printf(`cmd=%v args=%v`, cmd, args)
	// fmt.Println(``)
	c.debugMode = true
	c.cmd = cmd
	c.args = args
	c.version = `v1.0`
	tmpProjectDir := `/projects`
	c.Framework = &Framework{
		dirs: []string{
			`.` + tmpProjectDir + `/cmd`,
			`.` + tmpProjectDir + `/app`,
			`.` + tmpProjectDir + `/config`,
			`.` + tmpProjectDir + `/global`,
			`.` + tmpProjectDir + `/middleware`,
			`.` + tmpProjectDir + `/service`,
			`.` + tmpProjectDir + `/repo`,
			`.` + tmpProjectDir + `/model`,
			`.` + tmpProjectDir + `/util`,
		},
		files: map[string]string{
			`.` + tmpProjectDir + `/README.md`:                 `./comer/template/v1/README.md`,
			`.` + tmpProjectDir + `/go.mod`:                    `./comer/template/v1/go.mod.tpl`,
			`.` + tmpProjectDir + `/main.go`:                   `./comer/template/v1/main.tpl`,
			`.` + tmpProjectDir + `/cmd/root.go`:               `./comer/template/v1/cmd.root.tpl`,
			`.` + tmpProjectDir + `/cmd/server.go`:             `./comer/template/v1/cmd.server.tpl`,
			`.` + tmpProjectDir + `/cmd/init.go`:               `./comer/template/v1/cmd.init.tpl`,
			`.` + tmpProjectDir + `/cmd/migrate.go`:            `./comer/template/v1/cmd.migrate.tpl`,
			`.` + tmpProjectDir + `/config/settings-local.yml`: `./comer/template/v1/config.settings-local.tpl`,
			`.` + tmpProjectDir + `/config/casbin.conf`:        `./comer/template/v1/config.casbin.tpl`,

			`.` + tmpProjectDir + `/middleware/CasbinMiddleware.go`: `./comer/template/v1/middleware.Casbin.tpl`,
			// `.` + tmpProjectDir + `/middleware/CrosMiddleware.go`:      `./comer/template/v1/middleware.Cros.tpl`,
			// `.` + tmpProjectDir + `/middleware/JWTAuthMiddleware.go`:   `./comer/template/v1/middleware.JWTAuth.tpl`,
			// `.` + tmpProjectDir + `/middleware/LoggerMiddleware.go`:    `./comer/template/v1/middleware.Logger.tpl`,
			// `.` + tmpProjectDir + `/middleware/RateLimitMiddleware.go`: `./comer/template/v1/middleware.RateLimit.tpl`,
		},
	}
	c.tplData = map[string]any{
		`moduleName`: `com.imoowi.comer.newProject`,
		`dbName`:     `com_imoowi_comer_newproject`,
	}
}
