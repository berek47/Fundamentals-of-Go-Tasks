# Task Management API Documentation

## MongoDB Integration

- Mongo URI: `mongodb://localhost:27017`
- Database: `taskdb`
- Collection: `tasks`

## Endpoints

### POST /tasks

Create a new task.

- Body: `{ "title": "...", "description": "...", "completed": false }`

### GET /tasks

Retrieve all tasks.

### GET /tasks/:id

Retrieve task by ID.

### PUT /tasks/:id

Update task.

### DELETE /tasks/:id

Delete task.

## Notes

- Uses Mongo Go Driver (`go.mongodb.org/mongo-driver`)
- All data is now persisted in MongoDB
