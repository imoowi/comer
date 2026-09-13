/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func (c *Comer) AddApp(cmd *cobra.Command, args []string) error {
	tplVersion, err := cmd.Flags().GetString(`tplVersion`)
	if err != nil {
		return err
	}
	appName, _ := cmd.Flags().GetString(`app`)

	if tplVersion == `1` || appName != `` {
		if tplVersion != `1` && appName != `` {
			fmt.Println(`提示：未指定 -v=1 但传入了 -a，将使用 v1(apps/) 布局`)
		}
		if err = c.initApp(cmd, args); err != nil {
			return err
		}
		c.showLogo()
		if err = c.generateAppDir(); err != nil {
			return err
		}
		if err = c.generateAppFiles(); err != nil {
			return err
		}
		if err = c.addAppsDepend(); err != nil {
			return err
		}
		if err = c.addAppRouterDepend(); err != nil {
			return err
		}
		c.showAppTips()
		return nil
	}

	if err = c.initAppV2(cmd, args); err != nil {
		return err
	}
	c.showLogo()
	if err = c.generateAppDir(); err != nil {
		return err
	}
	if err = c.generateAppFiles(); err != nil {
		return err
	}
	c.showAppTips()
	return nil
}

func (c *Comer) generateAppDir() error {
	for _, dir := range c.App.dirs {
		if err := c.generateDirByName(dir); err != nil {
			return err
		}
	}
	return nil
}

func (c *Comer) generateDirs(dirs []string) error {
	for _, dir := range dirs {
		if err := c.generateDirByName(dir); err != nil {
			return err
		}
	}
	return nil
}

func (c *Comer) generateAppFiles() error {
	for file, tpl := range c.App.files {
		if err := c.generateFileByMap(file, tpl, c.tplAppData, false); err != nil {
			return err
		}
	}
	return nil
}

func (c *Comer) generateFiles(files map[string]string, tplData any) error {
	for file, tpl := range files {
		if err := c.generateFileByMap(file, tpl, tplData, true); err != nil {
			return err
		}
	}
	return nil
}

func (c *Comer) showAppTips() {
	fmt.Println(`comer add app end.`)
}

// addAppsDepend 向共享的 apps.go 注入当前 app 的空白导入，保证 app 被自动加载。
func (c *Comer) addAppsDepend() error {
	appsFile := `./apps/apps.go`
	moduleName := cast.ToString(c.tplAppData[`moduleName`])
	appName := cast.ToString(c.tplAppData[`appName`])
	dashLine := "\t_ \"" + moduleName + "/apps/" + appName + "\""

	// 若 apps.go 不存在，先从模板渲染
	if err := c.generateFileByMap(appsFile, `templates/v1/apps/apps.tmpl`, c.tplAppData, false); err != nil {
		return err
	}

	return insertIntoImportBlock(appsFile, dashLine)
}

// addAppRouterDepend 向 app 的 router.go 追加一个 controller 的路由组。
// router.go 本身由 generateAppFiles 从 router.tmpl 渲染（含首个 controller 的路由），
// 本方法用于给已存在的 app 追加额外的 controller。
func (c *Comer) addAppRouterDepend() error {
	appName := cast.ToString(c.tplAppData[`appName`])
	routerFile := `./apps/` + appName + `/router.go`

	// 极端情况下 router.go 缺失时先从模板渲染
	if err := c.generateFileByMap(routerFile, `templates/v1/apps/genapp/router.tmpl`, c.tplAppData, false); err != nil {
		return err
	}

	lHandler := cast.ToString(c.tplAppData[`lHandlerName`])
	handlerName2Dash := cast.ToString(c.tplAppData[`handlerName2Dash`])
	handlerName := cast.ToString(c.tplAppData[`HandlerName`]) // 已是首字母大写

	controllerLine := lHandler + `s := api.Group("/` + handlerName2Dash + `s")`
	block := []string{
		"\t" + controllerLine,
		"\t{",
		"\t\t" + lHandler + `s.GET("",handlers.` + handlerName + `PageList) //分页列表`,
		"\t\t" + lHandler + `s.GET("/:id",handlers.` + handlerName + `One) //详情`,
		"\t\t" + lHandler + `s.POST("",handlers.` + handlerName + `Add) //新增`,
		"\t\t" + lHandler + `s.PUT("/:id",handlers.` + handlerName + `Update) //更新`,
		"\t\t" + lHandler + `s.PATCH("/:id",handlers.` + handlerName + `Patch) //部分更新`,
		"\t\t" + lHandler + `s.DELETE("/:id",handlers.` + handlerName + `Del) //删除，默认为软删除`,
		"\t}",
	}
	return insertBeforeSentinel(routerFile, `do-not-delete-this-line`, block)
}
