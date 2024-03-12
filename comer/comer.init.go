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
	moduleName := ``
	if len(args) > 0 {
		moduleName = args[0]
	}
	// fmt.Println(`args=`, args[0])
	// return false
	// moduleName, err := cmd.Flags().GetString(`module`)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return false
	// }
	if moduleName == `` {
		// fmt.Println(`请输入module,例如: -m=github.com/imoowi/comer-example(pls input module, e.g. -m=github.com/imoowi/comer-example)`)
		fmt.Println(`请输入module,例如: comer new github.com/imoowi/comer-example (pls input module, e.g. comer new github.com/imoowi/comer-example)`)
		return false
	}
	c.moduleName = moduleName
	moduleNameArr := strings.Split(c.moduleName, `/`)
	c.moduleProjectName = moduleNameArr[len(moduleNameArr)-1]

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
			// c.path + `/components`,
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
			c.path + `/test`,
			// c.path + `/utils`,
			// c.path + `/utils/copy`,
			// c.path + `/utils/format`,
			// c.path + `/utils/maker`,
			// c.path + `/utils/myfile`,
			// c.path + `/utils/mytime`,
			// c.path + `/utils/office`,
			// c.path + `/utils/password`,
			// c.path + `/utils/request`,
			// c.path + `/utils/response`,
			// c.path + `/utils/slice`,
			c.path + `/.vscode`,
			c.path + `/docs`,
		},
		files: map[string]string{
			c.path + `/README.md`:                 tplUri + `templates/v1/README.md.tmpl`,
			c.path + `/Makefile`:                  tplUri + `templates/v1/Makefile.tmpl`,
			c.path + `/Dockerfile`:                tplUri + `templates/v1/Dockerfile.tmpl`,
			c.path + `/docker-compose.yml`:        tplUri + `templates/v1/docker-compose.yml.tmpl`,
			c.path + `/start_server_in_docker.sh`: tplUri + `templates/v1/start_server_in_docker.sh.tmpl`,
			c.path + `/go.mod`:                    tplUri + `templates/v1/go.mod.tmpl`,
			c.path + `/main.go`:                   tplUri + `templates/v1/main.tmpl`,
			c.path + `/.air.toml`:                 tplUri + `templates/v1/air.tmpl`,
			c.path + `/.gitignore`:                tplUri + `templates/v1/gitignore.tmpl`,
			c.path + `/docs/doc.go`:               tplUri + `templates/v1/doc.tmpl`,
			c.path + `/.vscode/launch.json`:       tplUri + `templates/v1/vscode/launch.json`,
			c.path + `/.vscode/settings.json`:     tplUri + `templates/v1/vscode/settings.json`,
			c.path + `/cmd/root.go`:               tplUri + `templates/v1/cmd/cmd.root.tmpl`,
			c.path + `/cmd/server.go`:             tplUri + `templates/v1/cmd/cmd.server.tmpl`,
			c.path + `/cmd/init.go`:               tplUri + `templates/v1/cmd/cmd.init.tmpl`,
			c.path + `/cmd/migrate.go`:            tplUri + `templates/v1/cmd/cmd.migrate.tmpl`,
			// c.path + `/components/mysql.go`:                tplUri + `templates/v1/components/components.mysql.tmpl`,
			// c.path + `/components/redis.go`:                tplUri + `templates/v1/components/components.redis.tmpl`,
			// c.path + `/components/captcha.go`:              tplUri + `templates/v1/components/components.captcha.tmpl`,
			c.path + `/configs/settings-local.yml`:         tplUri + `templates/v1/config/config.settings-local.tmpl`,
			c.path + `/configs/casbin.conf`:                tplUri + `templates/v1/config/config.casbin.tmpl`,
			c.path + `/global/global.go`:                   tplUri + `templates/v1/global/global.tmpl`,
			c.path + `/global/global.userlog.go`:           tplUri + `templates/v1/global/global.userlog.tmpl`,
			c.path + `/global/cache.go`:                    tplUri + `templates/v1/global/global.cache.tmpl`,
			c.path + `/global/casbin.go`:                   tplUri + `templates/v1/global/global.casbin.tmpl`,
			c.path + `/global/config.go`:                   tplUri + `templates/v1/global/global.config.tmpl`,
			c.path + `/global/log.go`:                      tplUri + `templates/v1/global/global.log.tmpl`,
			c.path + `/global/mysql.go`:                    tplUri + `templates/v1/global/global.mysql.tmpl`,
			c.path + `/global/redis.go`:                    tplUri + `templates/v1/global/global.redis.tmpl`,
			c.path + `/middlewares/CasbinMiddleware.go`:    tplUri + `templates/v1/middleware/middleware.Casbin.tmpl`,
			c.path + `/middlewares/CrosMiddleware.go`:      tplUri + `templates/v1/middleware/middleware.Cros.tmpl`,
			c.path + `/middlewares/JWTAuthMiddleware.go`:   tplUri + `templates/v1/middleware/middleware.JWTAuth.tmpl`,
			c.path + `/middlewares/token/jwttoken.go`:      tplUri + `templates/v1/middleware/token/middleware.JWTToken.tmpl`,
			c.path + `/middlewares/LoggerMiddleware.go`:    tplUri + `templates/v1/middleware/middleware.Logger.tmpl`,
			c.path + `/middlewares/RateLimitMiddleware.go`: tplUri + `templates/v1/middleware/middleware.RateLimit.tmpl`,
			c.path + `/middlewares/VcodeMiddleware.go`:     tplUri + `templates/v1/middleware/middleware.Vcode.tmpl`,
			c.path + `/middlewares/UserlogMiddleware.go`:   tplUri + `templates/v1/middleware/middleware.Userlog.tmpl`,
			c.path + `/middlewares/RequestIdMiddleware.go`: tplUri + `templates/v1/middleware/middleware.Requestid.tmpl`,
			c.path + `/middlewares/middleware.go`:          tplUri + `templates/v1/middleware/middleware.tmpl`,
			c.path + `/router/router.go`:                   tplUri + `templates/v1/router/router.tmpl`,
			// c.path + `/utils/copy/copy.go`:                      tplUri + `templates/v1/utils/copy/copy.tmpl`,
			// c.path + `/utils/format/format.go`:                  tplUri + `templates/v1/utils/format/format.tmpl`,
			// c.path + `/utils/maker/maker.go`:                    tplUri + `templates/v1/utils/maker/maker.tmpl`,
			// c.path + `/utils/myfile/myfile.go`:                  tplUri + `templates/v1/utils/myfile/myfile.tmpl`,
			// c.path + `/utils/mytime/mytime.go`:                  tplUri + `templates/v1/utils/mytime/mytime.tmpl`,
			// c.path + `/utils/mytime/translater.go`:              tplUri + `templates/v1/utils/mytime/translater.tmpl`,
			// c.path + `/utils/mytime/week.go`:                    tplUri + `templates/v1/utils/mytime/week.tmpl`,
			// c.path + `/utils/office/excel.go`:                   tplUri + `templates/v1/utils/office/excel.tmpl`,
			// c.path + `/utils/password/password.go`:              tplUri + `templates/v1/utils/password/password.tmpl`,
			// c.path + `/utils/request/http.go`:                   tplUri + `templates/v1/utils/request/http.tmpl`,
			// c.path + `/utils/request/pages.go`:                  tplUri + `templates/v1/utils/request/pages.tmpl`,
			// c.path + `/utils/response/pages.go`:                 tplUri + `templates/v1/utils/response/pages.tmpl`,
			// c.path + `/utils/response/response.go`:              tplUri + `templates/v1/utils/response/response.tmpl`,
			// c.path + `/utils/slice/slice.go`:                    tplUri + `templates/v1/utils/slice/slice.tmpl`,
			// c.path + `/utils/logger.go`:                         tplUri + `templates/v1/utils/logger.tmpl`,
			// c.path + `/utils/utils.go`:                          tplUri + `templates/v1/utils/utils.tmpl`,
			c.path + `/apps/apps.go`:                            tplUri + `templates/v1/apps/apps.tmpl`,
			c.path + `/apps/swagger/router.go`:                  tplUri + `templates/v1/apps/swagger/router.tmpl`,
			c.path + `/apps/common/router.go`:                   tplUri + `templates/v1/apps/common/router.tmpl`,
			c.path + `/apps/common/handlers/captcha.handler.go`: tplUri + `templates/v1/apps/common/handlers/captcha.handler.tmpl`,
			c.path + `/apps/user/router.go`:                     tplUri + `templates/v1/apps/user/router.tmpl`,
			c.path + `/apps/user/handlers/auth.handler.go`:      tplUri + `templates/v1/apps/user/handlers/auth.handler.tmpl`,
			c.path + `/apps/user/migrates/role.migrate.go`:      tplUri + `templates/v1/apps/user/migrates/role.migrate.tmpl`,
			c.path + `/apps/user/migrates/user.migrate.go`:      tplUri + `templates/v1/apps/user/migrates/user.migrate.tmpl`,
			c.path + `/apps/user/migrates/userlog.migrate.go`:   tplUri + `templates/v1/apps/user/migrates/userlog.migrate.tmpl`,
			c.path + `/apps/user/migrates/userrole.migrate.go`:  tplUri + `templates/v1/apps/user/migrates/userrole.migrate.tmpl`,
			c.path + `/apps/user/models/role.model.go`:          tplUri + `templates/v1/apps/user/models/role.model.tmpl`,
			c.path + `/apps/user/models/user.model.go`:          tplUri + `templates/v1/apps/user/models/user.model.tmpl`,
			c.path + `/apps/user/models/userlog.model.go`:       tplUri + `templates/v1/apps/user/models/userlog.model.tmpl`,
			c.path + `/apps/user/models/userrole.model.go`:      tplUri + `templates/v1/apps/user/models/userrole.model.tmpl`,
			c.path + `/apps/user/repos/role.repo.go`:            tplUri + `templates/v1/apps/user/repos/role.repo.tmpl`,
			c.path + `/apps/user/repos/user.repo.go`:            tplUri + `templates/v1/apps/user/repos/user.repo.tmpl`,
			c.path + `/apps/user/repos/userlog.repo.go`:         tplUri + `templates/v1/apps/user/repos/userlog.repo.tmpl`,
			c.path + `/apps/user/repos/userrole.repo.go`:        tplUri + `templates/v1/apps/user/repos/userrole.repo.tmpl`,
			c.path + `/apps/user/services/role.service.go`:      tplUri + `templates/v1/apps/user/services/role.service.tmpl`,
			c.path + `/apps/user/services/user.service.go`:      tplUri + `templates/v1/apps/user/services/user.service.tmpl`,
			c.path + `/apps/user/services/userlog.service.go`:   tplUri + `templates/v1/apps/user/services/userlog.service.tmpl`,
			c.path + `/apps/user/services/userrole.service.go`:  tplUri + `templates/v1/apps/user/services/userrole.service.tmpl`,
			c.path + `/test/login.go`:                           tplUri + `templates/v1/test/login.tmpl`,
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
