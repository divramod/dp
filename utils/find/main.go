package dfind

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/ktr0731/go-fuzzyfinder"
)

// DDir ...
type DDir struct {
	Name    	string
	Size    	int64
	Mode    	os.FileMode
	ModTime 	time.Time
	PathLocal  	string
	PathGlobal 	string
}

// FindDirs ...
func FindDirs(path string) ([]DDir, error) {
	var dirs []DDir
	var selectedDirs []DDir

	// --- [cd] change to search directory
	currentDir, err := os.Getwd()
	if err != nil {
		return selectedDirs, err
	}
	os.Chdir(path)
	pathSearch, err := os.Getwd()
	if err != nil {
		os.Chdir(currentDir)
		return selectedDirs, err
	}
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// fmt.Println("path:", path)
			var dir DDir
			dir.Name = info.Name()
			dir.Size = info.Size()
			dir.Mode = info.Mode()
			dir.ModTime = info.ModTime()
			dir.PathLocal = path
			dir.PathGlobal = pathSearch + "/" + path
			dirs = append(dirs, dir)
		}
		return nil
	})
	if err != nil {
		os.Chdir(currentDir)
		return selectedDirs, err
	}
	idx, err := fuzzyfinder.FindMulti(
		dirs,
		func(i int) string {
			return dirs[i].PathLocal
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Name: %s \nSize: %v\nMode: %s\nTime: %s\nPath local: %s\nPath global: %s",
				dirs[i].Name,
				dirs[i].Size,
				dirs[i].Mode,
				dirs[i].ModTime.Format("2006-01-02 15:04:05"),
				dirs[i].PathLocal,
				dirs[i].PathGlobal,
			)
		}))
	if err != nil {
		os.Chdir(currentDir)
		return selectedDirs, err
	}
	for _, i := range idx {
		selectedDirs= append(selectedDirs, dirs[i])
	}
	os.Chdir(currentDir)
	return selectedDirs, nil
}
// DFile ...
type DFile struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	PathLocal    string
	PathGlobal 		string
}

// FindFiles ...
func FindFiles(path string) ([]DFile, error) {
	var files []DFile
	var selectedFiles []DFile

	// --- [cd] change to search directory
	currentDir, err := os.Getwd()
	if err != nil {
		return selectedFiles, err
	}
	os.Chdir(path)
	pathSearch, err := os.Getwd()
	if err != nil {
		os.Chdir(currentDir)
		return selectedFiles, err
	}
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// fmt.Println("path:", path)
			var file DFile
			file.Name = info.Name()
			file.Size = info.Size()
			file.Mode = info.Mode()
			file.ModTime = info.ModTime()
			file.PathLocal = path
			file.PathGlobal = pathSearch + "/" + path
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		os.Chdir(currentDir)
		return selectedFiles, err
	}
	idx, err := fuzzyfinder.FindMulti(
		files,
		func(i int) string {
			return files[i].PathLocal
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			dat, err := ioutil.ReadFile(files[i].PathLocal)
			if err != nil {
				return "problem reading file"
			}
			return fmt.Sprintf("Name: %s \nSize: %v\nMode: %s\nTime: %s\nPath local: %s\nPath global: %s\nData:\n%s",
				files[i].Name,
				files[i].Size,
				files[i].Mode,
				files[i].ModTime.Format("2006-01-02 15:04:05"),
				files[i].PathLocal,
				files[i].PathGlobal,
				string(dat),
			)
		}))
	if err != nil {
		os.Chdir(currentDir)
		return selectedFiles, err
	}
	for _, i := range idx {
		selectedFiles = append(selectedFiles, files[i])
	}
	os.Chdir(currentDir)
	return selectedFiles, nil
}

