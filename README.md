# Gin Web CRUD
### Set .env file
Use the .env file to set your DB user and password, and configure the gin db connection.
```env
# docker-compose env
DB_ROOT_PASSWORD=
DB_USER=
DB_USER_PASSWORD=
```
```env
# web server env
PORT=
DB_URL="<user>:<password>@tcp(<host>:<port>)/<database>?charset=utf8mb4&parseTime=True&loc=Local"
```

### Running Gin and SQL
Start Docker containers:
```
docker-compose up -d
```
Enter the gin container:
```
docker exec -it go_gin_env bash
```
Create the schema when you first use:
```
go run migrate/migrate.go
```
Start Gin web server:
```
app
```

# API Instructions
### User
Register user : `POST` `{{url}}/auth/register`
```json
{
    "username": "Jarvis",
    "password": "123"
}
```

### Book
Create book : `POST` `{{url}}/api/book`
```json
{
    "name": "book1",
    "author": "author1",
    "publish_year": 2024
}
```


