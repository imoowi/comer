/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Comer) init(cmd *cobra.Command, args []string) bool {
	c.debugMode = true
	c.cmd = cmd
	c.args = args
	/*
		path, err := cmd.Flags().GetString(`path`)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		c.path = path
		_moduleName := ``
		moduleFile := `./go.mod`
		_, gErr := os.Stat(moduleFile)
		if os.IsNotExist(gErr) {
			log.Println(`go.mod not exists`)
			// return false
		} else {
			file, err := os.OpenFile(moduleFile, os.O_RDWR, 0544)
			if err != nil {
				fmt.Printf("File open failed! err: %v\n", err)
				return false
			}
			reader := bufio.NewReader(file)
			_moduleNameLine := ``
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
					_moduleNameLine = line
					break
				}
			}
			file.Close()
			if _moduleNameLine != `` {
				_moduleName = strings.Replace(_moduleNameLine, "module ", "", -1)
				_moduleName = strings.Replace(_moduleName, "\r", "", -1)
				_moduleName = strings.Replace(_moduleName, "\n", "", -1)
				_moduleName = strings.Replace(_moduleName, "\n\r", "", -1)
				_moduleName = strings.Replace(_moduleName, "\r\n", "", -1)
			}
		}
		if _moduleName == `` {
			moduleName, err := cmd.Flags().GetString(`module`)
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
			if _moduleName == `` && moduleName == `` {
				fmt.Println(`pls input module, e.g. -m=github.com/imoowi/comer-example (请输入go.mod文件的module,例如 -m=github.com/imoowi/comer-example)`)
				return false
			}
			c.moduleName = moduleName
		} else {
			c.moduleName = _moduleName
		}
		//*/
	moduleName, err := cmd.Flags().GetString(`module`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if moduleName == `` {
		// fmt.Println(`pls input module, e.g. -m=github.com/imoowi/comer-example (请输入go.mod文件的module,例如 -m=github.com/imoowi/comer-example)`)
		fmt.Println(`请输入module,例如: -m=github.com/imoowi/comer-example(pls input module, e.g. -m=github.com/imoowi/comer-example)`)
		return false
	}
	c.moduleName = moduleName
	/*
		if c.path == `` {
			c.path = c.moduleName
		}
		//*/
	c.path = c.moduleName
	tplUri := ``
	c.Framework = &Framework{
		dirs: []string{
			c.path + `/apps`,
			c.path + `/cmd`,
			c.path + `/components`,
			c.path + `/apps/common`,
			c.path + `/apps/common/handlers`,
			c.path + `/apps/swagger`,
			c.path + `/apps/user/handlers`,
			c.path + `/apps/user/migrates`,
			c.path + `/apps/user/models`,
			c.path + `/apps/user/repos`,
			c.path + `/apps/user/services`,
			c.path + `/configs`,
			c.path + `/global`,
			c.path + `/middlewares`,
			c.path + `/middlewares/token`,
			c.path + `/router`,
			c.path + `/runtime`,
			c.path + `/utils`,
			c.path + `/utils/copy`,
			c.path + `/utils/format`,
			c.path + `/utils/maker`,
			c.path + `/utils/myfile`,
			c.path + `/utils/mytime`,
			c.path + `/utils/office`,
			c.path + `/utils/password`,
			c.path + `/utils/request`,
			c.path + `/utils/response`,
			c.path + `/utils/slice`,
			c.path + `/.vscode`,
		},
		files: map[string]string{
			c.path + `/README.md`:                               tplUri + `templates/v1/README.md.tpl`,
			c.path + `/go.mod`:                                  tplUri + `templates/v1/go.mod.tpl`,
			c.path + `/main.go`:                                 tplUri + `templates/v1/main.tpl`,
			c.path + `/.air.toml`:                               tplUri + `templates/v1/air.tpl`,
			c.path + `/.vscode/launch.json`:                     tplUri + `templates/v1/vscode/launch.json`,
			c.path + `/.vscode/settings.json`:                   tplUri + `templates/v1/vscode/settings.json`,
			c.path + `/cmd/root.go`:                             tplUri + `templates/v1/cmd/cmd.root.tpl`,
			c.path + `/cmd/server.go`:                           tplUri + `templates/v1/cmd/cmd.server.tpl`,
			c.path + `/cmd/init.go`:                             tplUri + `templates/v1/cmd/cmd.init.tpl`,
			c.path + `/cmd/migrate.go`:                          tplUri + `templates/v1/cmd/cmd.migrate.tpl`,
			c.path + `/components/mysql.go`:                     tplUri + `templates/v1/components/components.mysql.tpl`,
			c.path + `/components/redis.go`:                     tplUri + `templates/v1/components/components.redis.tpl`,
			c.path + `/components/captcha.go`:                   tplUri + `templates/v1/components/components.captcha.tpl`,
			c.path + `/configs/settings-local.yml`:              tplUri + `templates/v1/config/config.settings-local.tpl`,
			c.path + `/configs/casbin.conf`:                     tplUri + `templates/v1/config/config.casbin.tpl`,
			c.path + `/global/global.go`:                        tplUri + `templates/v1/global/global.tpl`,
			c.path + `/global/global.userlog.go`:                tplUri + `templates/v1/global/global.userlog.tpl`,
			c.path + `/global/cache.go`:                         tplUri + `templates/v1/global/global.cache.tpl`,
			c.path + `/global/casbin.go`:                        tplUri + `templates/v1/global/global.casbin.tpl`,
			c.path + `/global/config.go`:                        tplUri + `templates/v1/global/global.config.tpl`,
			c.path + `/global/log.go`:                           tplUri + `templates/v1/global/global.log.tpl`,
			c.path + `/global/mysql.go`:                         tplUri + `templates/v1/global/global.mysql.tpl`,
			c.path + `/global/redis.go`:                         tplUri + `templates/v1/global/global.redis.tpl`,
			c.path + `/middlewares/CasbinMiddleware.go`:         tplUri + `templates/v1/middleware/middleware.Casbin.tpl`,
			c.path + `/middlewares/CrosMiddleware.go`:           tplUri + `templates/v1/middleware/middleware.Cros.tpl`,
			c.path + `/middlewares/JWTAuthMiddleware.go`:        tplUri + `templates/v1/middleware/middleware.JWTAuth.tpl`,
			c.path + `/middlewares/token/jwttoken.go`:           tplUri + `templates/v1/middleware/token/middleware.JWTToken.tpl`,
			c.path + `/middlewares/LoggerMiddleware.go`:         tplUri + `templates/v1/middleware/middleware.Logger.tpl`,
			c.path + `/middlewares/RateLimitMiddleware.go`:      tplUri + `templates/v1/middleware/middleware.RateLimit.tpl`,
			c.path + `/middlewares/VcodeMiddleware.go`:          tplUri + `templates/v1/middleware/middleware.Vcode.tpl`,
			c.path + `/middlewares/UserlogMiddleware.go`:        tplUri + `templates/v1/middleware/middleware.Userlog.tpl`,
			c.path + `/middlewares/middleware.go`:               tplUri + `templates/v1/middleware/middleware.tpl`,
			c.path + `/router/router.go`:                        tplUri + `templates/v1/router/router.tpl`,
			c.path + `/utils/copy/copy.go`:                      tplUri + `templates/v1/utils/copy/copy.tpl`,
			c.path + `/utils/format/format.go`:                  tplUri + `templates/v1/utils/format/format.tpl`,
			c.path + `/utils/maker/maker.go`:                    tplUri + `templates/v1/utils/maker/maker.tpl`,
			c.path + `/utils/myfile/myfile.go`:                  tplUri + `templates/v1/utils/myfile/myfile.tpl`,
			c.path + `/utils/mytime/mytime.go`:                  tplUri + `templates/v1/utils/mytime/mytime.tpl`,
			c.path + `/utils/mytime/translater.go`:              tplUri + `templates/v1/utils/mytime/translater.tpl`,
			c.path + `/utils/mytime/week.go`:                    tplUri + `templates/v1/utils/mytime/week.tpl`,
			c.path + `/utils/office/excel.go`:                   tplUri + `templates/v1/utils/office/excel.tpl`,
			c.path + `/utils/password/password.go`:              tplUri + `templates/v1/utils/password/password.tpl`,
			c.path + `/utils/request/http.go`:                   tplUri + `templates/v1/utils/request/http.tpl`,
			c.path + `/utils/request/pages.go`:                  tplUri + `templates/v1/utils/request/pages.tpl`,
			c.path + `/utils/response/pages.go`:                 tplUri + `templates/v1/utils/response/pages.tpl`,
			c.path + `/utils/response/response.go`:              tplUri + `templates/v1/utils/response/response.tpl`,
			c.path + `/utils/slice/slice.go`:                    tplUri + `templates/v1/utils/slice/slice.tpl`,
			c.path + `/utils/logger.go`:                         tplUri + `templates/v1/utils/logger.tpl`,
			c.path + `/utils/utils.go`:                          tplUri + `templates/v1/utils/utils.tpl`,
			c.path + `/apps/apps.go`:                            tplUri + `templates/v1/apps/apps.tpl`,
			c.path + `/apps/swagger/router.go`:                  tplUri + `templates/v1/apps/swagger/router.tpl`,
			c.path + `/apps/common/router.go`:                   tplUri + `templates/v1/apps/common/router.tpl`,
			c.path + `/apps/common/handlers/captcha.handler.go`: tplUri + `templates/v1/apps/common/handlers/captcha.handler.tpl`,
			c.path + `/apps/user/router.go`:                     tplUri + `templates/v1/apps/user/router.tpl`,
			c.path + `/apps/user/handlers/auth.handler.go`:      tplUri + `templates/v1/apps/user/handlers/auth.handler.tpl`,
			c.path + `/apps/user/migrates/role.migrate.go`:      tplUri + `templates/v1/apps/user/migrates/role.migrate.tpl`,
			c.path + `/apps/user/migrates/user.migrate.go`:      tplUri + `templates/v1/apps/user/migrates/user.migrate.tpl`,
			c.path + `/apps/user/migrates/userlog.migrate.go`:   tplUri + `templates/v1/apps/user/migrates/userlog.migrate.tpl`,
			c.path + `/apps/user/migrates/userrole.migrate.go`:  tplUri + `templates/v1/apps/user/migrates/userrole.migrate.tpl`,
			c.path + `/apps/user/models/role.model.go`:          tplUri + `templates/v1/apps/user/models/role.model.tpl`,
			c.path + `/apps/user/models/user.model.go`:          tplUri + `templates/v1/apps/user/models/user.model.tpl`,
			c.path + `/apps/user/models/userlog.model.go`:       tplUri + `templates/v1/apps/user/models/userlog.model.tpl`,
			c.path + `/apps/user/models/userrole.model.go`:      tplUri + `templates/v1/apps/user/models/userrole.model.tpl`,
			c.path + `/apps/user/repos/role.repo.go`:            tplUri + `templates/v1/apps/user/repos/role.repo.tpl`,
			c.path + `/apps/user/repos/user.repo.go`:            tplUri + `templates/v1/apps/user/repos/user.repo.tpl`,
			c.path + `/apps/user/repos/userlog.repo.go`:         tplUri + `templates/v1/apps/user/repos/userlog.repo.tpl`,
			c.path + `/apps/user/repos/userrole.repo.go`:        tplUri + `templates/v1/apps/user/repos/userrole.repo.tpl`,
			c.path + `/apps/user/services/role.service.go`:      tplUri + `templates/v1/apps/user/services/role.service.tpl`,
			c.path + `/apps/user/services/user.service.go`:      tplUri + `templates/v1/apps/user/services/user.service.tpl`,
			c.path + `/apps/user/services/userlog.service.go`:   tplUri + `templates/v1/apps/user/services/userlog.service.tpl`,
			c.path + `/apps/user/services/userrole.service.go`:  tplUri + `templates/v1/apps/user/services/userrole.service.tpl`,
		},
	}
	moduleNameSlic := strings.Split(c.moduleName, `/`)
	exeName := moduleNameSlic[len(moduleNameSlic)-1]
	c.tplData = map[string]any{
		`moduleName`: c.moduleName,
		`dbName`:     `comer_project`,
		`exeName`:    exeName,
	}
	return true
}
