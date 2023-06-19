package comer

import "github.com/spf13/cobra"

func (c *Comer) init(cmd *cobra.Command, args []string) {
	// fmt.Printf(`cmd=%v args=%v`, cmd, args)
	// fmt.Println(``)
	c.debugMode = true
	c.cmd = cmd
	c.args = args
	c.version = `v1.0`
	tmpProjectDir := `/../projects`
	c.Framework = &Framework{
		dirs: []string{
			`.` + tmpProjectDir + `/cmd`,
			`.` + tmpProjectDir + `/apps`,
			`.` + tmpProjectDir + `/configs`,
			`.` + tmpProjectDir + `/global`,
			`.` + tmpProjectDir + `/middlewares`,
			`.` + tmpProjectDir + `/middlewares/token`,
			`.` + tmpProjectDir + `/services`,
			`.` + tmpProjectDir + `/repos`,
			`.` + tmpProjectDir + `/models`,
			`.` + tmpProjectDir + `/utils`,
		},
		files: map[string]string{
			`.` + tmpProjectDir + `/README.md`: `./comer/template/v1/README.md`,
			`.` + tmpProjectDir + `/go.mod`:    `./comer/template/v1/go.mod.tpl`,
			`.` + tmpProjectDir + `/main.go`:   `./comer/template/v1/main.tpl`,

			`.` + tmpProjectDir + `/cmd/root.go`:    `./comer/template/v1/cmd/cmd.root.tpl`,
			`.` + tmpProjectDir + `/cmd/server.go`:  `./comer/template/v1/cmd/cmd.server.tpl`,
			`.` + tmpProjectDir + `/cmd/init.go`:    `./comer/template/v1/cmd/cmd.init.tpl`,
			`.` + tmpProjectDir + `/cmd/migrate.go`: `./comer/template/v1/cmd/cmd.migrate.tpl`,

			`.` + tmpProjectDir + `/configs/settings-local.yml`: `./comer/template/v1/config/config.settings-local.tpl`,
			`.` + tmpProjectDir + `/configs/casbin.conf`:        `./comer/template/v1/config/config.casbin.tpl`,

			`.` + tmpProjectDir + `/middlewares/CasbinMiddleware.go`:    `./comer/template/v1/middleware/middleware.Casbin.tpl`,
			`.` + tmpProjectDir + `/middlewares/CrosMiddleware.go`:      `./comer/template/v1/middleware/middleware.Cros.tpl`,
			`.` + tmpProjectDir + `/middlewares/JWTAuthMiddleware.go`:   `./comer/template/v1/middleware/middleware.JWTAuth.tpl`,
			`.` + tmpProjectDir + `/middlewares/token/jwttoken.go`:      `./comer/template/v1/middleware/middleware.JWTToken.tpl`,
			`.` + tmpProjectDir + `/middlewares/LoggerMiddleware.go`:    `./comer/template/v1/middleware/middleware.Logger.tpl`,
			`.` + tmpProjectDir + `/middlewares/RateLimitMiddleware.go`: `./comer/template/v1/middleware/middleware.RateLimit.tpl`,

			`.` + tmpProjectDir + `/global/global.go`: `./comer/template/v1/global/global.tpl`,
			`.` + tmpProjectDir + `/global/cache.go`:  `./comer/template/v1/global/global.cache.tpl`,
			`.` + tmpProjectDir + `/global/casbin.go`: `./comer/template/v1/global/global.casbin.tpl`,
			`.` + tmpProjectDir + `/global/config.go`: `./comer/template/v1/global/global.config.tpl`,
			`.` + tmpProjectDir + `/global/log.go`:    `./comer/template/v1/global/global.log.tpl`,
			`.` + tmpProjectDir + `/global/mysql.go`:  `./comer/template/v1/global/global.mysql.tpl`,
			`.` + tmpProjectDir + `/global/redis.go`:  `./comer/template/v1/global/global.redis.tpl`,
		},
	}
	c.tplData = map[string]any{
		`moduleName`: `github.com/imoowi/newProject`,
		`dbName`:     `com_imoowi_comer_newproject`,
	}
}
