# user-svc

## Prerequisite
- Go Version ``>= go 1.21.5``
- PostgreSQL Version ``>= version 16.0``

## How To Run
```
1. git clone https://github.com/luthfirahman/user-service.git
2. cd user-service
3. cp .env.example .env
4. configure .env with your postgres
DB_HOST = localhost
DB_USER = postgres
DB_PASS = 
DB_NAME = 
DB_PORT = 5432
5. Open terminal, follow the steps below:
- if you haven't downloaded postgres, you can download it first
- Run -> psql -U postgres
- Run -> Create database according to what you put in .env
- Run -> Exit
6. go run cmd/main.go
```
