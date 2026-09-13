package comer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSplitModuleName(t *testing.T) {
	if got := splitModuleName("github.com/imoowi/comer-example"); got != "comer-example" {
		t.Fatalf("got %q", got)
	}
}

func TestBuildAppTplData(t *testing.T) {
	d := buildAppTplData("github.com/Foo/Bar", "MyApp", "PostPlus", "PostService", "PostModel", "Tags")

	if d["moduleName"] != "github.com/Foo/Bar" {
		t.Fatalf("moduleName should stay exact, got %v", d["moduleName"])
	}
	if d["ModuleName"] != "github.com/Foo/Bar" {
		t.Fatalf("ModuleName wrong: %v", d["ModuleName"])
	}
	if d["handler_name"] != "post_plus" {
		t.Fatalf("handler_name should be snake_case, got %v", d["handler_name"])
	}
	if d["handler-name"] != "post-plus" {
		t.Fatalf("handler-name should be dash-case, got %v", d["handler-name"])
	}
	// Handler* 与 Controller* 是同一概念的别名
	if d["HandlerName"] != "PostPlus" || d["ControllerName"] != "PostPlus" {
		t.Fatalf("HandlerName/ControllerName mismatch: %v / %v", d["HandlerName"], d["ControllerName"])
	}
	if d["lHandlerName"] != "postPlus" || d["lControllerName"] != "postPlus" {
		t.Fatalf("lHandler/lController mismatch")
	}
	if d["appName"] != "myapp" {
		t.Fatalf("appName wrong: %v", d["appName"])
	}
}

func TestReadModuleName(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(old)

	if err := os.WriteFile("go.mod", []byte("module github.com/foo/bar\n\ngo 1.20\n"), 0644); err != nil {
		t.Fatal(err)
	}
	m, err := readModuleName()
	if err != nil {
		t.Fatalf("readModuleName: %v", err)
	}
	if m != "github.com/foo/bar" {
		t.Fatalf("got %q", m)
	}
}

func TestInsertIntoImportBlock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "apps.go")
	if err := os.WriteFile(f, []byte("package apps\n\nimport (\n\t_ \"a/b/common\"\n)\n"), 0644); err != nil {
		t.Fatal(err)
	}

	line := "\t_ \"a/b/student\""
	if err := insertIntoImportBlock(f, line); err != nil {
		t.Fatalf("insert: %v", err)
	}
	data, _ := os.ReadFile(f)
	if !strings.Contains(string(data), `"a/b/student"`) {
		t.Fatalf("insert not found: %s", data)
	}
	// 幂等：再次插入不应产生重复
	if err := insertIntoImportBlock(f, line); err != nil {
		t.Fatalf("reinsert: %v", err)
	}
	data2, _ := os.ReadFile(f)
	if got := strings.Count(string(data2), "a/b/student"); got != 1 {
		t.Fatalf("duplicate insert, count=%d: %s", got, data2)
	}
}

func TestInsertBeforeSentinel(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "router.go")
	if err := os.WriteFile(f, []byte("package x\n\nfunc R() {\n\t//!important:do-not-delete-this-line\n}\n"), 0644); err != nil {
		t.Fatal(err)
	}

	block := []string{"\tstudents := api.Group(\"/students\")", "\t{"}
	if err := insertBeforeSentinel(f, "do-not-delete-this-line", block); err != nil {
		t.Fatalf("insert: %v", err)
	}
	data, _ := os.ReadFile(f)
	if !strings.Contains(string(data), "students := api.Group") {
		t.Fatalf("block not inserted: %s", data)
	}
	// 幂等：block 首行已存在则不重复注入
	if err := insertBeforeSentinel(f, "do-not-delete-this-line", block); err != nil {
		t.Fatalf("reinsert: %v", err)
	}
	data2, _ := os.ReadFile(f)
	if got := strings.Count(string(data2), "students := api.Group"); got != 1 {
		t.Fatalf("duplicate insert, count=%d", got)
	}

	// 哨兵缺失时应报错（而非静默 no-op）
	if err := os.WriteFile(f, []byte("package x\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := insertBeforeSentinel(f, "do-not-delete-this-line", block); err == nil {
		t.Fatalf("expected error for missing sentinel")
	}
}
