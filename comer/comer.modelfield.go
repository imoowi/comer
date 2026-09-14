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
	Column  string // snake_case 列名，如 title
	Name    string // CamelCase 字段名，如 Title
	GoType  string // string / int / int64 / float64 / bool / time.Time / json.RawMessage / []string
	Tag     string // 组合好的反引号标签（json/form/gorm）
	BindTag string // go-playground/validator 校验规则，如 required,min=2；为空表示不校验
}

// fieldKind 表示字段类型所属分类，用于按类别应用 size 语义。
type fieldKind string

const (
	kindVarchar  fieldKind = `varchar`
	kindText     fieldKind = `text`
	kindDecimal  fieldKind = `decimal`
	kindInt      fieldKind = `int`
	kindFloat    fieldKind = `float`
	kindBool     fieldKind = `bool`
	kindDatetime fieldKind = `datetime`
	kindDate     fieldKind = `date`
	kindJSON     fieldKind = `json`
	kindSlice    fieldKind = `slice`
)

// typeKeyword 保存 type 关键字到 Go/GORM 类型及分类的映射。
type typeKeyword struct {
	goType   string
	gormType string
	kind     fieldKind
}

var typeKeywords = map[string]typeKeyword{
	`string`:   {`string`, `varchar`, kindVarchar},
	`text`:     {`string`, `text`, kindText},
	`int`:      {`int`, `int`, kindInt},
	`int64`:    {`int64`, `bigint`, kindInt},
	`bigint`:   {`int64`, `bigint`, kindInt},
	`uint`:     {`uint`, `int unsigned`, kindInt},
	`uint8`:    {`uint8`, `tinyint unsigned`, kindInt},
	`uint16`:   {`uint16`, `smallint unsigned`, kindInt},
	`uint32`:   {`uint32`, `int unsigned`, kindInt},
	`uint64`:   {`uint64`, `bigint unsigned`, kindInt},
	`float64`:  {`float64`, `decimal`, kindDecimal},
	`decimal`:  {`float64`, `decimal`, kindDecimal},
	`float`:    {`float64`, `decimal`, kindDecimal},
	`float32`:  {`float32`, `float`, kindFloat},
	`bool`:     {`bool`, `tinyint(1)`, kindBool},
	`datetime`: {`time.Time`, `datetime`, kindDatetime},
	`date`:     {`time.Time`, `date`, kindDate},
	`time`:     {`time.Time`, `datetime`, kindDatetime},
	`json`:     {`json.RawMessage`, `json`, kindJSON},
	`slice`:    {`[]string`, `json`, kindSlice},
	`array`:    {`[]string`, `json`, kindSlice},
}

// supportedTypes 返回支持的类型列表，用于错误提示。
func supportedTypes() string {
	return `string/text/int/int64/bigint/uint/uint8/uint16/uint32/uint64/float64/float32/decimal/float/bool/datetime/date/time/json/slice/array`
}

// parseModelField 解析单个字段 spec：name:type[:size][:comment][:validate]
func parseModelField(spec string) (ModelField, error) {
	parts := strings.Split(spec, `:`)
	if len(parts) < 2 {
		return ModelField{}, fmt.Errorf(`字段定义格式应为 name:type[:size][:comment][:validate]，收到: %q`, spec)
	}
	column := strings.TrimSpace(parts[0])
	if column == `` {
		return ModelField{}, fmt.Errorf(`字段名不能为空: %q`, spec)
	}
	typeKey := strings.TrimSpace(parts[1])
	tk, ok := typeKeywords[typeKey]
	if !ok {
		return ModelField{}, fmt.Errorf(`不支持的字段类型 %q（支持: %s）`, typeKey, supportedTypes())
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
	bindTag := ``
	if len(parts) > 4 {
		bindTag = strings.TrimSpace(parts[4])
	}

	gormType := tk.gormType
	switch tk.kind {
	case kindVarchar:
		if size == `` {
			size = `30`
		}
		gormType = `varchar(` + size + `)`
	case kindText:
		gormType = textType(size)
	case kindDecimal:
		gormType = `decimal(` + decimalPrecisionScale(size) + `)`
	}

	tag := fmt.Sprintf(`json:"%s" form:"%s" gorm:"column:%s;type:%s;not null;comment:%s`, column, column, column, gormType, comment)
	if tk.kind == kindJSON || tk.kind == kindSlice {
		tag += `;serializer:json`
	}
	tag += `"`

	return ModelField{
		Column:  column,
		Name:    snake2Camel(column),
		GoType:  tk.goType,
		Tag:     tag,
		BindTag: bindTag,
	}, nil
}

// textType 将 size 映射到 MySQL text 子类型，缺省 text。
func textType(size string) string {
	switch strings.ToLower(size) {
	case `tiny`, `tinytext`:
		return `tinytext`
	case `medium`, `mediumtext`:
		return `mediumtext`
	case `long`, `longtext`:
		return `longtext`
	default:
		return `text`
	}
}

// decimalPrecisionScale 返回 decimal 的 precision,scale，缺省 10,2。
// 分隔符兼容逗号（如 10,4）与点号（如 10.4）。
func decimalPrecisionScale(size string) string {
	if size == `` {
		return `10,2`
	}
	size = strings.ReplaceAll(size, `.`, `,`)
	if strings.Contains(size, `,`) {
		return size
	}
	return size + `,2`
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

	if specs, err := cmd.Flags().GetStringArray(`field`); err != nil {
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

// hasJSONField 判断是否存在 json.RawMessage 字段（需要 encoding/json import）。
func hasJSONField(fields []ModelField) bool {
	for _, f := range fields {
		if f.GoType == `json.RawMessage` {
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
