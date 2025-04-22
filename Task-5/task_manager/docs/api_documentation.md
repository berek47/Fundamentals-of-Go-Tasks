📝 Task Management API Documentation (With MongoDB Integration)
Base URL:
http://localhost:8080

Postman Documentation:
Click here to view the Postman collection

📌 Endpoints

1. GET /tasks
   Description: Fetch all tasks from MongoDB.

URL:
GET http://localhost:8080/tasks
Response Example:
[
{
"id": "1",
"title": "Sample Task",
"description": "This is a test",
"due_date": "2025-04-30",
"status": "pending"
}
] 2. GET /tasks/{id}
Description: Fetch a single task by ID from MongoDB.

URL:
GET http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to retrieve.

Response Example:
{
"id": "1",
"title": "Sample Task",
"description": "This is a test",
"due_date": "2025-04-30",
"status": "pending"
}

3. POST /tasks
   Description: Create a new task in MongoDB.

URL:
POST http://localhost:8080/tasks
Request Body Example:
{
"title": "New Task",
"description": "This is a newly created task",
"due_date": "2025-05-01",
"status": "pending"
}
Response Example:
{
"id": "2",
"title": "New Task",
"description": "This is a newly created task",
"due_date": "2025-05-01",
"status": "pending"
} 4. PUT /tasks/{id}
Description: Update an existing task completely (replace all fields).

URL:
PUT http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to update.

Request Body Example:

{
"title": "Updated Task Title",
"description": "Updated task description",
"due_date": "2025-06-01",
"status": "in progress"
}
Response Example:
{
"id": "1",
"title": "Updated Task Title",
"description": "Updated task description",
"due_date": "2025-06-01",
"status": "in progress"
}

5. PATCH /tasks/{id}
   Description: Partially update a task (e.g., update only the status).

URL:
PATCH http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to update.

Request Body Example:
{
"status": "completed"
}

Response Example:
{
"id": "1",
"title": "Sample Task",
"description": "This is a test",
"due_date": "2025-04-30",
"status": "completed"
}

6. DELETE /tasks/{id}
   Description: Delete a task by ID from MongoDB.

URL:

DELETE http://localhost:8080/tasks/{id}
URL Parameters:

id: ID of the task to delete.

Response Example:

{
"message": "Task with id 1 has been deleted."
}

📌 Example Error Responses

1. Task Not Found
   Error Code: 404

Error Message Example:

{
"error": "Task with ID 999 not found."
}

2. Invalid Request Body
   Error Code: 400

Error Message Example:

{
"error": "Request body is missing required fields."
}

3. Invalid Task Status
   Error Code: 400

Error Message Example:

{
"error": "Invalid status value. Accepted values are 'pending', 'in progress', 'completed'."
}

📌 Example Success Responses

1. Task Created Successfully
   Success Code: 201

Success Message Example:

{
"id": "2",
"title": "New Task",
"description": "This is a newly created task",
"due_date": "2025-05-01",
"status": "pending"
}

2. Task Updated Successfully
   Success Code: 200

Success Message Example:

{
"id": "1",
"title": "Updated Task Title",
"description": "Updated task description",
"due_date": "2025-06-01",
"status": "in progress"
}

3. Task Deleted Successfully
   Success Code: 200

Success Message Example:

3. Task Deleted Successfully
   Success Code: 200

Success Message Example:

{
"message": "Task with id 1 has been deleted."
}

📌 MongoDB Integration Documentation

1. MongoDB Setup
   Step 1: Set up a MongoDB instance (either locally or via a cloud service like MongoDB Atlas).

Step 2: Install the MongoDB Go Driver package.
go get go.mongodb.org/mongo-driver

Step 3: Configure your Go application to connect to MongoDB by setting the connection URI and database name.
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, err := mongo.Connect(context.Background(), clientOptions)
if err != nil {
log.Fatal(err)
}

2. CRUD Operations Using MongoDB
   Create a Task:
   taskCollection := client.Database("taskDB").Collection("tasks")
   newTask := Task{Title: "New Task", Description: "A new task", DueDate: "2025-05-01", Status: "pending"}
   result, err := taskCollection.InsertOne(context.Background(), newTask)

Retrieve All Tasks:
cursor, err := taskCollection.Find(context.Background(), bson.D{{}})

Update Task:

filter := bson.D{{"id", taskID}}
update := bson.D{{"$set", bson.D{{"status", "completed"}}}}
\_, err := taskCollection.UpdateOne(context.Background(), filter, update)

Delete Task:
filter := bson.D{{"id", taskID}}
\_, err := taskCollection.DeleteOne(context.Background(), filter)

3. Error Handling for MongoDB
   Ensure proper error handling for all MongoDB operations, including:

Connection errors: when unable to connect to the database.

Query errors: when the database query is invalid.

Data validation errors: when fields are missing or incorrect.

4. Verifying MongoDB Data
   Use MongoDB Compass or a query tool to verify that tasks are correctly stored in the database after performing CRUD operations.

📌 Folder Structure
task_manager/
├── main.go
├── controllers/
│ └── task_controller.go
├── models/
│ └── task.go
├── data/
│ └── task_service.go
├── router/
│ └── router.go
├── docs/
│ └── api_documentation.md
└── go.mod

main.go: Entry point of the application.

controllers/task_controller.go: Handles incoming HTTP requests and invokes the appropriate service methods.

models/task.go: Defines the data structures used in the application.

data/task_service.go: Contains business logic and data manipulation functions. MongoDB interaction is implemented here.

router/router.go: Sets up routes and initializes the Gin router.

docs/api_documentation.md: Contains the updated API documentation with MongoDB integration instructions.

go.mod: Defines the module and its dependencies.

Evaluation Criteria:
Successful MongoDB integration.

Correct implementation of CRUD operations with MongoDB.

Proper error handling and validation for MongoDB operations.

Testing API endpoints and verifying data correctness.

Clear and comprehensive documentation of the MongoDB integration process.
