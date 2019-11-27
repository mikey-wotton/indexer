package indexer

import (
	"os"
)

type directory struct {
	Directories *[]*directory
	Files       *map[string]os.FileInfo
}

//CreateIndex will create an index file of the provided directory at the provided outputPath location
func CreateIndex(directoryPath, outputPath string) error {
	directory, err := parseDirectory(directoryPath)
	if err != nil {
		return err
	}

	err = parseIndex(directory)
	if err != nil {
		return err
	}

	return nil
}

func parseDirectory(dirPath string) (*directory, error) {
	return nil, nil
}

func parseIndex(dir *directory) error {
	return nil
}
