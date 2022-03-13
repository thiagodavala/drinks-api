# Drinks API

## GET /swagger

```bash
curl http://localhost:8080/swagger/index.html
```

## GET /cocktails

List Cocktails

```bash
curl -X 'GET' 'http://localhost:8080/cocktails'
```

## GET /cocktail/id 

Get Cocktail by Id

```bash
curl -X 'GET' 'http://localhost:8080/cocktail/1'
```

## PUT /cocktail

Add Cocktail

```bash
curl -X PUT localhost:8080/cocktail -d '{"Name":"Negroni","Ingredients":"30 ml Gin, 30 ml Bitter Campari, 30 ml Sweet Red Vermouth","Method":"Pour all ingredients directly into a chilled old fashioned glass filled with ice, Stir gently.","Garnish":"Garnish with a half orange slice."}'
```

## DELETE /cocktail/id

Add Cocktail

```bash
curl -X DELETE localhost:8080/cocktail/1
```

## POST /login

```
curl -X POST localhost:8080/login -d '{"Email":"Negroni","Password":"30 ml Gin"}'
```

## Start Project

```
go mod init drinks-api
go mod tidy
```

## Run project

1. Start BD

```
db/docker-compose up
```

2. Start app

```
go run main.go
```

## Update Swagger

```
go get -u github.com/swaggo/swag/cmd/swag
swag init
```