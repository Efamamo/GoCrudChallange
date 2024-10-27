# Go-crud-challenge

## Implement Simple CRUD API

Your task is to implement simple CRUD API using in-memory database underneath.

## Details:

1. The task must be solved using only Go, or Go related framework
2. API path `/person`:
   - **GET** `/person` or `/person/${personId}` should return all persons or person with corresponding `personId`
   - **POST** `/person` is used to create record about new person and store it in database
   - **PUT** `/person/${personId}` is used to update record about existing person
   - **DELETE** `/person/${personId}` is used to delete record about existing person from database
3. Persons are stored as `objects` that have following properties:
   - `id` — unique identifier (`string`, `uuid`) generated on server side
   - `name` — person's name (`string`, **required**)
   - `age` — person's age (`number`, **required**)
   - `hobbies` — person's hobbies (`array` of `strings` or empty `array`, **required**)
4. Requests to non-existing endpoints (e.g. `/some-non/existing/resource`) should be handled.
5. Internal server errors should be handled and processed correctly.
6. Make sure the api is accesible by frontend apps hosted on a different domain (cross-site resource sharing)

# Go CRUD Challenge

## Overview

This project is a CRUD (Create, Read, Update, Delete) application built with Go, implementing a Command Query Responsibility Segregation (CQRS) pattern. The application manages person entities and provides a simple interface for performing CRUD operations.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (version 1.16 or later)
- **Make** (for using the Makefile)
- **Docker** (optional, if using Docker for dependencies)

## Getting Started

### Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Efamamo/GoCrudChallange.git
   cd GoCrudChallange
   ```

2. **Install dependencies:**

   If you are using Go modules, run the following command to install dependencies:

   ```bash
   go mod tidy
   ```

### Using the Makefile

The project includes a `Makefile` for building, running, and testing the application. Below are the available commands.

#### Build

To build the application, run:

```bash
make build
```

This will compile the Go code and output an executable named `main` in the `bin` directory.

#### Run

To run the application, use the following command:

```bash
make run
```

This will execute the compiled binary.

#### Run in Development Mode

To run the application in development mode with live reloading, use:

```bash
make run-dev
```

This requires the `air` tool for live reloading. If you haven't installed it yet, you can do so with:

```bash
go install github.com/cosmtrek/air@latest
```

#### Testing

To run tests, execute:

```bash
make test -cov
```

This will compile and run the tests in the `test` directory. You can check the output for results and coverage information.
