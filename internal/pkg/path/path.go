package path

import (
	"os"
	"path/filepath"
)

func BasePath(path ...string) string {
	base, _ := os.Getwd()
	paths := append([]string{base}, path...)

	return filepath.Join(paths...)
}
