/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imoowi/comer/utils/format"
)

// readModuleName 从项目根目录的 go.mod 中读取 module 名称。
func readModuleName() (string, error) {
	if _, err := os.Stat(`go.mod`); os.IsNotExist(err) {
		return ``, fmt.Errorf(`项目根目录下没有 go.mod 文件`)
	}
	file, err := os.Open(`go.mod`)
	if err != nil {
		return ``, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 && fields[0] == `module` {
			return fields[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return ``, err
	}
	return ``, fmt.Errorf(`go.mod 中未找到 module 声明`)
}

// splitModuleName 返回 module 路径的最后一段（即项目名/可执行名）。
func splitModuleName(moduleName string) string {
	parts := strings.Split(moduleName, `/`)
	return parts[len(parts)-1]
}

// buildFrameworkTplData 构建生成项目框架时的模板数据（v1/v2 共用）。
// cfg 非空时用其覆盖 dbName/exeName/swagger 默认值。
func buildFrameworkTplData(moduleName, projectName string, cfg *FrameworkConfig) map[string]any {
	dbName := `comer_project`
	exeName := projectName
	swaggerTitle := `Comer API`
	swaggerVersion := `1.0`
	swaggerDescription := `This is a comer-example server.`

	if cfg != nil {
		if cfg.DBName != `` {
			dbName = cfg.DBName
		}
		if cfg.ExeName != `` {
			exeName = cfg.ExeName
		}
		if cfg.Swagger.Title != `` {
			swaggerTitle = cfg.Swagger.Title
		}
		if cfg.Swagger.Version != `` {
			swaggerVersion = cfg.Swagger.Version
		}
		if cfg.Swagger.Description != `` {
			swaggerDescription = cfg.Swagger.Description
		}
	}

	return map[string]any{
		`moduleName`:        moduleName,
		`dbName`:            dbName,
		`exeName`:           exeName,
		`moduleProjectName`: projectName,
		`swaggerTitle`:       swaggerTitle,
		`swaggerVersion`:     swaggerVersion,
		`swaggerDescription`: swaggerDescription,
	}
}

// buildAppTplData 构建生成单个应用/控制器时的模板数据。
// controllerName 对应 v1 的 handler 概念与 v2 的 controller 概念，两者共用同一套别名。
func buildAppTplData(moduleName, appName, controllerName, serviceName, modelName, swaggerTags string) map[string]any {
	return map[string]any{
		`ModuleName`: moduleName,
		`moduleName`: moduleName,
		`AppName`:    format.FirstUpper(appName),
		`appName`:    strings.ToLower(appName),

		// Handler* 与 Controller* 是同一概念在不同模板版本中的别名，统一提供。
		`HandlerName`:       format.FirstUpper(controllerName),
		`lHandlerName`:      format.FirstLower(controllerName),
		`handlerName`:       strings.ToLower(controllerName),
		`handler_name`:      format.Camel2Snake(controllerName),
		`handler-name`:      format.Camel2Dash(controllerName),
		`handlerName2Dash`:  format.Camel2Dash(controllerName),
		`handlerName2Snake`: format.Camel2Snake(controllerName),

		`ControllerName`:       format.FirstUpper(controllerName),
		`lControllerName`:      format.FirstLower(controllerName),
		`controllerName`:       strings.ToLower(controllerName),
		`controller_name`:      format.Camel2Snake(controllerName),
		`controller-name`:      format.Camel2Dash(controllerName),
		`controllerName2Dash`:  format.Camel2Dash(controllerName),
		`controllerName2Snake`: format.Camel2Snake(controllerName),

		`ServiceName`: format.FirstUpper(serviceName),
		`serviceName`: strings.ToLower(serviceName),
		`ModelName`:   format.FirstUpper(modelName),
		`modelName`:   strings.ToLower(modelName),
		`model_name`:  format.Camel2Snake(modelName),
		`model-name`:  format.Camel2Dash(modelName),
		`SwaggerTags`: swaggerTags,
	}
}

// writeFileAtomic 通过临时文件 + rename 原子地写文件，避免写一半损坏原文件。
func writeFileAtomic(fileName string, data []byte) error {
	tmp := fileName + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	if err := os.Rename(tmp, fileName); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}

// insertIntoImportBlock 在文件的 import ( ... ) 块中追加一行 import，幂等。
func insertIntoImportBlock(fileName, line string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	text := string(data)

	// 幂等：已存在同内容行（trim 后）则跳过
	for _, l := range strings.Split(text, "\n") {
		if strings.TrimSpace(l) == strings.TrimSpace(line) {
			return nil
		}
	}

	start := strings.Index(text, "import (")
	if start < 0 {
		return fmt.Errorf("未找到 import ( 块: %s", fileName)
	}
	closeRel := strings.Index(text[start:], "\n)")
	if closeRel < 0 {
		return fmt.Errorf("import 块未闭合: %s", fileName)
	}
	insertAt := start + closeRel + 1 // 指向闭合括号前一行末尾
	newText := text[:insertAt] + line + "\n" + text[insertAt:]
	return writeFileAtomic(fileName, []byte(newText))
}

// insertBeforeSentinel 在包含 sentinel 子串的那一行之前插入 block（每行一个元素），幂等。
// 若 block 第一行已存在（trim 后精确匹配）则视为已注入，直接返回 nil；找不到哨兵行则报错。
func insertBeforeSentinel(fileName, sentinel string, block []string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")

	first := strings.TrimSpace(block[0])
	for _, l := range lines {
		if strings.TrimSpace(l) == first {
			return nil
		}
	}

	insertIdx := -1
	for i, l := range lines {
		if strings.Contains(l, sentinel) {
			insertIdx = i
			break
		}
	}
	if insertIdx < 0 {
		return fmt.Errorf("未找到哨兵行 %q，无法注入路由", sentinel)
	}

	out := make([]string, 0, len(lines)+len(block))
	out = append(out, lines[:insertIdx]...)
	out = append(out, block...)
	out = append(out, lines[insertIdx:]...)
	return writeFileAtomic(fileName, []byte(strings.Join(out, "\n")))
}
