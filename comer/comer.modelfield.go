/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ModelField 表示一个模型字段。
type ModelField struct {
	Column string // snake_case 列名，如 title
	Name   string // CamelCase 字段名，如 Title
	GoType string // string / int / int64 / float64 / bool / time.Time
	Tag    string // 组合好的反引号标签
}

// typeKeyword 保存 type 关键字到 Go/GORM 类型的映射。
type typeKeyword struct {
	goType   string
	gormType string
}

var typeKeywords = map[string]typeKeyword{
	`string`:   {`string`, `varchar`},
	`text`:     {`string`, `text`},
	`int`:      {`int`, `int`},
	`int64`:    {`int64`, `bigint`},
	`bigint`:   {`int64`, `bigint`},
	`float64`:  {`float64`, `decimal(10,2)`},
	`decimal`:  {`float64`, `decimal(10,2)`},
	`float`:    {`float64`, `decimal(10,2)`},
	`bool`:     {`bool`, `tinyint(1)`},
	`datetime`: {`time.Time`, `datetime`},
	`date`:     {`time.Time`, `date`},
}

// parseModelField 解析单个字段 spec：name:type[:size][:comment]
func parseModelField(spec string) (ModelField, error) {
	parts := strings.Split(spec, `:`)
	if len(parts) < 2 {
		return ModelField{}, fmt.Errorf(`字段定义格式应为 name:type[:size][:comment]，收到: %q`, spec)
	}
	column := strings.TrimSpace(parts[0])
	if column == `` {
		return ModelField{}, fmt.Errorf(`字段名不能为空: %q`, spec)
	}
	typeKey := strings.TrimSpace(parts[1])
	tk, ok := typeKeywords[typeKey]
	if !ok {
		return ModelField{}, fmt.Errorf(`不支持的字段类型 %q（支持: string/text/int/int64/bigint/float64/decimal/float/bool/datetime/date）`, typeKey)
	}

	size := ``
	if len(parts) > 2 {
		size = strings.TrimSpace(parts[2])
	}
	comment := column
	if len(parts) > 3 {
		if c := strings.TrimSpace(parts[3]); c != `` {
			comment = c
		}
	}

	gormType := tk.gormType
	if gormType == `varchar` {
		if size == `` {
			size = `30`
		}
		gormType = `varchar(` + size + `)`
	}

	tag := fmt.Sprintf(`json:"%s" form:"%s" gorm:"column:%s;type:%s;not null;comment:%s"`, column, column, column, gormType, comment)
	return ModelField{
		Column: column,
		Name:   snake2Camel(column),
		GoType: tk.goType,
		Tag:    tag,
	}, nil
}

// parseModelFields 解析多个字段 spec。
func parseModelFields(specs []string) ([]ModelField, error) {
	fields := make([]ModelField, 0, len(specs))
	for _, s := range specs {
		if strings.TrimSpace(s) == `` {
			continue
		}
		f, err := parseModelField(s)
		if err != nil {
			return nil, err
		}
		fields = append(fields, f)
	}
	return fields, nil
}

// parseModelFieldsFile 读取字段配置文件，一行一个字段，跳过空行与 # 注释。
func parseModelFieldsFile(path string) ([]ModelField, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	specs := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == `` || strings.HasPrefix(line, `#`) {
			continue
		}
		specs = append(specs, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return parseModelFields(specs)
}

// resolveFields 先读 fieldConfig 文件，再追加 -f 字段。
func resolveFields(cmd *cobra.Command) ([]ModelField, error) {
	var all []ModelField

	if cfg, err := cmd.Flags().GetString(`fieldConfig`); err != nil {
		return nil, err
	} else if cfg != `` {
		fs, err := parseModelFieldsFile(cfg)
		if err != nil {
			return nil, err
		}
		all = append(all, fs...)
	}

	if specs, err := cmd.Flags().GetStringSlice(`field`); err != nil {
		return nil, err
	} else if len(specs) > 0 {
		fs, err := parseModelFields(specs)
		if err != nil {
			return nil, err
		}
		all = append(all, fs...)
	}

	return all, nil
}

// firstStringColumn 返回第一个 string/text 字段的列名，缺省 "name"。
func firstStringColumn(fields []ModelField) string {
	for _, f := range fields {
		if f.GoType == `string` {
			return f.Column
		}
	}
	return `name`
}

// hasTimeField 判断是否存在 time.Time 字段。
func hasTimeField(fields []ModelField) bool {
	for _, f := range fields {
		if f.GoType == `time.Time` {
			return true
		}
	}
	return false
}

// snake2Camel 将 snake_case 转为 CamelCase（is_active -> IsActive）。
func snake2Camel(s string) string {
	parts := strings.Split(s, `_`)
	var b strings.Builder
	for _, p := range parts {
		if p == `` {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	return b.String()
}
