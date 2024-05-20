# Gin Web CRUD
### Set .env file
Use the .env file to set your DB user and password, and configure the gin db connection.
```
# docker-compose env
DB_ROOT_PASSWORD=
DB_USER=
DB_USER_PASSWORD=
```
```
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
Start Gin web server:
```
app
```
