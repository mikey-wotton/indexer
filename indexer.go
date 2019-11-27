package indexer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"go.uber.org/zap"
)

type directory struct {
	Name        string
	Path        string
	TotalSize   int64
	Directories *[]*directory
	Files       *map[string]os.FileInfo
}

var logger, _ = zap.NewProduction()

//CreateIndex will create an index file of the provided directory at the provided outputPath location
func CreateIndex(directoryPath, outputPath string) error {
	directory, _, err := parseDirectory(directoryPath)
	if err != nil {
		return err
	}

	return parse(outputPath, directory)
}

func parse(output string, dir *directory) error {
	fm := template.FuncMap{"bytesToKiloBytes": func(a int64) int {
		return int(a / 1000)
	}, "bytesToMegaBytes": func(a int64) int {
		return int(a / 1000 / 1000)
	}}

	t, err := template.New("index.tmpl").Funcs(fm).Parse(tmpl)
	if err != nil {
		logger.Error("failed to parse tmpl", zap.Any("err", err))
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		logger.Error("failed to create file", zap.Any("err", err))
		return err
	}
	defer f.Close()

	err = t.Execute(f, *dir)
	if err != nil {
		logger.Error("failed to execute tmpl", zap.Any("err", err))
		return err
	}

	return nil
}

func parseDirectory(dirPath string) (*directory, int64, error) {
	directories := make([]*directory, 0)
	files := make(map[string]os.FileInfo, 0)
	directory := &directory{
		filepath.Base(dirPath),
		dirPath,
		0, //initial directory size
		&directories,
		&files}

	fileList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logger.Error("failed to readDir", zap.Any("err", err))
		return nil, 0, err
	}

	for _, file := range fileList {
		if file.IsDir() {
			directoryName := dirPath + "\\" + file.Name()
			dir, size, err := parseDirectory(directoryName)
			if err != nil {
				logger.Warn(fmt.Sprintf("failed to parse directory named %s, skipping...", directoryName))
				logger.Error(fmt.Sprintf("directory %s failed to parse due to error", directoryName), zap.Any("err", err))
				continue
			}
			directories = append(directories, dir)
			directory.TotalSize += size
		} else {
			files[dirPath+"\\"+file.Name()] = file
			directory.TotalSize += file.Size()
		}
	}

	return directory, directory.TotalSize, nil
}
