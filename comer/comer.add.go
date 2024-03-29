/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/imoowi/comer/utils/format"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func (c *Comer) AddApp(cmd *cobra.Command, args []string) {
	/*
		if !c.initApp(cmd, args) {
			return
		}
		//*/

	tplVersion, err := cmd.Flags().GetString(`tplVersion`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	appName, _ := cmd.Flags().GetString(`app`)

	if tplVersion == `1` || appName != `` {
		if !c.initApp(cmd, args) {
			return
		}
		c.showLogo()
		c.generateAppDir()
		c.generateAppFiles()
		c.addAppsDepend()
		c.addAppRouterDepend()
		c.showAppTips()
	} else {
		if !c.initAppV2(cmd, args) {
			return
		}
		c.showLogo()
		c.generateAppDir()
		c.generateAppFiles()
		c.showAppTips()
	}
}

func (c *Comer) generateAppDir() {
	if len(c.App.dirs) > 0 {
		for _, dir := range c.App.dirs {
			c.generateDirByName(dir)
		}
	}
}

func (c *Comer) generateDirs(dirs []string) {
	if len(dirs) > 0 {
		for _, dir := range dirs {
			c.generateDirByName(dir)
		}
	}
}

func (c *Comer) generateAppFiles() {
	if len(c.App.files) > 0 {
		for file, tpl := range c.App.files {
			c.generateFileByMap(file, tpl, c.tplAppData, false)
		}
	}
}
func (c *Comer) generateFiles(files map[string]string, tplData any) {
	if len(files) > 0 {
		for file, tpl := range files {
			c.generateFileByMap(file, tpl, tplData, true)
		}
	}
}

func (c *Comer) showAppTips() {
	fmt.Println(`comer add app end.`)
}

func (c *Comer) addAppsDepend() {
	//*
	appsGenFile := `./apps/apps.go`
	_, gErr := os.Stat(appsGenFile)
	if os.IsNotExist(gErr) {
		//定义模板
		tplFileName := `templates/v1/apps/apps.tpl`
		file, err := os.OpenFile(appsGenFile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("file ["+appsGenFile+"] open failed", err.Error())
			return
		}
		defer file.Close()

		// t, err := template.ParseFiles(tpl)
		t, err := template.ParseFS(tplLocal, tplFileName)
		if err != nil {
			fmt.Println(`err:`, err.Error())
			return
		}
		err = t.Execute(file, c.tplAppData)
		if err != nil {
			fmt.Println(`err=`, err.Error())
		}

		fmt.Println(`file [`, appsGenFile, `] created`)
	} else {
		fmt.Println(`file [`, appsGenFile, `] already exists`)
	}
	dashLine := `	_ "` + cast.ToString(c.tplAppData[`moduleName`]) + `/apps/` + strings.ToLower(cast.ToString(c.tplAppData[`appName`])) + `"`

	//判断路由里有没有对应控制器
	// 打开要操作的file  os.O_RDWR: 可读可写
	file, err := os.OpenFile(appsGenFile, os.O_RDWR, 0544)
	if err != nil {
		fmt.Printf("File open failed! err: %v\n", err)
		return
	}
	reader := bufio.NewReader(file)
	linesPos := make([]string, 0)
	needAddDashLine := true

	for {
		line, err := reader.ReadString('\n') // 依次读一行
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("File raed failed! err: %v\n", err)
			return
		}
		if strings.Contains(line, dashLine) {
			needAddDashLine = false
			continue
		}
		linesPos = append(linesPos, line)
	}

	file.Close()
	if needAddDashLine {
		tempFile, err := os.OpenFile(appsGenFile+".tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Temp create failed! err: %v\n", err)
			return
		}
		writer := bufio.NewWriter(tempFile)
		_ = writer.Flush()

		for _, v := range linesPos {
			// 写入临时file
			_, _ = writer.WriteString(v)
			if needAddDashLine && strings.Contains(v, `import`) {
				_, _ = writer.WriteString(dashLine + "\r\n")
			}
		}
		_ = writer.Flush()

		tempFile.Close()
		err = os.Rename(appsGenFile+".tmp", appsGenFile)
		if err != nil {
			fmt.Printf("Rename file raed failed! err: %v\n", err)
			return
		}
	}
}

