// Package util is a helper library for go2ts
package util

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// GetGoModPath - Gitリポジトリ直下から幅優先探索でgo.modを探す
func GetGoModPath(in string) (string, error) {
	base := filepath.Clean(in)
	prev := base

	for {
		goMod := filepath.Join(base, "go.mod")

		_, err := os.Stat(goMod)

		if err == nil {
			return goMod, nil
		}

		base, err = filepath.Abs(filepath.Join(base, ".."))

		if err != nil {
			return "", fmt.Errorf("not found")
		}

		if prev == base {
			return "", fmt.Errorf("not found")
		}

		prev = base
	}
}

// GetGoModule - Goのルートパッケージ名をgo.modから取得する
func GetGoModule(goMod string) (string, error) {
	d, err := ioutil.ReadFile(goMod)
	if err != nil {
		return "", err
	}

	f, err := modfile.Parse("", d, nil)
	if err != nil {
		return "", err
	}

	if len(f.Module.Mod.Path) == 0 {
		return "", fmt.Errorf("package name was not found")
	}

	return f.Module.Mod.Path, nil
}

// SplitPackegeStruct - package.structを分割
func SplitPackegeStruct(s string) (string, string) {
	idx := strings.LastIndex(s, ".")

	return s[:idx], s[idx+1:]
}

// GetPackageNameFromPath - パスの最後の要素からパッケージ名を取得
func GetPackageNameFromPath(s string) string {
	arr := strings.Split(s, "/")

	return arr[len(arr)-1]
}

// SHA1 generate SHA1 in hex encoding from s
func SHA1(s string) string {
	hash := sha1.Sum([]byte(s))

	return fmt.Sprintf("%x", hash[:])
}
