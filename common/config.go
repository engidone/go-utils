package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/engidone/utils/log"
	"gopkg.in/yaml.v3"
)

type Paths struct {
	Root   string
	Config string
}

func backLevels(n int) []string {
	backevels := make([]string, n)

	for i := range backevels {
		backevels[i] = ".."
	}

	return backevels
}

func findFileAbsolute(basePath, fileName string) (string, error) {
	var result string

	// Convert basePath to absolute first
	absPath, err := filepath.Abs(basePath)
	if err != nil {
		return "", fmt.Errorf("error getting absolute path of %s: %v", basePath, err)
	}

	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			result = path // path is already full/absolute
			return filepath.SkipDir
		}
		return nil
	})

	if result == "" {
		return "", fmt.Errorf("file '%s' not found in '%s'", fileName, absPath)
	}
	return result, nil
}

func NewConfigPaths(path string) Paths {
	wd, _ := os.Getwd()
	configPaths := strings.Split(path, "/")
	return Paths{
		Config: filepath.Join(wd, filepath.Join(configPaths...)),
		Root:   wd,
	}
}

func GetConfigPath(fileName string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting work dir -> %s", err.Error())
	}

	pathFile, err := findFileAbsolute(wd, fileName)

	if err != nil {
		log.Fatalf("error finding absolute path  for %s file -> %s",  fileName, err.Error())
	}

	return pathFile
}

func LoadFile[T any](path string) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	expandedYaml := os.ExpandEnv(string(data))

	var config T
	if err := yaml.Unmarshal([]byte(expandedYaml), &config); err != nil {
		log.Fatalf("failed to unmarshal yaml: %v", err)
		return nil, err
	}

	return &config, nil
}
