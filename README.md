# Golang visual Programming backend

For init project clone this repository

- then create main.go file

    ``mkdir main.go``

- initilize mod

    ``go mod init kenny.com/api/gorilla-mux``

### Install dependencies

- `go get -u github.com/go-chi/chi/v5`

- ``go get -u github.com/gorilla/mux``

- `go get -u github.com/dgraph-io/dgo/v210`

- `go get github.com/joho/godotenv/cmd/godotenv`
  
  go >= 1.17

- `go install github.com/joho/godotenv/cmd/godotenv@latest`


CompileDaemon:
    
- ``go get github.com/githubnemo/CompileDaemon``

For go 1.18

- `go install github.com/githubnemo/CompileDaemon@latest`

- `CompileDaemon`

- ``CompileDaemon -command="WORKING_DIRECTORY.exe"``

- `CompileDaemon -command="visual-programming.exe"`

### Run Project

``go run main.go``

### Patron repository
- concreto  -> handler -> GetUserByIdPostgres -> ... . ... . .. .. . . 
  el codigo es volátil 

- Abstraccion
- handler - GetUserById -> User
        -  Postgres
        - MongoDB
        - ...
        
  inyeccion de dependencias:      
        - Postgres
        - MongoDB
        - ...

pertenece a uno de los principios de SOLID

## Instalacion dgraph


## Install Dgraph DB with Docker

> docker run --rm -it -p "8080:8080" -p "9080:9080" -p "8000:8000" -v ~/dgraph:/dgraph "dgraph/standalone:v21.12.0"


- port: 8000 -> UI
- port: 8080 -> RESTAPI

## Flujo de REST API

principal -> main


## Para implementar en Postgres

> go get github.com/lib/pq