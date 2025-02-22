# Go Todo List API

A RESTful API for managing todos built with Go and Gin framework. This API provides endpoints for creating, reading, updating, and deleting todo items.

## Tech Stack
- Go
- Gin Web Framework
- CORS support
- PostgreSQL

## Features
- Create new todos
- Retrieve all todos
- Update existing todos
- Delete todos
- CORS enabled for cross-origin requests

## PROJECT STRUCTURE 

todo-list/
├── controllers/
│   └── todo.go
├── database/
│   └── database.go
├── models/
│   └── todo.go
├── repository/
│   └── todo.go
├── routes/
│   └── route.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md

## API Endpoints

### Get All Todos
- **URL**: `/api/todos`
- **Method**: `GET`

### Create Todo
- **URL**: `/api/todos`
- **Method**: `POST`
- **Request Body**:
json
{
    "taskName": "Sample todo",
    "taskDescription": "Todo description",
    "dueDate": false
}

### Edit Todo
- **URL**: `/api/todo/:id`
- **Method**: `PUT`
- **Request Body**:
json
{
    "taskName": "Sample todo",
    "taskDescription": "Todo description",
    "dueDate": false
}

### Delete Todo
- **URL**: `/api/todo/:id`
- **Method**: `DELETE`

```