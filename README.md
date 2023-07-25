## Golang Api Framework

This template should help get you started developing with Golang

Current Project is Based on Gorm and Gin-Gonic


## Project Setup
Clone project and run
```sh
go get
```
# Project Setup
## Install Cli
```sh
go install https://github.com/usama-tariq1/leet-astro
```
Test Cli
```sh
leet-astro
```
## Create project 
```sh
leet-astro init ProjectName "ModUser"
```
example
```sh
leet-astro init Test "github.com/usama-tariq1"
```

assuming go 1.18 or greater is already installed on system

## Install Dependencies 
```sh
go get
```

## Database Config .env
copy .env.example and make .env file fill in Database configs 

## Compile with 
This will also create migrations in database if tables are not already created

### With Cli
```sh
leet-astro serve
```
### With Out Cli
```sh
go run main.go
```
note that project does not support hot reload 

# Cli Create Commands
## Create Controller 
With Resource and methods already built
```sh
leet-astro create controller ControllerName --model=ModelName
```
With out methods
```sh
leet-astro create controller ControllerName
```

## Create Model
create just model with query methods
```sh
leet-astro create model ModalName
```
With controller auto generated 
```sh
leet-astro create model Name --controller=true
```
With controller and router auto generated
```sh
leet-astro create model Name --controller=true --router=true
```

## Create Router
```sh
leet-astro create router NameRouter
```

## Create Middleware
```sh
leet-astro create middleware Name
```

#

## Router functions , http and response handling Docs
Gin-Gonic [Docs](https://github.com/gin-gonic/gin).

#
## Database and ORM Docs
GORM [Docs](https://gorm.io/docs/).


