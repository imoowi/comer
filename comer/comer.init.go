package comer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	tmpPath, err := cmd.Flags().GetString(`path`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if tmpPath == `` {
		tmpPath = `.`
	}
	gomod := tmpPath + `/go.mod`
	_, err = os.Stat(gomod)
	if os.IsNotExist(err) {
		moduleName, err := cmd.Flags().GetString(`module`)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		if moduleName == `` {
			fmt.Println(`pls input module, e.g. --module=test007 (请输入go.mod文件的module,例如 --module=test007)`)
			return false
		}
		c.moduleName = moduleName
	} else {
		data, err := ioutil.ReadFile(gomod)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		lines := strings.Split(string(data), "\n")
		c.moduleName = strings.Replace(lines[0], "module ", "", -1)
		if c.moduleName == `` {
			log.Println(`没有在当前目录找到go.mod里的module配置`)
			return false
		}

	}

	tplUri := `./comer/`
	tplUri = ``
	// return false
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
			tmpPath + `/README.md`: tplUri + `templates/v1/README.md`,
			tmpPath + `/go.mod`:    tplUri + `templates/v1/go.mod.tpl`,
			tmpPath + `/main.go`:   tplUri + `templates/v1/main.tpl`,
			tmpPath + `/.air.toml`: tplUri + `templates/v1/air.tpl`,

			tmpPath + `/.vscode/launch.json`:   tplUri + `templates/v1/.vscode/launch.json`,
			tmpPath + `/.vscode/settings.json`: tplUri + `templates/v1/.vscode/settings.json`,

			tmpPath + `/cmd/root.go`:    tplUri + `templates/v1/cmd/cmd.root.tpl`,
			tmpPath + `/cmd/server.go`:  tplUri + `templates/v1/cmd/cmd.server.tpl`,
			tmpPath + `/cmd/init.go`:    tplUri + `templates/v1/cmd/cmd.init.tpl`,
			tmpPath + `/cmd/migrate.go`: tplUri + `templates/v1/cmd/cmd.migrate.tpl`,

			tmpPath + `/components/mysql.go`:   tplUri + `templates/v1/components/components.mysql.tpl`,
			tmpPath + `/components/redis.go`:   tplUri + `templates/v1/components/components.redis.tpl`,
			tmpPath + `/components/captcha.go`: tplUri + `templates/v1/components/components.captcha.tpl`,

			tmpPath + `/configs/settings-local.yml`: tplUri + `templates/v1/config/config.settings-local.tpl`,
			tmpPath + `/configs/casbin.conf`:        tplUri + `templates/v1/config/config.casbin.tpl`,

			tmpPath + `/global/global.go`: tplUri + `templates/v1/global/global.tpl`,
			tmpPath + `/global/cache.go`:  tplUri + `templates/v1/global/global.cache.tpl`,
			tmpPath + `/global/casbin.go`: tplUri + `templates/v1/global/global.casbin.tpl`,
			tmpPath + `/global/config.go`: tplUri + `templates/v1/global/global.config.tpl`,
			tmpPath + `/global/log.go`:    tplUri + `templates/v1/global/global.log.tpl`,
			tmpPath + `/global/mysql.go`:  tplUri + `templates/v1/global/global.mysql.tpl`,
			tmpPath + `/global/redis.go`:  tplUri + `templates/v1/global/global.redis.tpl`,

			tmpPath + `/middlewares/CasbinMiddleware.go`:    tplUri + `templates/v1/middleware/middleware.Casbin.tpl`,
			tmpPath + `/middlewares/CrosMiddleware.go`:      tplUri + `templates/v1/middleware/middleware.Cros.tpl`,
			tmpPath + `/middlewares/JWTAuthMiddleware.go`:   tplUri + `templates/v1/middleware/middleware.JWTAuth.tpl`,
			tmpPath + `/middlewares/token/jwttoken.go`:      tplUri + `templates/v1/middleware/token/middleware.JWTToken.tpl`,
			tmpPath + `/middlewares/LoggerMiddleware.go`:    tplUri + `templates/v1/middleware/middleware.Logger.tpl`,
			tmpPath + `/middlewares/RateLimitMiddleware.go`: tplUri + `templates/v1/middleware/middleware.RateLimit.tpl`,
			tmpPath + `/middlewares/VcodeMiddleware.go`:     tplUri + `templates/v1/middleware/middleware.Vcode.tpl`,
			tmpPath + `/middlewares/middleware.go`:          tplUri + `templates/v1/middleware/middleware.tpl`,

			tmpPath + `/router/router.go`: tplUri + `templates/v1/router/router.tpl`,

			tmpPath + `/utils/copy/copy.go`:         tplUri + `templates/v1/utils/copy/copy.tpl`,
			tmpPath + `/utils/format/format.go`:     tplUri + `templates/v1/utils/format/format.tpl`,
			tmpPath + `/utils/maker/maker.go`:       tplUri + `templates/v1/utils/maker/maker.tpl`,
			tmpPath + `/utils/myfile/myfile.go`:     tplUri + `templates/v1/utils/myfile/myfile.tpl`,
			tmpPath + `/utils/mytime/mytime.go`:     tplUri + `templates/v1/utils/mytime/mytime.tpl`,
			tmpPath + `/utils/mytime/translater.go`: tplUri + `templates/v1/utils/mytime/translater.tpl`,
			tmpPath + `/utils/mytime/week.go`:       tplUri + `templates/v1/utils/mytime/week.tpl`,
			tmpPath + `/utils/office/excel.go`:      tplUri + `templates/v1/utils/office/excel.tpl`,
			tmpPath + `/utils/password/password.go`: tplUri + `templates/v1/utils/password/password.tpl`,
			tmpPath + `/utils/request/http.go`:      tplUri + `templates/v1/utils/request/http.tpl`,
			tmpPath + `/utils/request/pages.go`:     tplUri + `templates/v1/utils/request/pages.tpl`,
			tmpPath + `/utils/response/pages.go`:    tplUri + `templates/v1/utils/response/pages.tpl`,
			tmpPath + `/utils/response/response.go`: tplUri + `templates/v1/utils/response/response.tpl`,
			tmpPath + `/utils/slice/slice.go`:       tplUri + `templates/v1/utils/slice/slice.tpl`,
			tmpPath + `/utils/logger.go`:            tplUri + `templates/v1/utils/logger.tpl`,
			tmpPath + `/utils/utils.go`:             tplUri + `templates/v1/utils/utils.tpl`,

			tmpPath + `/apps/apps.go`:           tplUri + `templates/v1/apps/apps.tpl`,
			tmpPath + `/apps/swagger/router.go`: tplUri + `templates/v1/apps/swagger/router.tpl`,

			tmpPath + `/apps/common/router.go`:                   tplUri + `templates/v1/apps/common/router.tpl`,
			tmpPath + `/apps/common/handlers/captcha.handler.go`: tplUri + `templates/v1/apps/common/handlers/captcha.handler.tpl`,

			tmpPath + `/apps/user/router.go`:                   tplUri + `templates/v1/apps/user/router.tpl`,
			tmpPath + `/apps/user/handlers/auth.handler.go`:    tplUri + `templates/v1/apps/user/handlers/auth.handler.tpl`,
			tmpPath + `/apps/user/migrate/role.migrate.go`:     tplUri + `templates/v1/apps/user/migrate/role.migrate.tpl`,
			tmpPath + `/apps/user/migrate/user.migrate.go`:     tplUri + `templates/v1/apps/user/migrate/user.migrate.tpl`,
			tmpPath + `/apps/user/migrate/userlog.migrate.go`:  tplUri + `templates/v1/apps/user/migrate/userlog.migrate.tpl`,
			tmpPath + `/apps/user/migrate/userrole.migrate.go`: tplUri + `templates/v1/apps/user/migrate/userrole.migrate.tpl`,

			tmpPath + `/apps/user/models/role.model.go`:     tplUri + `templates/v1/apps/user/models/role.model.tpl`,
			tmpPath + `/apps/user/models/user.model.go`:     tplUri + `templates/v1/apps/user/models/user.model.tpl`,
			tmpPath + `/apps/user/models/userlog.model.go`:  tplUri + `templates/v1/apps/user/models/userlog.model.tpl`,
			tmpPath + `/apps/user/models/userrole.model.go`: tplUri + `templates/v1/apps/user/models/userrole.model.tpl`,

			tmpPath + `/apps/user/repos/role.repo.go`:     tplUri + `templates/v1/apps/user/repos/role.repo.tpl`,
			tmpPath + `/apps/user/repos/user.repo.go`:     tplUri + `templates/v1/apps/user/repos/user.repo.tpl`,
			tmpPath + `/apps/user/repos/userlog.repo.go`:  tplUri + `templates/v1/apps/user/repos/userlog.repo.tpl`,
			tmpPath + `/apps/user/repos/userrole.repo.go`: tplUri + `templates/v1/apps/user/repos/userrole.repo.tpl`,

			tmpPath + `/apps/user/services/role.service.go`:     tplUri + `templates/v1/apps/user/services/role.service.tpl`,
			tmpPath + `/apps/user/services/user.service.go`:     tplUri + `templates/v1/apps/user/services/user.service.tpl`,
			tmpPath + `/apps/user/services/userlog.service.go`:  tplUri + `templates/v1/apps/user/services/userlog.service.tpl`,
			tmpPath + `/apps/user/services/userrole.service.go`: tplUri + `templates/v1/apps/user/services/userrole.service.tpl`,
		},
	}
	c.tplData = map[string]any{
		`moduleName`: c.moduleName,
		`dbName`:     fmt.Sprintf(`comer_project_db_%s`, maker.MakeRandStr(6)),
	}
	fmt.Println(c.moduleName)
	// return false
	return true
}
