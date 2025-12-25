# Todo API Server

## Cтруктура

```
todo-app/
├── cmd/
│   └── server/              
│       └── main.go
├── internal/
│   ├── handlers/            
│   │   └── handlers.go
│   ├── middleware/          
│   │   └── middleware.go
│   ├── models/              
│   │   └── models.go
│   ├── router/              
│   │   └── router.go
│   ├── service/             
│   │   └── service.go
│   └── storage/
│       ├── storage.go        
│       └── storage_test.go
├── Dockerfile
├── Makefile
└── go.mod
```

## Запуск

Локально:

```shell
go run cmd/server/main.go
```

Docker:

```shell
make docker-build
```

```shell
make docker-up
```

## Запуск тестов

```shell
make tests
```

## Endpoints

**POST /todos - создать новую задачу**

_Request:_
```json
{
    "title": "test",
    "body": "test"
}
```

_Response (201 Created):_
```json
{
    "id": 1,
    "title": "test",
    "body": "test",
    "isDone": false
}
```

**GET /todos - получить список всех задач**

_Response (200 OK):_
```json
[
    {
        "id": 1,
        "title": "test",
        "body": "test",
        "isDone": false
    }
]
```

**GET /todos/{id} - получить задачу по идентификатору**

_Response (200 OK):_
```json
{
    "id": 1,
    "title": "test",
    "body": "test",
    "isDone": false
}
```

**PUT /todos/{id} - обновить задачу по идентификатору**

_Request:_
```json
{
    "title": "update",
    "body": "update"
}
```

_Response (200 OK):_
```json
{
    "status": "updated"
}
```

**DELETE /todos/{id} - удалить задачу по идентификатору** 

_Response (204 No Content)_