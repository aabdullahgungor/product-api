# Go RESTful API with Gin Web Framework & PostgreSQL
This is an example golang backend application using PostgreSQL database with clean architecture.

## Features
* Go Web Framework ([gin-gonic](https://github.com/gin-gonic/gin))
* Containerize ([docker](https://www.docker.com/))
* Swagger ([swaggo](https://github.com/swaggo/swag))
* CRUD operations
* Mock: [golang/mock](https://github.com/golang/mock)
* Database: [MongoDB](https://www.mongodb.com/)
* Test Assertions: [stretchr/testify](https://github.com/stretchr/testify)

## Getting Started

```sh
# download the project
git clone https://github.com/aabdullahgungor/product-api.git

cd product-api
```

### Run the Project

```bash
go run main.go
```
### Access API using 

```bash
http://localhost:8000/api/v1
```

### Sample of Endpoints

- GET localhost:8000/api/v1/products
- GET localhost:8000/api/v1/products/:id
- POST localhost:8000/api/v1/products
- PUT localhost:8000/api/v1/products
- DELETE localhost:8000/api/v1/products/:id
- ........

## Open API Doc Preview
http://localhost:8000/api/v1/swagger/index.html

![Swagger](.github/images/Swagger.png)

## How to run the test?

```bash
# Run tests
go test ./test/test_controller -v
go test ./test/test_service -v
```
