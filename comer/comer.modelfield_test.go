package comer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseModelField(t *testing.T) {
	cases := []struct {
		spec   string
		name   string
		goType string
		gorm   string // 期望的 gorm type 片段
	}{
		{`title:string:100`, `Title`, `string`, `type:varchar(100)`},
		{`content:text`, `Content`, `string`, `type:text`},
		{`status:int`, `Status`, `int`, `type:int`},
		{`count:int64`, `Count`, `int64`, `type:bigint`},
		{`price:decimal`, `Price`, `float64`, `type:decimal(10,2)`},
		{`is_active:bool`, `IsActive`, `bool`, `type:tinyint(1)`},
		{`created_at:datetime`, `CreatedAt`, `time.Time`, `type:datetime`},
		{`birthday:date`, `Birthday`, `time.Time`, `type:date`},
	}
	for _, c := range cases {
		f, err := parseModelField(c.spec)
		if err != nil {
			t.Fatalf("parseModelField(%q) error: %v", c.spec, err)
		}
		if f.Name != c.name {
			t.Errorf("parseModelField(%q).Name = %q, want %q", c.spec, f.Name, c.name)
		}
		if f.GoType != c.goType {
			t.Errorf("parseModelField(%q).GoType = %q, want %q", c.spec, f.GoType, c.goType)
		}
		if !strings.Contains(f.Tag, c.gorm) {
			t.Errorf("parseModelField(%q).Tag = %q, want contain %q", c.spec, f.Tag, c.gorm)
		}
	}
}

func TestParseModelFieldComment(t *testing.T) {
	// 带自定义 comment
	f, err := parseModelField(`title:string:100:标题`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(f.Tag, `comment:标题`) {
		t.Errorf("comment not applied: %q", f.Tag)
	}
	// comment 缺省用字段名
	f2, _ := parseModelField(`title:string`)
	if !strings.Contains(f2.Tag, `comment:title`) {
		t.Errorf("default comment should be column name: %q", f2.Tag)
	}
}

func TestParseModelFieldError(t *testing.T) {
	for _, spec := range []string{``, `onlyname`, `:string`, `title:unknown`} {
		if _, err := parseModelField(spec); err == nil {
			t.Errorf("parseModelField(%q) should error", spec)
		}
	}
}

func TestParseModelFieldsFile(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "model.fields")
	content := "# 注释\n\ntitle:string:100:标题\ncontent:text\n"
	if err := os.WriteFile(f, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	fields, err := parseModelFieldsFile(f)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("got %d fields, want 2", len(fields))
	}
	if fields[0].Name != `Title` || fields[1].Name != `Content` {
		t.Errorf("unexpected fields: %+v", fields)
	}
}

func TestSnake2Camel(t *testing.T) {
	cases := map[string]string{
		`title`:     `Title`,
		`is_active`: `IsActive`,
		`a_b_c`:     `ABC`,
	}
	for in, want := range cases {
		if got := snake2Camel(in); got != want {
			t.Errorf("snake2Camel(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFirstStringColumnAndHasTime(t *testing.T) {
	fields := []ModelField{
		{Column: `status`, GoType: `int`},
		{Column: `title`, GoType: `string`},
		{Column: `created_at`, GoType: `time.Time`},
	}
	if got := firstStringColumn(fields); got != `title` {
		t.Errorf("firstStringColumn = %q, want title", got)
	}
	if !hasTimeField(fields) {
		t.Errorf("hasTimeField should be true")
	}
	if got := firstStringColumn(nil); got != `name` {
		t.Errorf("firstStringColumn(nil) = %q, want name", got)
	}
	if hasTimeField(nil) {
		t.Errorf("hasTimeField(nil) should be false")
	}
}
