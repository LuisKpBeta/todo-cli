# Table of Contents

* [DevChallenge](#devchallenge) 
* [Challenge](#challenge)
* [Requirements](#requirements)
* [Build](#build)


# DevChallenge
<a href="https://devchallenge.now.sh/"> DevChallenge</a> allows you to evolve your skills as a programmer! Join our <a href="https://discord.gg/yvYXhGj">community</a> o/

# Challenge
Your challenge is to create a command-line interface that allows you to maintain a task list with pending and done activities!



# Requirements:
The following requirements should be implemented in a command-line interface. This CLI should be controlled by using commands (including subcommands and arguments), a menu with options or other creative methods.
1. **[optional]** If you decide to use commands keep in mind that they have to be an explicit action, while arguments use a dash (`-a` or `--argument`) and are responsible for modifying a command's behavior. Example:

```bash
task <subcommand> # Accepts add, complete, delete, list and next as subcommands

task add <description> [-p <priority>] # Adds a pending task. Can set the task's priority to low, normal or high with the -p (or --priority) option

task complete <id> # Marks a task as done
task delete <id> # Deletes a task
task list [-a] # Shows pending tasks. The -a (or --all) option shows all tasks
task next # Shows the next task of each priority
```
2. **[add task]** : It should be possible to register a new task.

    A task must have a unique id, a description, a creation date, a status (shows if a task is pending or done) and a priority (high, normal or low). Example:
    
    ```javascript
    { id: 1, description: 'Buy 6 eggs', created: 2021-04-01T20:54:19.410Z, status: 'pending', priority: 'high' };
    ```
3. **[mark task as done]** : Should be able to update a task's status to done.
4. **[delete task]** : Should be able to delete a task by providing its corresponding id.
5. **[list tasks]** : Should be able to list the tasks with status that isn't done.
    
    Instead of showing the creation date of each task, that property should be substituted by a new property that displays how long ago the task was created (1 month). Example:
    ```javascript
    [{ id: 1, description: 'Buy 6 eggs', age: 22 hours, status: 'pending', priority: 'high' }]
    ```
6. **[list all tasks]** : Should be able to list all tasks, including those with status done.
    
    Instead of showing the creation date of each task, that property should be substituted by a new property that displays how long ago the task was created (1 month).
7. **[list next tasks]** : Should be able to list a task from each priority. In other words, a high priority task, a normal priority task and a low priority task, if they exist. The listed task of each priority with be the oldest from its group.

    Instead of showing the creation date of each task, that property should be substituted by a new property that displays how long ago the task was created (1 month).
8. **[local file or database]** : Should be able to persist data so it isn't lost after the CLI finishes execution.

# Build

To build the application for the linux environment, you can use the standard command
``` 
go build cmd/todo.go
```
To build for windows OS, you need to input some flags for correct output
```
GOOS=windows GOARCH=386 \
CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ \
CC=i686-w64-mingw32-gcc \
go build cmd/todo.go 
```

If you are using the project inside the devcontainer, the gcc dependencies are already installed

To use builded app, you only need to add aplication to PATH env of your OS