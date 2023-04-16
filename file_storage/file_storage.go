package file_storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boksadawid/godo-cli/task_manager"
)

const (
	configPath = "godo-cli"
	fileName   = "tasks.json"
)

var configDirForTesting string

func SetConfigDirForTesting(dir string) {
	configDirForTesting = dir
}

func GetFilePath() (string, error) {
	configDir := configDirForTesting
	var err error
	if configDir == "" {
		configDir, err = os.UserConfigDir()
		if err != nil {
			return "", fmt.Errorf("error getting user's config directory: %w", err)
		}
	}

	appConfigDir := filepath.Join(configDir, configPath)
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return "", fmt.Errorf("error creating app config directory: %w", err)
	}

	taskFilePath := filepath.Join(appConfigDir, fileName)

	return taskFilePath, nil
}

func SaveToJson(tasks []task_manager.Task) error {
	jsonBytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	filepath, err := GetFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadFromJson() ([]task_manager.Task, error) {
	filepath, err := GetFilePath()
	if err != nil {
		return nil, err
	}

	jsonBytes, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var tasks []task_manager.Task

	err = json.Unmarshal(jsonBytes, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
