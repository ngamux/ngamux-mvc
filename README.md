# ngamux-mvc
Model-View-Controller Using Ngamux

# Run Project
1. Run MongoDB first, you can use this command:
```bash
sudo mongod --dbpath /var/db/mongo
```

2. Configure the connection URL in [config.ConnectDatabase](https://github.com/ngamux/ngamux-mvc/blob/0c6088cac7df4fef96b6d7b4755c9b26e8a562d3/config/config.go#L11) function.
3. Run project use this command:
```bash
go run main.go
```

# API Docs
Host: http://localhost:8080

## Get All Users
### Request
Method: GET
URL: /users

### Response
Body:
```json
{
  "users": [
    {
      "email": "my@email.com",
      "password": "123123123"
    }
  ]
}```

## Create User
### Request
Method: GET
URL: /users
Query: ?email=my@email.com&password=123123123

### Response
Body:
```json
{
  "user": {
    "email": "my@email.com",
    "password": "123123123"
  }
}```
