package main

import (
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fileSystem http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	//Open File
	file, err := nfs.fileSystem.Open(path)
	if err != nil {
		return nil, err
	}

	//Check if the file is a directory
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	//If the file is a directory check for an index.html otherwise return nothing
	if fileInfo.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fileSystem.Open(index); err != nil {
			closeErr := file.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return file, nil
}
