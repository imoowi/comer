package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func executeCommand(args ...string) (string, error) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return buf.String(), err
}

// chdirTemp 切换到临时目录并返回恢复函数。
func chdirTemp(t *testing.T) (dir string, restore func()) {
	t.Helper()
	dir = t.TempDir()
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	return dir, func() { _ = os.Chdir(old) }
}

func TestRoot(t *testing.T) {
	out, err := executeCommand()
	if err != nil {
		t.Fatalf("root failed: %v", err)
	}
	if !strings.Contains(out, "comer") {
		t.Fatalf("root help missing 'comer': %q", out)
	}
}

func TestVersion(t *testing.T) {
	out, err := executeCommand("version")
	if err != nil {
		t.Fatalf("version failed: %v", err)
	}
	if !strings.Contains(out, "Comer version") {
		t.Fatalf("version output missing version: %q", out)
	}
}

func TestNewV2(t *testing.T) {
	dir, restore := chdirTemp(t)
	defer restore()

	if _, err := executeCommand("new", "github.com/imoowi/example"); err != nil {
		t.Fatalf("new failed: %v", err)
	}
	base := filepath.Join(dir, "github.com", "imoowi", "example")
	if _, err := os.Stat(filepath.Join(base, "go.mod")); err != nil {
		t.Fatalf("go.mod not created: %v", err)
	}
	if _, err := os.Stat(filepath.Join(base, "internal", "controllers", "user.controller.go")); err != nil {
		t.Fatalf("user.controller.go not created: %v", err)
	}
}

func TestAddV2Idempotent(t *testing.T) {
	dir, restore := chdirTemp(t)
	defer restore()

	if _, err := executeCommand("new", "github.com/imoowi/example"); err != nil {
		t.Fatalf("new failed: %v", err)
	}
	base := filepath.Join(dir, "github.com", "imoowi", "example")
	if err := os.Chdir(base); err != nil {
		t.Fatalf("chdir into project: %v", err)
	}

	if _, err := executeCommand("add", "-c=post", "-s=post", "-m=post"); err != nil {
		t.Fatalf("add failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("internal", "controllers", "post.controller.go")); err != nil {
		t.Fatalf("post.controller.go not created: %v", err)
	}
	// 再次 add，应幂等、无报错
	if _, err := executeCommand("add", "-c=post", "-s=post", "-m=post"); err != nil {
		t.Fatalf("second add failed: %v", err)
	}
}

func TestAddV1InjectionIdempotent(t *testing.T) {
	dir, restore := chdirTemp(t)
	defer restore()

	if _, err := executeCommand("new", "-v=1", "github.com/imoowi/example"); err != nil {
		t.Fatalf("new v1 failed: %v", err)
	}
	base := filepath.Join(dir, "github.com", "imoowi", "example")
	if err := os.Chdir(base); err != nil {
		t.Fatalf("chdir into project: %v", err)
	}

	args := []string{"add", "-a=student", "-c=student", "-s=student", "-m=student"}
	if _, err := executeCommand(args...); err != nil {
		t.Fatalf("add v1 failed: %v", err)
	}
	appsGo, err := os.ReadFile("apps/apps.go")
	if err != nil {
		t.Fatalf("read apps.go: %v", err)
	}
	if got := strings.Count(string(appsGo), `apps/student`); got != 1 {
		t.Fatalf("expected 1 student import in apps.go, got %d", got)
	}

	// 幂等：再次 add 不应重复注入
	if _, err := executeCommand(args...); err != nil {
		t.Fatalf("second add v1 failed: %v", err)
	}
	appsGo, _ = os.ReadFile("apps/apps.go")
	if got := strings.Count(string(appsGo), `apps/student`); got != 1 {
		t.Fatalf("expected still 1 student import after idempotent add, got %d", got)
	}
}
