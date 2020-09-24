package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
