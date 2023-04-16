package task_manager

import "time"

type Task struct {
	ID          int
	Name        string
	Description string
	Cron        string
	Deadline    time.Time
	Backlog     bool
	Done        bool
}

type TaskManager struct {
	Tasks []Task
}

func (t *TaskManager) ListNotDone() []Task {
	notDoneTasks := []Task{}
	for _, task := range t.Tasks {
		if !task.Done {
			notDoneTasks = append(notDoneTasks, task)
		}
	}
	return notDoneTasks
}

func (t *TaskManager) Get(id int) *Task {
	for _, task := range t.Tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}

func (t *TaskManager) Add(name string, description string, cron string, deadline time.Time, backlog bool) *TaskManager {
	id := t.nextId()

	task := Task{
		ID:          id,
		Name:        name,
		Description: description,
		Cron:        cron,
		Deadline:    deadline,
		Backlog:     backlog,
		Done:        false,
	}

	t.Tasks = append(t.Tasks, task)
	return t
}

func (t *TaskManager) Delete(id int) *TaskManager {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			break
		}
	}
	return t
}

func (t *TaskManager) Update(id int, name string, description string, cron string, deadline time.Time, backlog bool, done bool) *TaskManager {
	for i, task := range t.Tasks {
		if task.ID == id {
			// Update the task's properties
			t.Tasks[i].Name = name
			t.Tasks[i].Description = description
			t.Tasks[i].Cron = cron
			t.Tasks[i].Deadline = deadline
			t.Tasks[i].Backlog = backlog
			t.Tasks[i].Done = done
			break
		}
	}
	return t
}

func (t *TaskManager) nextId() int {
	maxId := 0

	// Find biggest ID from the existing Tasks
	for _, task := range t.Tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	return maxId + 1
}
