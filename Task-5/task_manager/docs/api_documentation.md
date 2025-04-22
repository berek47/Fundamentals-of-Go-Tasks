📝 Task Management API Documentation
Base URL: http://localhost:8080
Postman Documentation: [Click here to view the Postman collection](https://documenter.getpostman.com/view/18965273/2sB2iwGEyT)
📌 Endpoints
1. GET /tasks
Description: Fetch all tasks.
URL: GET http://localhost:8080/tasks
Response Example:
```json
[
  {
    "id": 1,
    "title": "Sample Task",
    "description": "This is a test",
    "due_date": "2025-04-30",
    "status": "pending"
  }
]```
2. GET /tasks/{id}
Description: Fetch a single task by ID.
URL: GET http://localhost:8080/tasks/{id}
URL Parameters:
- id: ID of the task to retrieve.
Response Example:
```json
{
  "id": 1,
  "title": "Sample Task",
  "description": "This is a test",
  "due_date": "2025-04-30",
  "status": "pending"
}```
3. POST /tasks
Description: Create a new task.
URL: POST http://localhost:8080/tasks
Request Body Example:
```json
{
  "title": "New Task",
  "description": "This is a newly created task",
  "due_date": "2025-05-01",
  "status": "pending"
}```
Response Example:
```json
{
  "id": 2,
  "title": "New Task",
  "description": "This is a newly created task",
  "due_date": "2025-05-01",
  "status": "pending"
}```
4. PUT /tasks/{id}
Description: Update a task completely (replace all fields).
URL: PUT http://localhost:8080/tasks/{id}
URL Parameters:
- id: ID of the task to update.
Request Body Example:
```json
{
  "title": "Updated Task Title",
  "description": "Updated task description",
  "due_date": "2025-06-01",
  "status": "in progress"
}```
Response Example:
```json
{
  "id": 1,
  "title": "Updated Task Title",
  "description": "Updated task description",
  "due_date": "2025-06-01",
  "status": "in progress"
}```
5. PATCH /tasks/{id}
Description: Partially update a task (e.g., update only the status).
URL: PATCH http://localhost:8080/tasks/{id}
URL Parameters:
- id: ID of the task to update.
Request Body Example:
```json
{
  "status": "completed"
}```
Response Example:
```json
{
  "id": 1,
  "title": "Sample Task",
  "description": "This is a test",
  "due_date": "2025-04-30",
  "status": "completed"
}```
6. DELETE /tasks/{id}
Description: Delete a task by ID.
URL: DELETE http://localhost:8080/tasks/{id}
URL Parameters:
- id: ID of the task to delete.
Response Example:
```json
{
  "message": "Task with id 1 has been deleted."
}```
📌 Example Error Responses
1. Task Not Found
Error Code: 404
Error Message Example:
```json
{
  "error": "Task with ID 999 not found."
}```
2. Invalid Request Body
Error Code: 400
Error Message Example:
```json
{
  "error": "Request body is missing required fields."
}```
3. Invalid Task Status
Error Code: 400
Error Message Example:
```json
{
  "error": "Invalid status value. Accepted values are 'pending', 'in progress', 'completed'."
}```
📌 Example Success Responses
1. Task Created Successfully
Success Code: 201
Success Message Example:
```json
{
  "id": 2,
  "title": "New Task",
  "description": "This is a newly created task",
  "due_date": "2025-05-01",
  "status": "pending"
}```
2. Task Updated Successfully
Success Code: 200
Success Message Example:
```json
{
  "id": 1,
  "title": "Updated Task Title",
  "description": "Updated task description",
  "due_date": "2025-06-01",
  "status": "in progress"
}```
3. Task Deleted Successfully
Success Code: 200
Success Message Example:
```json
{
  "message": "Task with id 1 has been deleted."
}```