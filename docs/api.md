### 基础信息

host  
http://127.0.0.1:3000  


### 测试

URI: /auth/test

参数：

| name | type | info |
| ---- | ---- | ---- |
| auth_token | string | 验证信息 |

返回值示例：

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