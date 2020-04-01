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
{"code":5,"data":{"id":"34c1bd1f-b27c-4dcc-9521-3dcccf673e53","dialog":[],"steps":[],"started":false,"host":{"id":"937cce34-3ecf-401e-953d-d85ae3d16960","name":"unnamed","status":"leisure","role":"host","color":0,"turn":false,"ready":false},"challenger":{"id":"","name":"","status":"","role":"","color":0,"turn":false,"ready":false},"spectators":[]}}
```

## Enter Room

Request from a player

```json
{
    "code": 6,
    "data": {
        "rid": "a6c2d2b3-d592-4a55-a1e3-dffd398b58f2",
        "role": "challenger"
    }
}
```

Response to all players

```json
{"code":6,"data":{"id":"a6c2d2b3-d592-4a55-a1e3-dffd398b58f2","dialog":[],"steps":[],"started":false,"host":{"id":"b5448fcd-2626-46d1-8f9b-81b3ab184a86","name":"unnamed","status":"leisure","role":"host","color":0,"turn":false,"ready":false},"challenger":{"id":"34405814-1a1e-425e-aa27-445631c5afa5","name":"unnamed","status":"leisure","role":"challenger","color":1,"turn":false,"ready":false},"spectators":[]}}
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