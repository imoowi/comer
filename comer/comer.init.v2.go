/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Comer) initV2(cmd *cobra.Command, args []string) bool {
	c.debugMode = true
	c.cmd = cmd
	c.args = args
	moduleName := ``
	if len(args) > 0 {
		moduleName = args[0]
	}
	if moduleName == `` {
		fmt.Println(`请输入module,例如: comer new github.com/imoowi/comer-example (pls input module, e.g. comer new github.com/imoowi/comer-example)`)
		return false
	}
	c.moduleName = moduleName
	moduleNameArr := strings.Split(c.moduleName, `/`)
	c.moduleProjectName = moduleNameArr[len(moduleNameArr)-1]

	c.path = c.moduleName
	tplUri := ``
	c.Framework = &Framework{
		dirs: []string{
			c.path + `/.vscode`,
			c.path + `/cmd`,
			c.path + `/configs`,
			c.path + `/docs`,
			c.path + `/runtime`,
			c.path + `/test`,
			c.path + `/internal/controllers`,
			c.path + `/internal/global`,
			c.path + `/internal/middlewares`,
			c.path + `/internal/middlewares/token`,
			c.path + `/internal/migrates`,
			c.path + `/internal/models`,
			c.path + `/internal/repos`,
			c.path + `/internal/router`,
			c.path + `/internal/services`,
		},
		files: map[string]string{
			c.path + `/README.md`:                 tplUri + `templates/v2/README.md.tmpl`,
			c.path + `/Makefile`:                  tplUri + `templates/v2/Makefile.tmpl`,
			c.path + `/Dockerfile`:                tplUri + `templates/v2/Dockerfile.tmpl`,
			c.path + `/docker-compose.yml`:        tplUri + `templates/v2/docker-compose.yml.tmpl`,
			c.path + `/start_server_in_docker.sh`: tplUri + `templates/v2/start_server_in_docker.sh.tmpl`,
			c.path + `/go.mod`:                    tplUri + `templates/v2/go.mod.tmpl`,
			c.path + `/main.go`:                   tplUri + `templates/v2/main.go.tmpl`,
			c.path + `/.air.toml`:                 tplUri + `templates/v2/air.toml.tmpl`,
			c.path + `/.gitignore`:                tplUri + `templates/v2/gitignore.tmpl`,
			c.path + `/docs/init.go`:              tplUri + `templates/v2/docs/init.go.tmpl`,

			//.vscode√
			c.path + `/.vscode/launch.json`:   tplUri + `templates/v2/vscode/launch.json.tmpl`,
			c.path + `/.vscode/settings.json`: tplUri + `templates/v2/vscode/settings.json.tmpl`,

			//cmd√
			c.path + `/cmd/root.go`:    tplUri + `templates/v2/cmd/root.go.tmpl`,
			c.path + `/cmd/server.go`:  tplUri + `templates/v2/cmd/server.go.tmpl`,
			c.path + `/cmd/init.go`:    tplUri + `templates/v2/cmd/init.go.tmpl`,
			c.path + `/cmd/migrate.go`: tplUri + `templates/v2/cmd/migrate.go.tmpl`,

			//configs√
			c.path + `/configs/settings-local.yml`: tplUri + `templates/v2/configs/settings-local.yml.tmpl`,
			c.path + `/configs/casbin.conf`:        tplUri + `templates/v2/configs/casbin.conf.tmpl`,

			//controllers√
			c.path + `/internal/controllers/auth.controller.go`:    tplUri + `templates/v2/internal/controllers/auth.controller.go.tmpl`,
			c.path + `/internal/controllers/captcha.controller.go`: tplUri + `templates/v2/internal/controllers/captcha.controller.go.tmpl`,
			c.path + `/internal/controllers/user.controller.go`:    tplUri + `templates/v2/internal/controllers/user.controller.go.tmpl`,

			//global√
			c.path + `/internal/global/global.go`:         tplUri + `templates/v2/internal/global/global.go.tmpl`,
			c.path + `/internal/global/global.userlog.go`: tplUri + `templates/v2/internal/global/global.userlog.go.tmpl`,
			c.path + `/internal/global/cache.go`:          tplUri + `templates/v2/internal/global/cache.go.tmpl`,
			c.path + `/internal/global/casbin.go`:         tplUri + `templates/v2/internal/global/casbin.go.tmpl`,
			c.path + `/internal/global/config.go`:         tplUri + `templates/v2/internal/global/config.go.tmpl`,
			c.path + `/internal/global/log.go`:            tplUri + `templates/v2/internal/global/log.go.tmpl`,
			c.path + `/internal/global/mysql.go`:          tplUri + `templates/v2/internal/global/mysql.go.tmpl`,
			c.path + `/internal/global/redis.go`:          tplUri + `templates/v2/internal/global/redis.go.tmpl`,

			//middlewares√
			c.path + `/internal/middlewares/CasbinMiddleware.go`:    tplUri + `templates/v2/internal/middlewares/CasbinMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/CrosMiddleware.go`:      tplUri + `templates/v2/internal/middlewares/CrosMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/JWTAuthMiddleware.go`:   tplUri + `templates/v2/internal/middlewares/JWTAuthMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/token/jwttoken.go`:      tplUri + `templates/v2/internal/middlewares/token/jwttoken.go.tmpl`,
			c.path + `/internal/middlewares/LoggerMiddleware.go`:    tplUri + `templates/v2/internal/middlewares/LoggerMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/RateLimitMiddleware.go`: tplUri + `templates/v2/internal/middlewares/RateLimitMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/VcodeMiddleware.go`:     tplUri + `templates/v2/internal/middlewares/VcodeMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/UserlogMiddleware.go`:   tplUri + `templates/v2/internal/middlewares/UserlogMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/RequestIdMiddleware.go`: tplUri + `templates/v2/internal/middlewares/RequestIdMiddleware.go.tmpl`,
			c.path + `/internal/middlewares/middleware.go`:          tplUri + `templates/v2/internal/middlewares/middleware.go.tmpl`,

			//migrates√
			c.path + `/internal/migrates/init.go`:              tplUri + `templates/v2/internal/migrates/init.go.tmpl`,
			c.path + `/internal/migrates/role.migrate.go`:      tplUri + `templates/v2/internal/migrates/role.migrate.go.tmpl`,
			c.path + `/internal/migrates/user.migrate.go`:      tplUri + `templates/v2/internal/migrates/user.migrate.go.tmpl`,
			c.path + `/internal/migrates/user_log.migrate.go`:  tplUri + `templates/v2/internal/migrates/user_log.migrate.go.tmpl`,
			c.path + `/internal/migrates/user_role.migrate.go`: tplUri + `templates/v2/internal/migrates/user_role.migrate.go.tmpl`,

			//models√
			c.path + `/internal/models/role.filter.go`:      tplUri + `templates/v2/internal/models/role.filter.go.tmpl`,
			c.path + `/internal/models/role.model.go`:       tplUri + `templates/v2/internal/models/role.model.go.tmpl`,
			c.path + `/internal/models/user.filter.go`:      tplUri + `templates/v2/internal/models/user.filter.go.tmpl`,
			c.path + `/internal/models/user.model.go`:       tplUri + `templates/v2/internal/models/user.model.go.tmpl`,
			c.path + `/internal/models/user_log.filter.go`:  tplUri + `templates/v2/internal/models/user_log.filter.go.tmpl`,
			c.path + `/internal/models/user_log.model.go`:   tplUri + `templates/v2/internal/models/user_log.model.go.tmpl`,
			c.path + `/internal/models/user_role.filter.go`: tplUri + `templates/v2/internal/models/user_role.filter.go.tmpl`,
			c.path + `/internal/models/user_role.model.go`:  tplUri + `templates/v2/internal/models/user_role.model.go.tmpl`,

			//repos√
			c.path + `/internal/repos/init.go`:           tplUri + `templates/v2/internal/repos/init.go.tmpl`,
			c.path + `/internal/repos/role.repo.go`:      tplUri + `templates/v2/internal/repos/role.repo.go.tmpl`,
			c.path + `/internal/repos/user.repo.go`:      tplUri + `templates/v2/internal/repos/user.repo.go.tmpl`,
			c.path + `/internal/repos/user_log.repo.go`:  tplUri + `templates/v2/internal/repos/user_log.repo.go.tmpl`,
			c.path + `/internal/repos/user_role.repo.go`: tplUri + `templates/v2/internal/repos/user_role.repo.go.tmpl`,

			//router√
			c.path + `/internal/router/init.go`:           tplUri + `templates/v2/internal/router/init.go.tmpl`,
			c.path + `/internal/router/auth.router.go`:    tplUri + `templates/v2/internal/router/auth.router.go.tmpl`,
			c.path + `/internal/router/common.router.go`:  tplUri + `templates/v2/internal/router/common.router.go.tmpl`,
			c.path + `/internal/router/swagger.router.go`: tplUri + `templates/v2/internal/router/swagger.router.go.tmpl`,

			//services√
			c.path + `/internal/services/init.go`:              tplUri + `templates/v2/internal/services/init.go.tmpl`,
			c.path + `/internal/services/role.service.go`:      tplUri + `templates/v2/internal/services/role.service.go.tmpl`,
			c.path + `/internal/services/user.service.go`:      tplUri + `templates/v2/internal/services/user.service.go.tmpl`,
			c.path + `/internal/services/user_log.service.go`:  tplUri + `templates/v2/internal/services/user_log.service.go.tmpl`,
			c.path + `/internal/services/user_role.service.go`: tplUri + `templates/v2/internal/services/user_role.service.go.tmpl`,

			//test
			c.path + `/test/login.go`: tplUri + `templates/v2/test/login.go.tmpl`,
		},
	}
	moduleNameSlic := strings.Split(c.moduleName, `/`)
	exeName := moduleNameSlic[len(moduleNameSlic)-1]
	c.tplData = map[string]any{
		`moduleName`:        c.moduleName,
		`dbName`:            `comer_project`,
		`exeName`:           exeName,
		`moduleProjectName`: c.moduleProjectName,
	}
	return true
}
