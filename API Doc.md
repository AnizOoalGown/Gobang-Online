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

Response to all players in the room

```json
{"code":6,"data":{"id":"a6c2d2b3-d592-4a55-a1e3-dffd398b58f2","dialog":[],"steps":[],"started":false,"host":{"id":"b5448fcd-2626-46d1-8f9b-81b3ab184a86","name":"unnamed","status":"leisure","role":"host","color":0,"turn":false,"ready":false},"challenger":{"id":"34405814-1a1e-425e-aa27-445631c5afa5","name":"unnamed","status":"leisure","role":"challenger","color":1,"turn":false,"ready":false},"spectators":[]}}
```

## Leave Room

Request from a player

```json
{
    "code": 7,
    "data": "aae57069-9ceb-4296-a05d-0c010324dbac"
}
```

Response to the players left in the room if there is a host

```json
{"code":7,"data":{"id":"aae57069-9ceb-4296-a05d-0c010324dbac","dialog":[],"steps":[],"started":false,"host":{"id":"989be167-04fc-4b38-8132-da49addde697","name":"unnamed","status":"leisure","role":"challenger","color":1,"turn":false,"ready":false},"challenger":{"id":"","name":"","status":"","role":"","color":0,"turn":false,"ready":false},"spectators":[]}}
```

Response to the players left in the room if there is no host or challenger

```json
{
    "code": 8,
    "data": "aae57069-9ceb-4296-a05d-0c010324dbac"
}
```

## Del Room

Response to the player

```json
{
    "code": 8,
    "data": "84c2dfd3-7996-4658-b2b1-e09f3c517fcb"
}
```

## Room Chat

Request from a player

```json
{
    "code": 9,
    "data": {
        "from": "Tom",
        "content": "Hello!",
        "rid": "bbc8dca5-55e7-4954-a1ce-7c4afc06ed1f"
    }
}
```

Response to the players in the room

```json
{
    "code": 9,
    "data": {
        "rid": "bbc8dca5-55e7-4954-a1ce-7c4afc06ed1f",
        "time": "2020-04-04 21:55:07",
        "from": "Tom",
        "content": "Hello!"
    }
}
```

## Get Player

Request from a player

```json
{
    "code": 10
}
```



## Get Players

Request from a player

```json
{
    "code": 11
}
```

Response to the player

```json
{"code":11,"data":[{"id":"cf67d081-7637-4901-93af-9d4b3c824ba8","name":"unnamed","status":"leisure"},{"id":"c6f2cf85-503e-4a97-b104-06e7964c08e1","name":"unnamed","status":"leisure"},{"id":"bcfccae8-5027-4a33-ae1b-11e789951eaa","name":"unnamed","status":"leisure"},{"id":"e54d5c3c-5618-447f-b255-eb0666f90450","name":"unnamed","status":"leisure"},{"id":"41c1faa2-d88f-4b96-a17f-fb6400672b39","name":"unnamed","status":"leisure"},{"id":"74d655af-3916-4400-9611-e8ad1618d5e4","name":"unnamed","status":"leisure"},{"id":"12b2b259-1812-4095-80c7-67420d146b67","name":"unnamed","status":"leisure"},{"id":"e5e3cc75-da05-4e89-9fce-4639a9821768","name":"unnamed","status":"leisure"},{"id":"f1572da9-8710-48d6-99c1-d8b85bc368df","name":"unnamed","status":"leisure"},{"id":"937cce34-3ecf-401e-953d-d85ae3d16960","name":"unnamed","status":"leisure"}]}
```

## Player Rename

Request from a player

```json
{
    "code": 12,
    "data": "Tom"
}
```

Response to all players

```json
{"code":11,"data":[{"id":"cf67d081-7637-4901-93af-9d4b3c824ba8","name":"Tom","status":"leisure"},{"id":"c6f2cf85-503e-4a97-b104-06e7964c08e1","name":"unnamed","status":"leisure"},{"id":"bcfccae8-5027-4a33-ae1b-11e789951eaa","name":"unnamed","status":"leisure"},{"id":"e54d5c3c-5618-447f-b255-eb0666f90450","name":"unnamed","status":"leisure"},{"id":"41c1faa2-d88f-4b96-a17f-fb6400672b39","name":"unnamed","status":"leisure"},{"id":"74d655af-3916-4400-9611-e8ad1618d5e4","name":"unnamed","status":"leisure"},{"id":"12b2b259-1812-4095-80c7-67420d146b67","name":"unnamed","status":"leisure"},{"id":"e5e3cc75-da05-4e89-9fce-4639a9821768","name":"unnamed","status":"leisure"},{"id":"f1572da9-8710-48d6-99c1-d8b85bc368df","name":"unnamed","status":"leisure"},{"id":"937cce34-3ecf-401e-953d-d85ae3d16960","name":"unnamed","status":"leisure"}]}
```



## Set Player Status

Request from a player

```json
{
    "code": 13,
    "data": "chessing"
}
```

Response to all players

```json
{"code":11,"data":[{"id":"cf67d081-7637-4901-93af-9d4b3c824ba8","name":"unnamed","status":"leisure"},{"id":"c6f2cf85-503e-4a97-b104-06e7964c08e1","name":"unnamed","status":"leisure"},{"id":"bcfccae8-5027-4a33-ae1b-11e789951eaa","name":"unnamed","status":"leisure"},{"id":"e54d5c3c-5618-447f-b255-eb0666f90450","name":"unnamed","status":"leisure"},{"id":"41c1faa2-d88f-4b96-a17f-fb6400672b39","name":"unnamed","status":"leisure"},{"id":"74d655af-3916-4400-9611-e8ad1618d5e4","name":"unnamed","status":"leisure"},{"id":"12b2b259-1812-4095-80c7-67420d146b67","name":"unnamed","status":"leisure"},{"id":"e5e3cc75-da05-4e89-9fce-4639a9821768","name":"unnamed","status":"leisure"},{"id":"f1572da9-8710-48d6-99c1-d8b85bc368df","name":"unnamed","status":"leisure"},{"id":"937cce34-3ecf-401e-953d-d85ae3d16960","name":"unnamed","status":"leisure"}]}
```



## Set Ready

Request from a player

```json
{
    "code": 14,
    "data": {
        "rid": "27f71f22-07d9-4cad-9008-1f9806fbd1af",
        "ready": true
    }
}
```

Response to the players in the room. Data is room.

```json
{
    "code": 14,
    "data": {
        
    }
}
```



## Make Step

Request from a player

```json
{
    "code": 15,
    "data": {
        "rid": "27f71f22-07d9-4cad-9008-1f9806fbd1af",
        "i": 7,
        "j": 7
    }
}
```



## Retract Step

## Ask Draw