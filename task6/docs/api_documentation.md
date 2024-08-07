# Task Management API

## Overview

The Task Management API is a RESTful service built using the Go programming language and the Gin framework. This API provides endpoints for user registration, login, and task management functionalities, including JWT-based authentication and authorization. The API utilizes MongoDB for data storage, allowing registered users to perform CRUD operations on tasks.

## Features

- **User Registration and Login:** Supports user registration and login with JWT authentication, storing tokens in HTTP-only cookies.
- **CRUD Operations on Tasks:** Users can create, read, update, and delete tasks.
- **Secure Access:** Protected routes ensure that only authenticated users can access task operations.
- **MongoDB Integration:** Tasks and user information are stored in MongoDB.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
  - [User Registration](#user-registration)
  - [User Login](#user-login)
- [Project Structure](#project-structure)

## Installation

### Prerequisites

Ensure you have the following installed:

- **Go**: [Download and install Go](https://golang.org/doc/install)
- **MongoDB**: [Download and install MongoDB](https://docs.mongodb.com/manual/installation/)
- **Git**: [Download and install Git](https://git-scm.com/downloads)

### Clone the Repository

```bash
git clone https://github.com/yourusername/task_manager.git
cd task_manager
```

## Install Dependencies  

`go mod tidy`

## Configuration

### Environment Variables

Create a .env file in the root of your project directory to configure the MongoDB connection string and JWT secret key:

``` MONGO_URI=mongodb://localhost:27017
DB_NAME=task_manager
JWT_SECRET=your_jwt_secret_key
```

## MongoDB Setup

Ensure MongoDB is running on your local  machine or update the `MONGO_URI` in the .env file to connect to a remote MongoDB instance.

## Running the Application

Start the Server:

` go run main.go `

The server will start on <http://localhost:8080>

## API Documentation

### User Registration

-Endpoint: /register
-Method: POST
-Description: Registers a new user.
Request:

- Content-Type: `application/json`

``` bash
{
  "username": "string (3-20 characters, required)",
  "password": "string (6+ characters, required)"
}
```

Response:
Status Code: `201 Created`

``` bash
{
  "message": "User registered successfully"
}
```

Errors:

- Status Code: `400 Bad Request`

``` bash
{
  "error": "Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag"
}
```

- Status Code: `409 Conflict`

``` bash
{
  "error": "Username already taken"
}
```

## Project Structure

Here's an overview of the project's directory structure:

``` text
task_manager/
|-- controllers/
|   |-- authController.go
|   |-- taskController.go
|-- models/
|   |-- user.go
|   |-- task.go
|-- routes/
|   |-- authRoutes.go
|   |-- taskRoutes.go
|-- utils/
|   |-- jwt.go
|   |-- validator.go
|-- .env
|-- main.go
|-- go.mod
|-- README.md
```
