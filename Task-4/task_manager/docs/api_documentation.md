📝 Task Management API Documentation (With MongoDB Integration)
Base URL
http://localhost:8080

Postman Documentation
Click here to view the Postman collection [Replace with actual Postman link]

📌 Endpoints

1. GET /tasks
   Description: Fetch all tasks from MongoDB
   URL: GET http://localhost:8080/tasks
   Response Example:

json
[
{
"id": "65f1c45a9e7d8b2a1d4e3f5a",
"title": "Sample Task",
"description": "This is a test",
"due_date": "2025-04-30",
"status": "pending"
}
] 2. GET /tasks/{id}
Description: Fetch a single task by ID from MongoDB
URL: GET http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to retrieve

Response Example:

json
{
"id": "65f1c45a9e7d8b2a1d4e3f5a",
"title": "Sample Task",
"description": "This is a test",
"due_date": "2025-04-30",
"status": "pending"
} 3. POST /tasks
Description: Create a new task in MongoDB
URL: POST http://localhost:8080/tasks
Request Body Example:

json
{
"title": "New Task",
"description": "This is a newly created task",
"due_date": "2025-05-01",
"status": "pending"
}
Response Example:

json
{
"id": "65f1c45a9e7d8b2a1d4e3f5b",
"title": "New Task",
"description": "This is a newly created task",
"due_date": "2025-05-01",
"status": "pending"
} 4. PUT /tasks/{id}
Description: Update an existing task completely (replace all fields)
URL: PUT http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to update

Request Body Example:

json
{
"title": "Updated Task Title",
"description": "Updated task description",
"due_date": "2025-06-01",
"status": "in progress"
}
Response Example:

json
{
"id": "65f1c45a9e7d8b2a1d4e3f5a",
"title": "Updated Task Title",
"description": "Updated task description",
"due_date": "2025-06-01",
"status": "in progress"
} 5. PATCH /tasks/{id}
Description: Partially update a task (e.g., update only the status)
URL: PATCH http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to update

Request Body Example:

json
{
"status": "completed"
}
Response Example:

json
{
"id": "65f1c45a9e7d8b2a1d4e3f5a",
"title": "Sample Task",
"description": "This is a test",
"due_date": "2025-04-30",
"status": "completed"
} 6. DELETE /tasks/{id}
Description: Delete a task by ID from MongoDB
URL: DELETE http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to delete

Response Example:

json
{
"message": "Task with id 65f1c45a9e7d8b2a1d4e3f5a has been deleted."
}
📌 Example Error Responses

1. Task Not Found
   Error Code: 404
   Example:

json
{
"error": "Task with ID 65f1c45a9e7d8b2a1d4e3f99 not found."
} 2. Invalid Request Body
Error Code: 400
Example:

json
{
"error": "Request body is missing required fields."
} 3. Invalid Task Status
Error Code: 400
Example:

json
{
"error": "Invalid status value. Accepted values are 'pending', 'in progress', 'completed'."
}
📌 MongoDB Integration Documentation

1. MongoDB Setup
   Set up MongoDB: Install locally or use MongoDB Atlas

Install Go Driver:

bash
go get go.mongodb.org/mongo-driver
Connection Code:

go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, err := mongo.Connect(context.Background(), clientOptions) 2. CRUD Operations
Create Task:

go
taskCollection := client.Database("taskDB").Collection("tasks")
result, err := taskCollection.InsertOne(ctx, task)
Get All Tasks:

go
cursor, err := taskCollection.Find(ctx, bson.D{})
Update Task:

go
filter := bson.D{{"_id", taskID}}
update := bson.D{{"$set", updateData}}
\_, err := taskCollection.UpdateOne(ctx, filter, update)
Delete Task:

go
\_, err := taskCollection.DeleteOne(ctx, bson.D{{"_id", taskID}}) 3. Data Verification
Use MongoDB Compass to verify data integrity after operations.

📌 Example Success Responses

1. Task Created (201)
   json
   {
   "id": "65f1c45a9e7d8b2a1d4e3f5b",
   "title": "New Task",
   "description": "This is a newly created task",
   "due_date": "2025-05-01",
   "status": "pending"
   }
2. Task Updated (200)
   json
   {
   "id": "65f1c45a9e7d8b2a1d4e3f5a",
   "title": "Updated Task Title",
   "description": "Updated task description",
   "due_date": "2025-06-01",
   "status": "in progress"
   }
3. Task Deleted (200)
   json
   {
   "message": "Task with id 65f1c45a9e7d8b2a1d4e3f5a has been deleted."
   }
