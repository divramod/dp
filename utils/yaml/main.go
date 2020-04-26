package dpyaml

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// GetFile returns a yaml file
func GetFile(filePath string, value interface{}) interface{} {

	filename, _ := filepath.Abs(filePath)
	fmt.Println("filename:", filename)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, value)
	if err != nil {
		panic(err)
	}

	return value
}
