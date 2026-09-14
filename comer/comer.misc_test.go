package comer

import (
	"runtime"
	"testing"

	"github.com/spf13/cobra"
)

func TestShowTips(t *testing.T) {
	c := &Comer{path: "github.com/imoowi/example"}
	if got := c.showTips(); got != "showTips called" {
		t.Errorf("showTips = %q, want %q", got, "showTips called")
	}
}

func TestGoVersion(t *testing.T) {
	c := &Comer{}
	if got := c.goVersion(); got != runtime.Version() {
		t.Errorf("goVersion = %q, want %q", got, runtime.Version())
	}
}

func TestInitEmptyModule(t *testing.T) {
	if err := (&Comer{}).init(&cobra.Command{}, []string{}); err == nil {
		t.Fatal("init with empty module should error")
	}
}

func TestInitV2EmptyModule(t *testing.T) {
	if err := (&Comer{}).initV2(&cobra.Command{}, []string{}); err == nil {
		t.Fatal("initV2 with empty module should error")
	}
}
