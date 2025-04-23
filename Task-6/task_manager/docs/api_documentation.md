🛡️ Task Management API – Authentication & Authorization with JWT

This API adds secure user authentication and authorization to a task management backend using JWT (JSON Web Tokens). It allows users to register, log in, and access protected routes based on user roles (admin, user).

🎯 Objective

Enhance the Task Management API by:

    Creating user accounts with roles

    Implementing JWT-based authentication

    Protecting routes with middleware

    Restricting access based on roles

    Storing credentials securely

    Documenting endpoints and access control

📁 Folder Structure

task_manager/
├── main.go → Application entry point
├── go.mod → Go module and dependencies

├── controllers/
│ └── controller.go → Handles HTTP routes for users and tasks

├── models/
│ ├── task.go → Task struct definition
│ └── user.go → User struct definition

├── data/
│ ├── task_service.go → Business logic for tasks
│ └── user_service.go → Business logic for users + password hashing

├── middleware/
│ └── auth_middleware.go → JWT validation & role-based access control

├── router/
│ └── router.go → API route setup with middleware

├── docs/
│ └── api_documentation.md → API documentation

🌐 Base URL

http://localhost:8080
🔐 Authentication Endpoints

POST /register
Registers a new user.

Request:

{
"username": "johndoe",
"password": "securepassword",
"role": "admin"
}

Response (201 Created):

{
"user": {
"id": 1,
"username": "johndoe",
"role": "admin"
}
}

Errors:

    400 Bad Request – Username already exists or input is invalid

POST /login
Logs in a user and returns a signed JWT.

Request:

{
"username": "johndoe",
"password": "securepassword"
}

Response (200 OK):

{
"token": "eyJhbGciOiJIUzI1NiIsInR..."
}

Errors:

    401 Unauthorized – Invalid credentials

🛡️ JWT Authorization

All protected endpoints require a valid JWT token in the Authorization header.

Format:

Authorization: <JWT Token>

Example:

Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

🔒 Protected Endpoints

GET /api/tasks
Accessible by any authenticated user.

Headers:

Authorization: <JWT Token>

Response:

{
"message": "Get tasks - authenticated"
}

GET /api/admin/dashboard
Accessible only by admin users.

Headers:

Authorization: <Admin JWT Token>

Response:

{
"message": "Admin dashboard"
}

Errors:

    401 Unauthorized – Missing or invalid token

    403 Forbidden – Valid token, not admin

🔧 JWT Claims

Tokens include:

    username

    role (admin or user)

    exp (expiration timestamp)

🔐 Security Features

    Passwords hashed using bcrypt

    JWTs signed with HMAC SHA256

    Token validated through middleware

    Role-based access control handled with middleware

🧪 Postman Testing Guide
Step 1: Register a User

    POST http://localhost:8080/register

{
"username": "johndoe",
"password": "securepassword",
"role": "admin"
}

Step 2: Login

    POST http://localhost:8080/login

{
"username": "johndoe",
"password": "securepassword"
}

Copy the token from the response.
Step 3: Call Protected Route

    GET http://localhost:8080/api/tasks

    Header: Authorization: <JWT Token>

Step 4: Call Admin Route

    GET http://localhost:8080/api/admin/dashboard

    Header: Authorization: <JWT Token>

Postman Script to Auto-Save Token:

Add this to the Tests tab in the /login request:

pm.environment.set("jwt", pm.response.json().token);

Use in headers:

Authorization: {{jwt}}

❌ Negative Test Cases
Scenario Expected Response
No token 401 Unauthorized
Invalid or expired token 401 Unauthorized
Regular user on admin route 403 Forbidden
Duplicate registration 400 Bad Request
✅ Evaluation Checklist
Requirement Status
User registration/login ✅
Password hashing with bcrypt ✅
JWT generation & claim management ✅
Authentication middleware ✅
Role-based authorization ✅
Protected routes ✅
Admin-only route ✅
Postman test flow ✅
Project folder structure and documentation ✅
