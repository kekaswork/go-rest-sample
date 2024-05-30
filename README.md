# Go REST Sample

A small REST API application created as part of a test assessment.

## Overview

This application demonstrates a simple REST API using Golang and PostgreSQL. It includes basic operations to manage students, subjects, and marks.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [Golang](https://golang.org/dl/)

## Getting Started

### Step 1: Start the PostgreSQL Database

Run the following command to start the PostgreSQL database using Docker Compose:

```sh
docker-compose up
```

### Step 2: Run the Application

Use the following command to build and run the Go application:

```sh
make run
```

The make run command is a shortcut for:

```sh
go run cmd/go-rest-sample/main.go
```


## Initial Database Setup

On the first run, the database will be initialized and some sample data will be added automatically.

## API Endpoints

Here are some of the API endpoints you can interact with:


### Students

- `GET /students` - Retrieve all students
- `POST /students` - Add a new student
- `PUT /students/:id` - Update an existing student by ID
- `DELETE /students/:id` - Delete a student by ID


### Subjects
- `GET /subjects` - Retrieve all subjects
- `POST /subjects` - Add a new subject
- `PUT /subjects/:id` - Update an existing subject by ID
- `DELETE /subjects/:id` - Delete a subject by ID


### Marks
- `GET /marks` - Retrieve all marks
- `POST /marks` - Add a new mark
- `PUT /marks/:id` - Update an existing mark by ID
- `DELETE /marks/:id` - Delete a mark by ID

### Reports

- GET `/report` - Generate a report based on query parameters

Example Usage
To get students with a last name starting with a specific prefix and an average mark above a certain value for a specific subject:

```
GET /report?prefix=B&subject_id=1&min_avg_mark=2.0
```


## Configuration

The application uses environment variables for configuration. Create a .env file in the root directory with the following content:

```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
```