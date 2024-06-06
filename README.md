# Go CRUD Rest API Demo App

This is a simple CRUD application for employees data. Written in Golang with in-memory data store (database).


### Tech Stack
- [Golang v1.21](https://go.dev/doc/install)
- [Echo v4](https://github.com/labstack/echo) for router management
- [Testify](https://github.com/stretchr/testify) for unit testing
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) for validate request data


### How to run
- Clone the repository at local machine. 
> $ git clone https://github.com/iamganeshagrawal/go-crud-api-assignment.git
- Open terminal into cloned dir
- Run `go mod download ` to install all deps
- Run `go run .` to start web application
- Use **VsCode** + **REST Client** to access APIs OR use **Postman** with base url `http://localhost:8080`


### REST API
- `GET http://localhost:8080/ping` - Health check rest api
- `GET http://localhost:8080/api/v1/employees` - Get List of employees 
  - Query Params
    - `page` - get specific page (default 1)
    - `limit` - limit of data on a page (default 10, no limit = -1)
- `POST http://localhost:8080/api/v1/employees` - Create a new employee
```
// Content-Type: application/json
{
    "name": "Ganesh Agrawal",
    "position": "Software Engineer",
    "salary": 9999999.00
}
```
- `GET http://localhost:8080/api/v1/employees/{id}` - Get a employee data using ID
- `DELETE http://localhost:8080/api/v1/employees/{id}` - Delete a employee data using ID
- `PUT http://localhost:8080/api/v1/employees/{id}` - Update a employee data using ID
```
// Content-Type: application/json
{
    "name": "Ganesh Agrawal",
    "position": "Software Engineer",
    "salary": 9999999.00
}
```


### Folder Structure
- `/models` - database schemas struct
- `/repository` - database access layer
- `/internal` - internal helpers
  - `/datatypes` - user defined datatypes
- `/main.go` - entry point file
- `go.*` - golang dep managemnt files