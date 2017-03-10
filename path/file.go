package path

import (
	"os"
	"path/filepath"
)

func ReadDirectory(baseDir string) []string {
	fileList := []string{}
	filepath.Walk(baseDir, func(path string, f os.FileInfo, err error) error {
		if ! f.IsDir(){
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}
