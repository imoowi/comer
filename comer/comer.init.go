package comer

import (
	"fmt"

	"github.com/imoowi/comer/utils/maker"
	"github.com/spf13/cobra"
)

func (c *Comer) init(cmd *cobra.Command, args []string) bool {
	// fmt.Printf(`cmd=%v args=%v`, cmd, args)
	// return false
	// fmt.Println(``)
	c.debugMode = true
	c.cmd = cmd
	c.args = args
	moduleName, err := cmd.Flags().GetString(`moduleName`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	c.moduleName = moduleName
	// c.moduleName = `github.com/imoowi/newProject`
	// tmpPath := `../projects`
	tmpPath, err := cmd.Flags().GetString(`path`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if tmpPath == `` {
		tmpPath = `.`
	}
	c.Framework = &Framework{
		dirs: []string{
			tmpPath + `/apps`,
			tmpPath + `/cmd`,
			tmpPath + `/components`,
			tmpPath + `/apps/common`,
			tmpPath + `/apps/common/handlers`,
			tmpPath + `/apps/swagger`,
			tmpPath + `/apps/user/handlers`,
			tmpPath + `/apps/user/migrate`,
			tmpPath + `/apps/user/models`,
			tmpPath + `/apps/user/repos`,
			tmpPath + `/apps/user/services`,
			tmpPath + `/configs`,
			tmpPath + `/global`,
			tmpPath + `/middlewares`,
			tmpPath + `/middlewares/token`,
			tmpPath + `/router`,
			tmpPath + `/runtime`,
			tmpPath + `/utils`,
			tmpPath + `/utils/copy`,
			tmpPath + `/utils/format`,
			tmpPath + `/utils/maker`,
			tmpPath + `/utils/myfile`,
			tmpPath + `/utils/mytime`,
			tmpPath + `/utils/office`,
			tmpPath + `/utils/password`,
			tmpPath + `/utils/request`,
			tmpPath + `/utils/response`,
			tmpPath + `/utils/slice`,
			tmpPath + `/.vscode`,
		},
		files: map[string]string{
			tmpPath + `/README.md`: `./comer/template/v1/README.md`,
			tmpPath + `/go.mod`:    `./comer/template/v1/go.mod.tpl`,
			tmpPath + `/main.go`:   `./comer/template/v1/main.tpl`,
			tmpPath + `/.air.toml`: `./comer/template/v1/air.tpl`,

			tmpPath + `/.vscode/launch.json`:   `./comer/template/v1/.vscode/launch.json`,
			tmpPath + `/.vscode/settings.json`: `./comer/template/v1/.vscode/settings.json`,

			tmpPath + `/cmd/root.go`:    `./comer/template/v1/cmd/cmd.root.tpl`,
			tmpPath + `/cmd/server.go`:  `./comer/template/v1/cmd/cmd.server.tpl`,
			tmpPath + `/cmd/init.go`:    `./comer/template/v1/cmd/cmd.init.tpl`,
			tmpPath + `/cmd/migrate.go`: `./comer/template/v1/cmd/cmd.migrate.tpl`,

			tmpPath + `/components/mysql.go`:   `./comer/template/v1/components/components.mysql.tpl`,
			tmpPath + `/components/redis.go`:   `./comer/template/v1/components/components.redis.tpl`,
			tmpPath + `/components/captcha.go`: `./comer/template/v1/components/components.captcha.tpl`,

			tmpPath + `/configs/settings-local.yml`: `./comer/template/v1/config/config.settings-local.tpl`,
			tmpPath + `/configs/casbin.conf`:        `./comer/template/v1/config/config.casbin.tpl`,

			tmpPath + `/global/global.go`: `./comer/template/v1/global/global.tpl`,
			tmpPath + `/global/cache.go`:  `./comer/template/v1/global/global.cache.tpl`,
			tmpPath + `/global/casbin.go`: `./comer/template/v1/global/global.casbin.tpl`,
			tmpPath + `/global/config.go`: `./comer/template/v1/global/global.config.tpl`,
			tmpPath + `/global/log.go`:    `./comer/template/v1/global/global.log.tpl`,
			tmpPath + `/global/mysql.go`:  `./comer/template/v1/global/global.mysql.tpl`,
			tmpPath + `/global/redis.go`:  `./comer/template/v1/global/global.redis.tpl`,

			tmpPath + `/middlewares/CasbinMiddleware.go`:    `./comer/template/v1/middleware/middleware.Casbin.tpl`,
			tmpPath + `/middlewares/CrosMiddleware.go`:      `./comer/template/v1/middleware/middleware.Cros.tpl`,
			tmpPath + `/middlewares/JWTAuthMiddleware.go`:   `./comer/template/v1/middleware/middleware.JWTAuth.tpl`,
			tmpPath + `/middlewares/token/jwttoken.go`:      `./comer/template/v1/middleware/token/middleware.JWTToken.tpl`,
			tmpPath + `/middlewares/LoggerMiddleware.go`:    `./comer/template/v1/middleware/middleware.Logger.tpl`,
			tmpPath + `/middlewares/RateLimitMiddleware.go`: `./comer/template/v1/middleware/middleware.RateLimit.tpl`,
			tmpPath + `/middlewares/VcodeMiddleware.go`:     `./comer/template/v1/middleware/middleware.Vcode.tpl`,
			tmpPath + `/middlewares/middleware.go`:          `./comer/template/v1/middleware/middleware.tpl`,

			tmpPath + `/router/router.go`: `./comer/template/v1/router/router.tpl`,

			tmpPath + `/utils/copy/copy.go`:         `./comer/template/v1/utils/copy/copy.tpl`,
			tmpPath + `/utils/format/format.go`:     `./comer/template/v1/utils/format/format.tpl`,
			tmpPath + `/utils/maker/maker.go`:       `./comer/template/v1/utils/maker/maker.tpl`,
			tmpPath + `/utils/myfile/myfile.go`:     `./comer/template/v1/utils/myfile/myfile.tpl`,
			tmpPath + `/utils/mytime/mytime.go`:     `./comer/template/v1/utils/mytime/mytime.tpl`,
			tmpPath + `/utils/mytime/translater.go`: `./comer/template/v1/utils/mytime/translater.tpl`,
			tmpPath + `/utils/mytime/week.go`:       `./comer/template/v1/utils/mytime/week.tpl`,
			tmpPath + `/utils/office/excel.go`:      `./comer/template/v1/utils/office/excel.tpl`,
			tmpPath + `/utils/password/password.go`: `./comer/template/v1/utils/password/password.tpl`,
			tmpPath + `/utils/request/http.go`:      `./comer/template/v1/utils/request/http.tpl`,
			tmpPath + `/utils/request/pages.go`:     `./comer/template/v1/utils/request/pages.tpl`,
			tmpPath + `/utils/response/pages.go`:    `./comer/template/v1/utils/response/pages.tpl`,
			tmpPath + `/utils/response/response.go`: `./comer/template/v1/utils/response/response.tpl`,
			tmpPath + `/utils/slice/slice.go`:       `./comer/template/v1/utils/slice/slice.tpl`,
			tmpPath + `/utils/logger.go`:            `./comer/template/v1/utils/logger.tpl`,
			tmpPath + `/utils/utils.go`:             `./comer/template/v1/utils/utils.tpl`,

			tmpPath + `/apps/apps.go`:           `./comer/template/v1/apps/apps.tpl`,
			tmpPath + `/apps/swagger/router.go`: `./comer/template/v1/apps/swagger/router.tpl`,

			tmpPath + `/apps/common/router.go`:                   `./comer/template/v1/apps/common/router.tpl`,
			tmpPath + `/apps/common/handlers/captcha.handler.go`: `./comer/template/v1/apps/common/handlers/captcha.handler.tpl`,

			tmpPath + `/apps/user/router.go`:                   `./comer/template/v1/apps/user/router.tpl`,
			tmpPath + `/apps/user/handlers/auth.handler.go`:    `./comer/template/v1/apps/user/handlers/auth.handler.tpl`,
			tmpPath + `/apps/user/migrate/role.migrate.go`:     `./comer/template/v1/apps/user/migrate/role.migrate.tpl`,
			tmpPath + `/apps/user/migrate/user.migrate.go`:     `./comer/template/v1/apps/user/migrate/user.migrate.tpl`,
			tmpPath + `/apps/user/migrate/userlog.migrate.go`:  `./comer/template/v1/apps/user/migrate/userlog.migrate.tpl`,
			tmpPath + `/apps/user/migrate/userrole.migrate.go`: `./comer/template/v1/apps/user/migrate/userrole.migrate.tpl`,

			tmpPath + `/apps/user/models/role.model.go`:     `./comer/template/v1/apps/user/models/role.model.tpl`,
			tmpPath + `/apps/user/models/user.model.go`:     `./comer/template/v1/apps/user/models/user.model.tpl`,
			tmpPath + `/apps/user/models/userlog.model.go`:  `./comer/template/v1/apps/user/models/userlog.model.tpl`,
			tmpPath + `/apps/user/models/userrole.model.go`: `./comer/template/v1/apps/user/models/userrole.model.tpl`,

			tmpPath + `/apps/user/repos/role.repo.go`:     `./comer/template/v1/apps/user/repos/role.repo.tpl`,
			tmpPath + `/apps/user/repos/user.repo.go`:     `./comer/template/v1/apps/user/repos/user.repo.tpl`,
			tmpPath + `/apps/user/repos/userlog.repo.go`:  `./comer/template/v1/apps/user/repos/userlog.repo.tpl`,
			tmpPath + `/apps/user/repos/userrole.repo.go`: `./comer/template/v1/apps/user/repos/userrole.repo.tpl`,

			tmpPath + `/apps/user/services/role.service.go`:     `./comer/template/v1/apps/user/services/role.service.tpl`,
			tmpPath + `/apps/user/services/user.service.go`:     `./comer/template/v1/apps/user/services/user.service.tpl`,
			tmpPath + `/apps/user/services/userlog.service.go`:  `./comer/template/v1/apps/user/services/userlog.service.tpl`,
			tmpPath + `/apps/user/services/userrole.service.go`: `./comer/template/v1/apps/user/services/userrole.service.tpl`,
		},
	}
	c.tplData = map[string]any{
		`moduleName`: c.moduleName,
		`dbName`:     fmt.Sprintf(`comer_project_db_%s`, maker.MakeRandStr(6)),
	}
	return true
}
