# 📝 Task Management API Documentation

Base URL: `http://localhost:8080`

---

## 📌 Endpoints

### 1. GET `/tasks`

- **Description:** Fetch all tasks.
- **Response Example:**

```json
[
  {
    "id": 1,
    "title": "Sample Task",
    "description": "This is a test",
    "due_date": "2025-04-30",
    "status": "pending"
  }
]
```
