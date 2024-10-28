package helpers

import (
	"path/filepath"
)

func Include(path string) []string {
	files, _ := filepath.Glob("admin/views/templates/*.html") //istediğm kriterlere uyan dosyaları bir string olarak geri döndürüyor
	path_files, _ := filepath.Glob("admin/views/" + path + "/*.html")

	for _, file := range path_files {
		files = append(files, file)
	}
	return files
}
