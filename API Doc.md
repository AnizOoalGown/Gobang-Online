# API Doc

## Hall Chat

Request from player named Tom

```json
{
    "code": 2,
    "data": "Hello!"
}
```

Response to all players

```json
{
    "code": 2,
    "data": {
        "time": "2020-03-31 16:43:00",
        "from": "Tom",
        "content": "Hello!"
    }
}
```



## Create Room

Request from a player

```json
{
    "code": 3,
    "data": "host"
}
```

Response to all players

```json
{
    "code": 3,
    "data": {
        
    }
}
```



## Enter Room

Request from a player

```json
{
    "code": 4,
    "data": {
        "rid": "1123wqe1e",
        "role": "host"
    }
}
```

Response to all players

```json

```

