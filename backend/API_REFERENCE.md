# API REFERENCE

| API | endpoint | description |
| :-: | :-: | :-: |
| message | /message/ | something |
| user | /user/ | something |

## Message

- message data に関するAPI

### SendMessage

- method `POST`
- URI `/message/`

#### Request

```json:SendMessage request
{
    "message": "message-text",
    "userId": 1,
    "roomId": 2,
    "timestamp": "2006-01-02-15:45.999(+0700)"
}
```

| key | variable type | mandatory | description |
| :-: | :-: | :-: | :-: |
| message | string | **o** | message sentence |
| userId | int | **o** | contributor user id |
| roomId | int | **o** | talk room id |
| timestamp | timestamp | **o** | message publish time |

#### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 400 | request parameter error |
| 500 | internal server error |

- response body

```json:GetMessage response
{
    "id": 1
}
```

| key | variable type | nullable | description |
| :-: | :-: | :-: | :-: |
| id | int | x | published message id |

### GetMessage

***未完成***

- method `GET`
- URI `/message/{message id}`

#### Request

```json:GetMessage request
{
    "id":[
        1,2,3
    ]
}
```

| key | variable type | description |
| :-: | :-: | :-: |
| id | []int | id array of message to retrieve |

#### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 500 | internal server error |

```json:message response
{
    "messages":[
        {
            "id":1,
            "message":"message-text1",
            "contributer":"user_name",
            "timestamp":"2006-01-02-15:45.999(+0700)"
        },
        {
            "id":2,
            "message":"message-text2",
            "contributer":"user_name",
            "timestamp":"2006-01-02-15:45.999(+0700)"
        },
        {
            "id":3,
            "message":"message-text3",
            "contributer":"user_name",
            "timestamp":"2006-01-02-15:45.999(+0700)"
        }
    ]
}
```

| key | variable type | nullable | description |
| :-: | :-: | :-: | :-: |
| messages | JSON array | x | message array |
| messages[*].id | int | x | id of message to retrieve |
| messages[*].message | string | x | message |
| messages[*].contributer | int | x | user name |
| messages[*].timestamp | int | x | message send time |

### MessageID
***無し, roomAPIに追加するかも***
- message id の配列を返却

## User

- user information

### Create User

- create new user

- method `POST`
- URI `/user/`

#### Request

```json:CreateUser request
{
    "name": "Albert Einstein",
    "mail": "albert.e@domain.com",
    "sex": 1
}
```

| key | variable type | mandatory | description |
| :-: | :-: | :-: | :-: |
| name | string | **o** | user name |
| mail | string | x | email address |
| sex | int | x | 1: male, 2: female, 3: other |

#### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 400 | request parameter error |
| 500 | internal server error |

- response body

```json:CreateUser response
{
    "id": 1
}
```

| key | variable type | nullable | description |
| :-: | :-: | :-: | :-: |
| id | int | x | created user id |

### Get User

- get user information

- method `GET`
- URI `/user/{user_id}`

#### Request

- query parameter

| name | variable type | description |
| :-: | :-: | :-: |
| user_id | int | user id to retrieve |

### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 500 | internal server error |

```ison:user response
{
    "userId":1,
    "userName":"Albert Einstein"
}
```

## Group

### CreateGroup

- method `POST`
- URI `/group/`

### Request

```json:group request
{
    "id":[
        1,2,3
    ]
}
```

| key | variable type | description |
| :-: | :-: | :-: |
| id | []int | id array of group to retrieve |

### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 500 | internal server error |

```ison:user response
{
    "groups"[
        {
            "groupId":1,
            "groupName":"nims"
        },
        {
            "groupId":2,
            "groupName":"cern"
        },
        {
            "groupId":3,
            "groupName":"nasa"
        }
    ]
}
```


| key | variable type | description |
| :-: | :-: | :-: |
| groups | JSON array | group array |
| groups[*].groupID | int | group id |
| groups[*].groupName | string | group name |

## GroupUser

- method `GET`
- URI `/group/{group id}`

### Request

- query parameter

| name | variable type | description |
| :-: | :-: | :-: |
| group id | int | group id to retrieve |

### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 500 | internal server error |

```ison:groupUser response
{
    "groupId":1,
    "groupName":"Albert Einstein",
    "userId":[
        1,2,3
    ]
}
```

## MessageUser

- method `POST`
- URI `/message/`

### Request

- query parameter

| name | variable type | description |
| :-: | :-: | :-: |
| group id | int | group id to retrieve |


### Response

- http status

| status code | description |
| :-: | :-: |
| 200 | success |
| 500 | internal server error |

```ison:userMessage response
{
    "userId":1,
    "userName":"Albert Einstein"
}
```
