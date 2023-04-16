# GODO-CLI

GODO-CLI is a command-line application for managing and scheduling tasks. The name is derived from the programming language `Go` and `To Do`, thus resulting in `Go Do`.

## Features

- [x] Manage tasks with simple CLI commands

### Roadmap

- [ ] Schedule tasks using cron expressions
- [ ] Receive shell notifications when tasks are due
- [ ] Block websites if tasks are not completed by their deadline
- [ ] Snooze tasks for a specified period


## Usage

### Add a task

To add a new task, use the `task add` command:

```
godo task add NAME DESCRIPTION CRON DEADLINE IN_BACKLOG

godo task add "New task name" "New task description" "0 0 * * *" "2023-04-16 15:04:05" "true"
```

### List tasks

To list all active tasks, use the `list` command:

```
godo list
```

To show all tasks, including completed ones, use the `list-all` command:

```
godo list-all
```

To get a specific task use "get" command:

```
godo get 1
```


### Update a task

To update an existing task, use the `update` command with the task ID:

```
godo update ID NAME DESCRIPTION CRON DEADLINE IN_BACKLOG

godo update 1 "New task name" "New task description" "0 0 * * *" "2023-04-16 15:04:05" "true"
```

### Delete a task

To delete a task, use the `delete` command with the task ID:

```
godo delete 1
```

### Mark a task as done

To mark a task as done, use the `done` command with the task ID:

```
godo done 1
```

## License

This project is licensed under the MIT License
