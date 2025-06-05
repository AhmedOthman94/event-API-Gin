# 📦 Event API with Gin (Go)

A powerful, secure RESTful API built using the **Gin** framework in **Go**, designed for managing user accounts and event data. Supports authentication, JWT protection, SQLite database, and clear modular structure.

![Gin Logo](https://raw.githubusercontent.com/gin-gonic/logo/master/color.png)

---

## 🚀 Features

✅ User registration and login with password encryption (bcrypt)  
🔐 JWT-based authentication and middleware protection  
📅 Create and fetch event data associated with each user  
🧱 SQLite database with schema auto-creation  
📂 Clean folder structure for maintainability  
🛡 Protected endpoints using custom middleware  
📄 Ready to deploy and extend

---

## 🗂 Project Structure

```
.
├── db/             # DB connection and schema creation
├── models/         # Structs and DB operations
├── routes/         # Public & protected route handlers
├── utils/          # Helper functions (JWT, hashing)
├── main.go         # Entry point
├── go.mod          # Go module
├── .gitignore
└── README.md
```

---

## 📡 API Endpoints

| Method | Endpoint    | Description             | Auth Required |
|--------|-------------|-------------------------|----------------|
| POST   | `/register` | Register a new user     | ❌             |
| POST   | `/login`    | Login and get token     | ❌             |
| GET    | `/events`   | Fetch all events        | ❌             |
| POST   | `/events`   | Create new event        | ✅             |

> ✅ Use token in `Authorization: Bearer <your_token>` for protected routes.

---

## 🛠️ Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/event-API-Gin.git
cd event-API-Gin
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Run the server
```bash
go run main.go
```

The server will be running on: `http://localhost:8080`

---

## 🔐 Example Usage

### 📥 Register

```json
POST /register
{
  "email": "user@example.com",
  "password": "securepass123"
}
```

### 🔑 Login

```json
POST /login
{
  "email": "user@example.com",
  "password": "securepass123"
}
```

You will receive a JWT token to access protected routes.

### ➕ Create Event

```json
POST /events
Authorization: Bearer <your_token>

{
  "name": "Golang Meetup",
  "description": "Learn Go & Network",
  "location": "Cairo",
  "date_time": "2025-06-05T18:30:00Z"
}
```

---

## 🧪 Testing

Use [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) `.http` files or Postman.

---

## 👨‍💻 Author

Made with ❤️ by [Ahmed Othman](https://github.com/AhmedOthman94)

---

## 📃 License

MIT License © 2025