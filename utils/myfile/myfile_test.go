package myfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsFileExist(t *testing.T) {
	dir := t.TempDir()
	if IsFileExist(filepath.Join(dir, "nope")) {
		t.Errorf("IsFileExist should be false for nonexistent path")
	}
	f := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(f, []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	if !IsFileExist(f) {
		t.Errorf("IsFileExist should be true for existing file")
	}
}

func TestCreateDir(t *testing.T) {
	dir := t.TempDir()
	newDir := filepath.Join(dir, "a", "b")
	if !CreateDir(newDir, false) {
		t.Fatalf("CreateDir should return true for new dir")
	}
	if !IsFileExist(newDir) {
		t.Fatalf("dir should exist after CreateDir")
	}
	// 已存在、remove=false → false
	if CreateDir(newDir, false) {
		t.Errorf("CreateDir should return false when dir exists and remove=false")
	}
	// 已存在、remove=true → 重建并返回 true
	if !CreateDir(newDir, true) {
		t.Errorf("CreateDir should return true when dir exists and remove=true")
	}
	if !IsFileExist(newDir) {
		t.Errorf("dir should be recreated")
	}
}

func TestRemoveDir(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "to-remove")
	if err := os.MkdirAll(target, 0755); err != nil {
		t.Fatal(err)
	}
	if err := RemoveDir(target); err != nil {
		t.Fatalf("RemoveDir: %v", err)
	}
	if IsFileExist(target) {
		t.Errorf("dir should be removed")
	}
}
