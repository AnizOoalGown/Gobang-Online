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

## Get Hall Dialog

Request from a player

```json
{
    "code": 3
}
```

Response to the player

```json
{
    "code": 3,
    "data": [{
        "time":"2020-03-31 14:57:00",
        "from":"123",
        "content":"666"
    }, {
        "time":"2020-03-31 15:54:38",
        "from":"5be6c60a-b7bf-4a4b-8938-a8877edbe8ec",
        "content":"哇塞期待期待期待的去"
    }]
}
```

## Get Rooms

Request from a player

```json
{
    "code": 4
}
```

Response to the player

```json
{
    "code":4,
    "data":[{
        "id":"r123",
        "dialog":null,
        "steps":null,
        "started":false,
        "host":{
            "id":"p125",
            "name":"Tom",
            "status":"leisure",
            "role":"host",
            "color":0,
            "turn":false,
            "ready":false
        },
        "challenger":{
            "id":"",
            "name":"",
            "status":"",
            "role":"",
            "color":0,
            "turn":false,
            "ready":false
        },
        "spectators":null
    }]
}
```



## Create Room

Request from a player

```json
{
    "code": 5,
    "data": 0
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

## Leave Room

## Room Chat

## Get Player

## Get Players

## Player Rename

## Set Player Status

## Set Ready

## Make Step

## Retract Step

## Ask Draw