# Task Management CLI Application

A command-line interface (CLI) application written in Go for managing tasks. This application allows users to create, update, delete, and track the status of tasks with persistent storage in JSON format.

## Features

- Add new tasks
- Update existing task descriptions
- Mark tasks as done or in-progress
- Delete tasks
- List all tasks
- Filter tasks by status (done, in-progress, pending)
- Persistent storage using JSON file
- Automatic timestamp tracking for task creation and updates

## Installation

1. Ensure you have Go installed on your system
2. Clone this repository
3. Navigate to the project directory
4. Build the application:
```bash
go build
```

## Usage

The application supports various command-line flags for different operations:

### Adding a Task
```bash
./task-manager -add "Complete project documentation"
```

### Updating a Task
```bash
./task-manager -update-id 1 -update-description "Updated task description"
```

### Marking Task Status
```bash
# Mark as done
./task-manager -mark-done 1

# Mark as in-progress
./task-manager -mark-in-progress 1
```

### Deleting a Task
```bash
./task-manager -delete 1
```

### Listing Tasks
```bash
# List all tasks
./task-manager -list

# List done tasks
./task-manager -list-done

# List in-progress tasks
./task-manager -list-in-progress

# List pending tasks
./task-manager -list-pending
```

## Data Structure

Tasks are stored in a JSON file with the following structure:

```json
{
  "Tasks": [
    {
      "Id": 1,
      "Description": "Task description",
      "Status": "pending",
      "CreatedAt": "2025-01-16T10:00:00Z",
      "UpdateAt": "2025-01-16T10:00:00Z"
    }
  ]
}
```

Task status can be one of three values:
- `pending`: Default status for new tasks
- `in-progress`: Tasks that are currently being worked on
- `done`: Completed tasks

## Error Handling

The application includes error handling for common scenarios:
- File operations (reading/writing JSON file)
- Invalid task IDs
- JSON marshaling/unmarshaling
- Missing required flags

## Dependencies

The application uses only Go standard library packages:
- `encoding/json`: For JSON operations
- `flag`: For command-line flag parsing
- `fmt`: For formatted I/O
- `log`: For error logging
- `os`: For file operations
- `time`: For timestamp management

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## Project Solution for:

https://roadmap.sh/projects/task-tracker

## License

This project is licensed under the MIT License - see the LICENSE file for details.
