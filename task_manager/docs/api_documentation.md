# Task Management API

This API allows users to manage tasks with basic CRUD operations. The API is built using Go and the Gin framework.

## Base URL

The base URL for all API endpoints is `http://localhost:8080`.

## Endpoints

- **GET /tasks**: Retrieve a list of all tasks.
- **GET /tasks/:id**: Retrieve details of a specific task by ID.
- **POST /tasks**: Create a new task.
- **PUT /tasks/:id**: Update an existing task by ID.
- **DELETE /tasks/:id**: Delete a task by ID.

## Usage

Ensure the server is running on the specified port before making requests.

## Status Codes

- `200 OK`: Request successful.
- `201 Created`: Resource successfully created.
- `400 Bad Request`: Invalid request parameters.
