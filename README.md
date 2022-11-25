## Golang Api Framework

This template should help get you started developing with Golang

Current Project is Based on Gorm and Gin-Gonic


## Project Setup
Clone project and run
```sh
go get
```
assuming go 1.18 or greater is already installed on system

## Database Config .env
copy .env.example and make .env file fill in Database configs 

### Compile with 
This will also create migrations in database if tables are not already created
```sh
go run main.go
```
note that project does not support hot reload 

## Router functions , http and response handling Docs
Gin-Gonic [Docs](https://github.com/gin-gonic/gin).

## Database and ORM Docs
GORM [Docs](https://gorm.io/docs/).


