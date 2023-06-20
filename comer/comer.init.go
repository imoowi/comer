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
	// tmpProjectDir = ``
	c.Framework = &Framework{
		dirs: []string{
			`.` + tmpProjectDir + `/apps`,
			`.` + tmpProjectDir + `/cmd`,
			`.` + tmpProjectDir + `/components`,
			`.` + tmpProjectDir + `/apps/common`,
			`.` + tmpProjectDir + `/apps/common/handlers`,
			`.` + tmpProjectDir + `/apps/swagger`,
			`.` + tmpProjectDir + `/apps/user/handlers`,
			`.` + tmpProjectDir + `/apps/user/migrate`,
			`.` + tmpProjectDir + `/apps/user/models`,
			`.` + tmpProjectDir + `/apps/user/repos`,
			`.` + tmpProjectDir + `/apps/user/services`,
			`.` + tmpProjectDir + `/configs`,
			`.` + tmpProjectDir + `/global`,
			`.` + tmpProjectDir + `/middlewares`,
			`.` + tmpProjectDir + `/middlewares/token`,
			`.` + tmpProjectDir + `/router`,
			`.` + tmpProjectDir + `/runtime`,
			`.` + tmpProjectDir + `/utils`,
			`.` + tmpProjectDir + `/utils/copy`,
			`.` + tmpProjectDir + `/utils/format`,
			`.` + tmpProjectDir + `/utils/maker`,
			`.` + tmpProjectDir + `/utils/myfile`,
			`.` + tmpProjectDir + `/utils/mytime`,
			`.` + tmpProjectDir + `/utils/office`,
			`.` + tmpProjectDir + `/utils/password`,
			`.` + tmpProjectDir + `/utils/request`,
			`.` + tmpProjectDir + `/utils/response`,
			`.` + tmpProjectDir + `/utils/slice`,
			`.` + tmpProjectDir + `/.vscode`,
		},
		files: map[string]string{
			`.` + tmpProjectDir + `/README.md`: `./comer/template/v1/README.md`,
			`.` + tmpProjectDir + `/go.mod`:    `./comer/template/v1/go.mod.tpl`,
			`.` + tmpProjectDir + `/main.go`:   `./comer/template/v1/main.tpl`,
			`.` + tmpProjectDir + `/.air.toml`: `./comer/template/v1/air.tpl`,

			`.` + tmpProjectDir + `/.vscode/launch.json`:   `./comer/template/v1/.vscode/launch.json`,
			`.` + tmpProjectDir + `/.vscode/settings.json`: `./comer/template/v1/.vscode/settings.json`,

			`.` + tmpProjectDir + `/cmd/root.go`:    `./comer/template/v1/cmd/cmd.root.tpl`,
			`.` + tmpProjectDir + `/cmd/server.go`:  `./comer/template/v1/cmd/cmd.server.tpl`,
			`.` + tmpProjectDir + `/cmd/init.go`:    `./comer/template/v1/cmd/cmd.init.tpl`,
			`.` + tmpProjectDir + `/cmd/migrate.go`: `./comer/template/v1/cmd/cmd.migrate.tpl`,

			`.` + tmpProjectDir + `/components/mysql.go`:   `./comer/template/v1/components/components.mysql.tpl`,
			`.` + tmpProjectDir + `/components/redis.go`:   `./comer/template/v1/components/components.redis.tpl`,
			`.` + tmpProjectDir + `/components/captcha.go`: `./comer/template/v1/components/components.captcha.tpl`,

			`.` + tmpProjectDir + `/configs/settings-local.yml`: `./comer/template/v1/config/config.settings-local.tpl`,
			`.` + tmpProjectDir + `/configs/casbin.conf`:        `./comer/template/v1/config/config.casbin.tpl`,

			`.` + tmpProjectDir + `/global/global.go`: `./comer/template/v1/global/global.tpl`,
			`.` + tmpProjectDir + `/global/cache.go`:  `./comer/template/v1/global/global.cache.tpl`,
			`.` + tmpProjectDir + `/global/casbin.go`: `./comer/template/v1/global/global.casbin.tpl`,
			`.` + tmpProjectDir + `/global/config.go`: `./comer/template/v1/global/global.config.tpl`,
			`.` + tmpProjectDir + `/global/log.go`:    `./comer/template/v1/global/global.log.tpl`,
			`.` + tmpProjectDir + `/global/mysql.go`:  `./comer/template/v1/global/global.mysql.tpl`,
			`.` + tmpProjectDir + `/global/redis.go`:  `./comer/template/v1/global/global.redis.tpl`,

			`.` + tmpProjectDir + `/middlewares/CasbinMiddleware.go`:    `./comer/template/v1/middleware/middleware.Casbin.tpl`,
			`.` + tmpProjectDir + `/middlewares/CrosMiddleware.go`:      `./comer/template/v1/middleware/middleware.Cros.tpl`,
			`.` + tmpProjectDir + `/middlewares/JWTAuthMiddleware.go`:   `./comer/template/v1/middleware/middleware.JWTAuth.tpl`,
			`.` + tmpProjectDir + `/middlewares/token/jwttoken.go`:      `./comer/template/v1/middleware/token/middleware.JWTToken.tpl`,
			`.` + tmpProjectDir + `/middlewares/LoggerMiddleware.go`:    `./comer/template/v1/middleware/middleware.Logger.tpl`,
			`.` + tmpProjectDir + `/middlewares/RateLimitMiddleware.go`: `./comer/template/v1/middleware/middleware.RateLimit.tpl`,
			`.` + tmpProjectDir + `/middlewares/VcodeMiddleware.go`:     `./comer/template/v1/middleware/middleware.Vcode.tpl`,
			`.` + tmpProjectDir + `/middlewares/middleware.go`:          `./comer/template/v1/middleware/middleware.tpl`,

			`.` + tmpProjectDir + `/router/router.go`: `./comer/template/v1/router/router.tpl`,

			`.` + tmpProjectDir + `/utils/copy/copy.go`:         `./comer/template/v1/utils/copy/copy.tpl`,
			`.` + tmpProjectDir + `/utils/format/format.go`:     `./comer/template/v1/utils/format/format.tpl`,
			`.` + tmpProjectDir + `/utils/maker/maker.go`:       `./comer/template/v1/utils/maker/maker.tpl`,
			`.` + tmpProjectDir + `/utils/myfile/myfile.go`:     `./comer/template/v1/utils/myfile/myfile.tpl`,
			`.` + tmpProjectDir + `/utils/mytime/mytime.go`:     `./comer/template/v1/utils/mytime/mytime.tpl`,
			`.` + tmpProjectDir + `/utils/mytime/translater.go`: `./comer/template/v1/utils/mytime/translater.tpl`,
			`.` + tmpProjectDir + `/utils/mytime/week.go`:       `./comer/template/v1/utils/mytime/week.tpl`,
			`.` + tmpProjectDir + `/utils/office/excel.go`:      `./comer/template/v1/utils/office/excel.tpl`,
			`.` + tmpProjectDir + `/utils/password/password.go`: `./comer/template/v1/utils/password/password.tpl`,
			`.` + tmpProjectDir + `/utils/request/http.go`:      `./comer/template/v1/utils/request/http.tpl`,
			`.` + tmpProjectDir + `/utils/request/pages.go`:     `./comer/template/v1/utils/request/pages.tpl`,
			`.` + tmpProjectDir + `/utils/response/pages.go`:    `./comer/template/v1/utils/response/pages.tpl`,
			`.` + tmpProjectDir + `/utils/response/response.go`: `./comer/template/v1/utils/response/response.tpl`,
			`.` + tmpProjectDir + `/utils/slice/slice.go`:       `./comer/template/v1/utils/slice/slice.tpl`,
			`.` + tmpProjectDir + `/utils/logger.go`:            `./comer/template/v1/utils/logger.tpl`,
			`.` + tmpProjectDir + `/utils/utils.go`:             `./comer/template/v1/utils/utils.tpl`,

			`.` + tmpProjectDir + `/apps/apps.go`:           `./comer/template/v1/apps/apps.tpl`,
			`.` + tmpProjectDir + `/apps/swagger/router.go`: `./comer/template/v1/apps/swagger/router.tpl`,

			`.` + tmpProjectDir + `/apps/common/router.go`:                   `./comer/template/v1/apps/common/router.tpl`,
			`.` + tmpProjectDir + `/apps/common/handlers/captcha.handler.go`: `./comer/template/v1/apps/common/handlers/captcha.handler.tpl`,

			`.` + tmpProjectDir + `/apps/user/router.go`:                   `./comer/template/v1/apps/user/router.tpl`,
			`.` + tmpProjectDir + `/apps/user/handlers/auth.handler.go`:    `./comer/template/v1/apps/user/handlers/auth.handler.tpl`,
			`.` + tmpProjectDir + `/apps/user/migrate/role.migrate.go`:     `./comer/template/v1/apps/user/migrate/role.migrate.tpl`,
			`.` + tmpProjectDir + `/apps/user/migrate/user.migrate.go`:     `./comer/template/v1/apps/user/migrate/user.migrate.tpl`,
			`.` + tmpProjectDir + `/apps/user/migrate/userlog.migrate.go`:  `./comer/template/v1/apps/user/migrate/userlog.migrate.tpl`,
			`.` + tmpProjectDir + `/apps/user/migrate/userrole.migrate.go`: `./comer/template/v1/apps/user/migrate/userrole.migrate.tpl`,

			`.` + tmpProjectDir + `/apps/user/models/role.model.go`:     `./comer/template/v1/apps/user/models/role.model.tpl`,
			`.` + tmpProjectDir + `/apps/user/models/user.model.go`:     `./comer/template/v1/apps/user/models/user.model.tpl`,
			`.` + tmpProjectDir + `/apps/user/models/userlog.model.go`:  `./comer/template/v1/apps/user/models/userlog.model.tpl`,
			`.` + tmpProjectDir + `/apps/user/models/userrole.model.go`: `./comer/template/v1/apps/user/models/userrole.model.tpl`,

			`.` + tmpProjectDir + `/apps/user/repos/role.repo.go`:     `./comer/template/v1/apps/user/repos/role.repo.tpl`,
			`.` + tmpProjectDir + `/apps/user/repos/user.repo.go`:     `./comer/template/v1/apps/user/repos/user.repo.tpl`,
			`.` + tmpProjectDir + `/apps/user/repos/userlog.repo.go`:  `./comer/template/v1/apps/user/repos/userlog.repo.tpl`,
			`.` + tmpProjectDir + `/apps/user/repos/userrole.repo.go`: `./comer/template/v1/apps/user/repos/userrole.repo.tpl`,

			`.` + tmpProjectDir + `/apps/user/services/role.service.go`:     `./comer/template/v1/apps/user/services/role.service.tpl`,
			`.` + tmpProjectDir + `/apps/user/services/user.service.go`:     `./comer/template/v1/apps/user/services/user.service.tpl`,
			`.` + tmpProjectDir + `/apps/user/services/userlog.service.go`:  `./comer/template/v1/apps/user/services/userlog.service.tpl`,
			`.` + tmpProjectDir + `/apps/user/services/userrole.service.go`: `./comer/template/v1/apps/user/services/userrole.service.tpl`,
		},
	}
	c.tplData = map[string]any{
		`moduleName`: `github.com/imoowi/newProject`,
		`dbName`:     `com_imoowi_comer_newproject`,
	}
}
