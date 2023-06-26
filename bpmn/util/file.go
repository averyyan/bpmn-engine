package bpmn_util

import (
	"fmt"
	"io/fs"
)

func ListDirFiles(path string, fs fs.ReadDirFS) ([]string, error) {
	var paths []string
	files, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			_paths, err := ListDirFiles(fmt.Sprintf("%s/%s", path, file.Name()), fs)
			if err != nil {

				fmt.Println("paths err", paths, err)
				return nil, err
			}
			paths = append(paths, _paths...)
		} else {
			paths = append(paths, fmt.Sprintf("%s/%s", path, file.Name()))
		}
	}
	return paths, nil
}
