package file_storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/boksadawid/godo-cli/task_manager"
	"github.com/stretchr/testify/assert"
)

func TestFileStorage(t *testing.T) {
	t.Run("Get file path", func(t *testing.T) {
		path, err := GetFilePath()
		assert.NoError(t, err)

		configDir, _ := os.UserConfigDir()
		appConfigDir := filepath.Join(configDir, "godo-cli")
		expectedPath := filepath.Join(appConfigDir, "tasks.json")

		assert.Equal(t, expectedPath, path)
	})

	t.Run("Save and load tasks from JSON", func(t *testing.T) {
		// Create a temporary directory for testing
		tmpDir, err := os.MkdirTemp("", "config")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Set the temporary directory as the config directory for testing
		SetConfigDirForTesting(tmpDir)

		taskMgr := &task_manager.TaskManager{}
		deadline := time.Now().Add(24 * time.Hour)
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", deadline, false)
		taskMgr.Add("Task 2", "Task description 2", "* * * * *", deadline, false)

		err = SaveToJson(taskMgr.Tasks)
		assert.NoError(t, err)

		// Test loading tasks from JSON
		loadedTasks, err := LoadFromJson()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(loadedTasks))
		assert.Equal(t, 1, loadedTasks[0].ID)
		assert.Equal(t, "Task 1", loadedTasks[0].Name)
	})
}
