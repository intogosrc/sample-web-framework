### basic

host  
http://127.0.0.1:3000  


### test apis

URI: /auth/test

parameters:

| name | type | info |
| ---- | ---- | ---- |
| auth_token | string | use to check auth |

example of response:

```json
{
    "code": 200,
    "data": {
        "all": [
            {
                "Id": 1,
                "Name": "name_1659411693"
            },
            {
                "Id": 5,
                "Name": "name_1660535886"
            }
        ],
        "one": {
            "Id": 1,
            "Name": "name_1659411693"
        }
    },
    "msg": ""
}
```