func (c *Comer) addAppRouterDepend() {

	//定义模板
	routerTpl := "templates/v1/apps/genapp/router.tpl"
	//创建 router.go
	routerFile := `./apps/` + strings.ToLower(cast.ToString(c.tplAppData[`appName`])) + `/router.go`
	_, gErr := os.Stat(routerFile)
	if os.IsNotExist(gErr) {
		file, err := os.OpenFile(routerFile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("file ["+routerFile+"] open failed", err.Error())
			return
		}
		defer file.Close()

		// t, err := template.ParseFiles(routerTpl)
		t, err := template.ParseFS(tplLocal, routerTpl)
		if err != nil {
			fmt.Println(`err:`, err.Error())
			return
		}
		err = t.Execute(file, c.tplAppData)
		if err != nil {
			fmt.Println(`err=`, err.Error())
		}
		fmt.Println(`file [`, routerFile, `] created`)
	} else {
		fmt.Println(`file [`, routerFile, `] already exists`)

	}

	//判断路由里有没有对应控制器
	// 打开要操作的file  os.O_RDWR: 可读可写
	file, err := os.OpenFile(routerFile, os.O_RDWR, 0544)
	if err != nil {
		fmt.Printf("File open failed! err: %v\n", err)
		return
	}
	reader := bufio.NewReader(file)
	linesPos := make([]string, 0)
	// controllerLine := strings.ToLower(cast.ToString(c.tplAppData[`handlerName`])) + `s := ` + strings.ToLower(cast.ToString(c.tplAppData[`appName`])) + `.Group("/` + strings.ToLower(cast.ToString(c.tplAppData[`handlerName`])) + `s")`
	// handlerName := c.tplAppData[`handlerName`]
	controllerLine := cast.ToString(c.tplAppData[`lHandlerName`]) + `s := ` + `api.Group("/` + strings.ToLower(cast.ToString(c.tplAppData[`handlerName2Dash`])) + `s")`
	needAddDashLine := true

	for {
		line, err := reader.ReadString('\n') // 依次读一行
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("File raed failed! err: %v\n", err)
			return
		}
		if strings.Contains(line, controllerLine) {
			needAddDashLine = false
		}
		if strings.Contains(line, `do-not-delete-this-line`) {
			linesPos = append(linesPos, `	`+controllerLine+"\r\n")
			linesPos = append(linesPos, `	{`+"\r\n")
			linesPos = append(linesPos, `		`+cast.ToString(c.tplAppData[`lHandlerName`])+`s.GET("",handlers.`+format.FirstUpper(cast.ToString(c.tplAppData[`HandlerName`]))+`PageList) //分页列表`+"\r\n")
			linesPos = append(linesPos, `		`+cast.ToString(c.tplAppData[`lHandlerName`])+`s.GET("/:id",handlers.`+format.FirstUpper(cast.ToString(c.tplAppData[`HandlerName`]))+`One) //详情`+"\r\n")
			linesPos = append(linesPos, `		`+cast.ToString(c.tplAppData[`lHandlerName`])+`s.POST("",handlers.`+format.FirstUpper(cast.ToString(c.tplAppData[`HandlerName`]))+`Add) //新增`+"\r\n")
			linesPos = append(linesPos, `		`+cast.ToString(c.tplAppData[`lHandlerName`])+`s.PUT("/:id",handlers.`+format.FirstUpper(cast.ToString(c.tplAppData[`HandlerName`]))+`Update) //更新`+"\r\n")
			linesPos = append(linesPos, `		`+cast.ToString(c.tplAppData[`lHandlerName`])+`s.PATCH("/:id",handlers.`+format.FirstUpper(cast.ToString(c.tplAppData[`HandlerName`]))+`Patch) //部分更新`+"\r\n")
			linesPos = append(linesPos, `		`+cast.ToString(c.tplAppData[`lHandlerName`])+`s.DELETE("/:id",handlers.`+format.FirstUpper(cast.ToString(c.tplAppData[`HandlerName`]))+`Del) //删除，默认为软删除`+"\r\n")
			linesPos = append(linesPos, `	}`+"\r\n")
		}
		linesPos = append(linesPos, line)
	}

	file.Close()
	if needAddDashLine {
		tempFile, err := os.OpenFile(routerFile+".tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Temp create failed! err: %v\n", err)
			return
		}
		writer := bufio.NewWriter(tempFile)
		_ = writer.Flush()

		for _, v := range linesPos {
			// 写入临时file
			_, _ = writer.WriteString(v)
		}
		_ = writer.Flush()

		tempFile.Close()
		err = os.Rename(routerFile+".tmp", routerFile)
		if err != nil {
			fmt.Printf("Rename file raed failed! err: %v\n", err)
			return
		}
	}

}
