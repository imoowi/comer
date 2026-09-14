package comer

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestAppFilePathsV2(t *testing.T) {
	paths := appFilePathsV2("PostPlus", "PostService", "PostModel")
	if len(paths) != 7 {
		t.Fatalf("appFilePathsV2 len = %d, want 7", len(paths))
	}
	if _, ok := paths["./internal/router/post_plus.router.go"]; !ok {
		t.Errorf("missing router path")
	}
	if _, ok := paths["./internal/models/post_model.model.go"]; !ok {
		t.Errorf("missing model path")
	}
	if _, ok := paths["./internal/services/post_service.service.go"]; !ok {
		t.Errorf("missing service path")
	}
}

func TestInitAppV2EmptyController(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(old)
	if err := os.WriteFile("go.mod", []byte("module github.com/foo/bar\n"), 0644); err != nil {
		t.Fatal(err)
	}

	cmd := &cobra.Command{}
	cmd.Flags().String(`controller`, ``, ``)
	if err := (&Comer{}).initAppV2(cmd, nil); err == nil {
		t.Fatal("initAppV2 with empty controller should error")
	}
}
