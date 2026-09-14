package comer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFrameworkConfig(t *testing.T) {
	// 空路径返回 nil
	cfg, err := loadFrameworkConfig(``)
	if err != nil || cfg != nil {
		t.Fatalf("loadFrameworkConfig(empty) = %v, %v; want nil, nil", cfg, err)
	}

	dir := t.TempDir()
	f := filepath.Join(dir, "comer.json5")
	content := `{ db_name: "my_db", exe_name: "myapp", swagger: { title: "T", version: "2.0", description: "D" } }`
	if err := os.WriteFile(f, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	cfg, err = loadFrameworkConfig(f)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.DBName != `my_db` || cfg.ExeName != `myapp` {
		t.Errorf("unexpected cfg: %+v", cfg)
	}
	if cfg.Swagger.Title != `T` || cfg.Swagger.Version != `2.0` || cfg.Swagger.Description != `D` {
		t.Errorf("unexpected swagger: %+v", cfg.Swagger)
	}

	// 不存在的文件报错
	if _, err := loadFrameworkConfig(filepath.Join(dir, "nope.json5")); err == nil {
		t.Errorf("expected error for missing config file")
	}
}

func TestBuildFrameworkTplDataWithConfig(t *testing.T) {
	d := buildFrameworkTplData(`github.com/imoowi/example`, `example`, nil)
	if d[`dbName`] != `comer_project` || d[`exeName`] != `example` {
		t.Errorf("defaults wrong: %+v", d)
	}
	if d[`swaggerTitle`] != `Comer API` {
		t.Errorf("default swagger title wrong: %v", d[`swaggerTitle`])
	}

	cfg := &FrameworkConfig{DBName: `mydb`, ExeName: `exe`, Swagger: SwaggerConfig{Title: `X API`}}
	d2 := buildFrameworkTplData(`github.com/imoowi/example`, `example`, cfg)
	if d2[`dbName`] != `mydb` || d2[`exeName`] != `exe` {
		t.Errorf("override wrong: %+v", d2)
	}
	if d2[`swaggerTitle`] != `X API` || d2[`swaggerVersion`] != `1.0` {
		t.Errorf("swagger override wrong: title=%v version=%v", d2[`swaggerTitle`], d2[`swaggerVersion`])
	}
}
