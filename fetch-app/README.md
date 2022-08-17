# MyAPI

## Description
Simple API using clean code architecture

This project has  4 Domain layer :
 * Entity Layer
 * Usecase Layer
 * Controller Layer  
 * Delivery Layer

## First Run

```bash
$ go mod tidy
```

## Running the app
```bash
# Without docker
$ go run .

# Using docker
$ docker-compose up -d
```

## Tools Used:
 * [PostgreSQL](https://github.com/lib/pq)
 * [Iris Lib](https://github.com/kataras/iris)
 * [Xorm Lib](https://xorm.io/xorm)
 * [JWT](https://github.com/dgrijalva/jwt-go)
