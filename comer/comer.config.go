/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"os"

	"github.com/yosuke-furukawa/json5/encoding/json5"
)

// SwaggerConfig 表示 swagger 文档元信息。
type SwaggerConfig struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// FrameworkConfig 表示 comer new 的可选配置，用于覆盖默认的数据库名/可执行名/swagger 元信息。
type FrameworkConfig struct {
	DBName  string        `json:"db_name"`
	ExeName string        `json:"exe_name"`
	Swagger SwaggerConfig `json:"swagger"`
}

// loadFrameworkConfig 读取并解析 JSON5 配置文件；path 为空时返回 nil。
func loadFrameworkConfig(path string) (*FrameworkConfig, error) {
	if path == `` {
		return nil, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(`[读取配置文件失败] %w`, err)
	}
	var cfg FrameworkConfig
	if err = json5.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf(`[解析配置文件失败] %w`, err)
	}
	return &cfg, nil
}
