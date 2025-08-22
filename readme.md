Task Tracker
# How to run
Clone the repository and run the following command:
```
git clone https://github.com/Summer8918/task_tracker.git

```

Build a binary and run the program:
```
go build -o task_tracker .
./task_tracker
```

# How to run test
Go to task_tracker directory, then run the following command, which run tests for all packages under the current folder, recursively.
```
go test ./...
```

# How to use
### Adding a new task
task_tracker add "Buy groceries"

### Updating and deleting tasks
task_tracker update 1 "Buy groceries and cook dinner"  
task_tracker delete 1

### Marking a task as in progress, todo, done or blocked
task_tracker mark-in-progress 1  
task_tracker mark-done 1  

### Listing all tasks
task_tracker list

### Listing tasks by status
task_tracker list done  
task_tracker list todo  
task_tracker list in-progress  
task_tracker list blocked  