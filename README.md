# User Service

This is a User Service app built with Go programming language, Docker for containerization, and PostgreSQL for data storage.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Environment Variables](#environment-variables)

## Getting Started

### Prerequisites

Before running the User Service app, ensure you have the following installed:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (if you want to run the application locally without Docker)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/luthfirahman/user-svc.git
```

2. Navigate to the project directory:

```bash
cd user-svc
```

3. Set up environment variables (see Environment Variables).

4. Build and run the Docker containers:

```bash
docker-compose up
```

## Usage

Once the Docker containers are up and running, you can access the API at http://localhost:8080.

If you want to run the application locally without Docker:

1. Make sure PostgreSQL is running on your machine and accessible at localhost:5432.

2. Set up the required environment variables (see [Environment Variables](#environment-variables)).

3. Build and run the application:

```bash
go run cmd/main.go
```

## Endpoints

The following endpoints are available:

- POST /register: Register a new user.
- POST /login: Authenticate user.
- GET /me: Get user profile.
- PATCH /: Update user profile.

## Environment Variables

Ensure these environment variables are set:

- DB_HOST: PostgreSQL host address.
- DB_PORT: PostgreSQL port (default is 5432).
- DB_NAME: PostgreSQL database name.
- DB_USER: PostgreSQL username.
- DB_PASS: PostgreSQL password.
