package spin

import (
	"errors"
	"os"
	"path"
	"strings"
)

var Path []string

func initPath() {
	paths := os.Getenv("SPINDIR")
	if paths == "" {
		Path = []string{"./project", "./lib"}
	} else {
		Path = strings.Split(paths, ":")
	}
}

func init() {
	initPath()
}

func FindFile(name string) (string, bool) {
	for _, p := range Path {
		fullName := path.Join(p, name)
		_, err := os.Stat(fullName)
		if !errors.Is(err, os.ErrNotExist) {
			return fullName, true
		}
	}
	return "", false
}

func ReadFile(name string) ([]byte, error) {
	fullName, ok := FindFile(name)
	if !ok {
		return nil, errors.New("file not found")
	}
	return os.ReadFile(fullName)
}
