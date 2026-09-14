package comer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAllDirAndTpl(t *testing.T) {
	v := TplSetting{
		Controller: []DirAndTpl{{Dir: "c1", Tpl: "c.tpl"}},
		Service:    []DirAndTpl{{Dir: "s1", Tpl: "s.tpl"}},
		Model:      []DirAndTpl{{Dir: "m1", Tpl: "m.tpl"}},
		Repo:       []DirAndTpl{{Dir: "r1", Tpl: "r.tpl"}},
		Migrate:    []DirAndTpl{{Dir: "g1", Tpl: "g.tpl"}},
		Router:     []DirAndTpl{{Dir: "t1", Tpl: "t.tpl"}},
	}
	if got := allDirAndTpl(v); len(got) != 6 {
		t.Fatalf("allDirAndTpl len = %d, want 6", len(got))
	}
}

func TestAppTplDirs(t *testing.T) {
	v := TplSetting{
		Controller: []DirAndTpl{{Dir: "internal/controllers", Tpl: "c.tpl"}},
		Service:    []DirAndTpl{{Dir: "internal/controllers", Tpl: "s.tpl"}}, // 重复 dir
		Model:      []DirAndTpl{{Dir: " internal/models ", Tpl: "m.tpl"}},     // 带空格
	}
	dirs, err := appTplDirs(v)
	if err != nil {
		t.Fatal(err)
	}
	if len(dirs) != 2 {
		t.Fatalf("appTplDirs len = %d, want 2 (去重): %v", len(dirs), dirs)
	}
	// 空 dir 报错
	if _, err := appTplDirs(TplSetting{Controller: []DirAndTpl{{Dir: "", Tpl: "c.tpl"}}}); err == nil {
		t.Fatal("expected error for empty dir")
	}
}

func TestAppTplFiles(t *testing.T) {
	v := TplSetting{
		Controller: []DirAndTpl{{Dir: "internal/controllers", Tpl: "controller.tpl"}},
		Service:    []DirAndTpl{{Dir: "internal/services", Tpl: "service.tpl"}},
		Model:      []DirAndTpl{{Dir: "internal/models", Tpl: "model.tpl"}},
		Repo:       []DirAndTpl{{Dir: "internal/models", Tpl: "repo.tpl"}},
		Migrate:    []DirAndTpl{{Dir: "internal/db/migrates", Tpl: "migrate.tpl"}},
		Router:     []DirAndTpl{{Dir: "internal/router", Tpl: "router.tpl"}},
	}
	vv := TplVar{ControllerName: "PostPlus", ServiceName: "PostPlus", ModelName: "PostPlus"}
	files := appTplFiles("/tpl", v, vv)
	if len(files) != 6 {
		t.Fatalf("appTplFiles len = %d, want 6", len(files))
	}
	want := map[string]string{
		"./internal/controllers/post_plus.controller.go": "/tpl/controller.tpl",
		"./internal/services/post_plus.service.go":       "/tpl/service.tpl",
		"./internal/models/post_plus.model.go":           "/tpl/model.tpl",
		"./internal/models/post_plus.repo.go":            "/tpl/repo.tpl",
		"./internal/db/migrates/post_plus.migrate.go":    "/tpl/migrate.tpl",
		"./internal/router/post_plus.router.go":          "/tpl/router.tpl",
	}
	for p, tpl := range want {
		if files[p] != tpl {
			t.Errorf("files[%q] = %q, want %q", p, files[p], tpl)
		}
	}
}

func TestValidateTplFiles(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "ok.tpl"), []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := validateTplFiles(dir, TplSetting{Controller: []DirAndTpl{{Dir: "d", Tpl: "ok.tpl"}}}); err != nil {
		t.Fatalf("validateTplFiles should pass: %v", err)
	}
	if err := validateTplFiles(dir, TplSetting{Controller: []DirAndTpl{{Dir: "d", Tpl: "missing.tpl"}}}); err == nil {
		t.Fatal("expected error for missing template")
	}
}
