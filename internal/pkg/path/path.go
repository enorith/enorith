package path

import (
	"os"
	"path/filepath"
)

func BasePath(path ...string) string {
	base, _ := filepath.Abs(filepath.Base(os.Args[0]))
	paths := append([]string{base}, path...)

	return filepath.Join(paths...)
}
