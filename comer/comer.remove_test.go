package comer

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Flags().String(`tplVersion`, ``, ``)
	cmd.Flags().String(`controller`, ``, ``)
	cmd.Flags().String(`service`, ``, ``)
	cmd.Flags().String(`model`, ``, ``)
	cmd.Flags().Bool(`dry-run`, false, ``)
	return cmd
}

func TestRemoveAppRejectV1(t *testing.T) {
	cmd := newRemoveCmd()
	_ = cmd.Flags().Set(`tplVersion`, `1`)
	if err := (&Comer{}).RemoveApp(cmd, nil); err == nil {
		t.Fatal("RemoveApp should reject v1 layout")
	}
}

func TestRemoveAppNoGoMod(t *testing.T) {
	dir := t.TempDir() // 无 go.mod
	old, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(old)

	cmd := newRemoveCmd()
	_ = cmd.Flags().Set(`controller`, `post`)
	if err := (&Comer{}).RemoveApp(cmd, nil); err == nil {
		t.Fatal("RemoveApp should error without go.mod")
	}
}
