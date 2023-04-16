package task_manager

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskManager(t *testing.T) {
	t.Run("Add task", func(t *testing.T) {
		taskMgr := TaskManager{}

		deadline := time.Now().Add(24 * time.Hour)
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", deadline, false)

		assert.Equal(t, 1, len(taskMgr.Tasks))
		assert.Equal(t, 1, taskMgr.Tasks[0].ID)
	})

	t.Run("List not done tasks", func(t *testing.T) {
		taskMgr := TaskManager{}
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", time.Now().Add(24*time.Hour), false)
		taskMgr.Add("Task 2", "Task description 2", "0 0 * * *", time.Now().Add(48*time.Hour), false)
		taskMgr.Update(2, "Task 2", "Task description 2", "0 0 * * *", time.Now().Add(48*time.Hour), false, true)

		notDoneTasks := taskMgr.ListNotDone()
		assert.Equal(t, 1, len(notDoneTasks))
		assert.Equal(t, 1, notDoneTasks[0].ID)
	})

	t.Run("Get task by ID", func(t *testing.T) {
		taskMgr := TaskManager{}
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", time.Now().Add(24*time.Hour), false)

		task := taskMgr.Get(1)
		assert.NotNil(t, task)
		assert.Equal(t, 1, task.ID)
	})

	t.Run("Delete task by ID", func(t *testing.T) {
		taskMgr := TaskManager{}
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", time.Now().Add(24*time.Hour), false)
		taskMgr.Add("Task 2", "Task description 2", "0 0 * * *", time.Now().Add(48*time.Hour), false)

		taskMgr.Delete(1)
		assert.Equal(t, 1, len(taskMgr.Tasks))
		assert.Equal(t, 2, taskMgr.Tasks[0].ID)
	})

	t.Run("Update task by ID", func(t *testing.T) {
		taskMgr := TaskManager{}
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", time.Now().Add(24*time.Hour), false)
		newDeadline := time.Now().Add(72 * time.Hour)
		taskMgr.Update(1, "Updated Task 1", "Updated description 1", "0 0 * * *", newDeadline, true, true)

		updatedTask := taskMgr.Get(1)
		assert.NotNil(t, updatedTask)
		assert.Equal(t, "Updated Task 1", updatedTask.Name)
		assert.Equal(t, "Updated description 1", updatedTask.Description)
		assert.Equal(t, "0 0 * * *", updatedTask.Cron)
		assert.Equal(t, newDeadline, updatedTask.Deadline)
		assert.Equal(t, true, updatedTask.Backlog)
		assert.Equal(t, true, updatedTask.Done)
	})

	t.Run("Ensure nextId generates unique ID", func(t *testing.T) {
		taskMgr := TaskManager{}
		taskMgr.Add("Task 1", "Task description 1", "* * * * *", time.Now().Add(24*time.Hour), false)
		taskMgr.Add("Task 2", "Task description 2", "0 0 * * *", time.Now().Add(48*time.Hour), false)

		nextID := taskMgr.nextId()
		assert.Equal(t, 3, nextID)
	})
}
