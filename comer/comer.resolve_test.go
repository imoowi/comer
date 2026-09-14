package comer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func newFieldCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Flags().StringArray(`field`, nil, ``)
	cmd.Flags().String(`fieldConfig`, ``, ``)
	return cmd
}

func TestResolveFields(t *testing.T) {
	// 空 → 空
	if fields, err := resolveFields(newFieldCmd()); err != nil || len(fields) != 0 {
		t.Fatalf("empty = %+v, %v", fields, err)
	}

	// -f
	cmd := newFieldCmd()
	_ = cmd.Flags().Set(`field`, `title:string:100:标题`)
	fields, err := resolveFields(cmd)
	if err != nil || len(fields) != 1 || fields[0].Name != `Title` {
		t.Fatalf("-f = %+v, %v", fields, err)
	}

	// fieldConfig 文件
	dir := t.TempDir()
	f := filepath.Join(dir, "fields.txt")
	if err := os.WriteFile(f, []byte("title:string:100\ncontent:text\n"), 0644); err != nil {
		t.Fatal(err)
	}
	cmd = newFieldCmd()
	_ = cmd.Flags().Set(`fieldConfig`, f)
	fields, err = resolveFields(cmd)
	if err != nil || len(fields) != 2 {
		t.Fatalf("fieldConfig = %+v, %v", fields, err)
	}

	// 两者结合
	cmd = newFieldCmd()
	_ = cmd.Flags().Set(`fieldConfig`, f)
	_ = cmd.Flags().Set(`field`, `status:int`)
	fields, err = resolveFields(cmd)
	if err != nil || len(fields) != 3 {
		t.Fatalf("combined = %+v, %v", fields, err)
	}
}
