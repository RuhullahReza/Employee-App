# Employee-App

# How To Run App Localy
For runing this app localy, first we need to clone this repository

```
git clone https://github.com/RuhullahReza/Employee-App.git
```

After we clone the repository, we can run command for docker compose

```
docker compose up
```

After the container is active, Employee app ready to accept request

# Employee API Documentation

## Create Employee

Endpoint to create a new employee.

- **URL:** `http://127.0.0.1:8080/api/employee`
- **Method:** `POST`
- **Content-Type:** `application/json`

### Request Body


| Field      | Type   | Description                                                      |
|------------|--------|------------------------------------------------------------------|
| first_name | string | First name of the employee. Should contain alphabets and spaces only. |
| last_name  | string | Last name of the employee. Should contain alphabets and spaces only.                                      |
| email      | string | Email address of the employee. Should be unique for each user.   |
| hire_date  | string | Hire date of the employee (YYYY-MM-DD).                           |

### Example
```json
{
    "first_name": "abc",
    "last_name": "def",
    "email": "abc.def@gmail.com",
    "hire_date": "2024-05-04"
}
```


### Response

HTTP Status Codes:

**Success Response**

**201 Created**
```json
{
  "code": "Created",
  "message": "Successfully created new employee",
  "data": {
    "id": 1,
    "first_name": "Abc",
    "last_name": "Def",
    "email": "abc.def@gmail.com",
    "hire_date": "2024-05-04T00:00:00Z",
    "created_at": "2024-05-05T11:42:21.962919678Z",
    "updated_at": "2024-05-05T11:42:21.962919678Z"
  },
  "serverTime": 1714909341970
}
```

**Error Response**
```json
{
    "code": "Bad Request",
    "message": "invalid name format",
    "serverTime": 1714909901155
}
```

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.

**500 Internal Server Error :** Something went wrong on the server side.

## Get Employee By Id

Endpoint to retrieve data for a specific employee.

- **URL:** `http://127.0.0.1:8080/api/employee/{employee_id}`
- **Method:** `GET`
- **Content-Type:** `application/json`

### Path Parameters
| Parameter    | Type     | Description                                 |
|--------------|----------|---------------------------------------------|
| employee_id  | integer  | The unique identifier of the employee.      |


### Request Body

### Response

HTTP Status Codes:

**Success Response**

**200 OK**
```json
{
    "code": "OK",
    "message": "Successfully get data for employee id 1",
    "data": {
        "id": 1,
        "first_name": "Abc",
        "last_name": "Def",
        "email": "abc.def@gmail.com",
        "hire_date": "2024-05-01T00:00:00Z",
        "created_at": "2024-05-05T08:57:10.729112Z",
        "updated_at": "2024-05-05T08:57:44.453695Z"
    },
    "serverTime": 1714910829304
}
```

**Error Response**
```json
{
    "code": "Not Found",
    "message": "employee with id 11 not found",
    "serverTime": 1714909901155
}
```

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.

**404 Not Found :** Employee with the specified ID does not exist

**500 Internal Server Error :** Something went wrong on the server side.

## Get All Employee

Endpoint to retrieve all employee data with pagination and sorting options.

- **URL:** `http://127.0.0.1:8080/api/employee`
- **Method:** `GET`
- **Content-Type:** `application/json`

### Path Parameters
| Parameter | Type    | Description                                  | Default Value |
|-----------|---------|----------------------------------------------|---------------|
| pageNum   | integer | Specifies the page number.                   | 1             |
| pageSize  | integer | Specifies the number of items per page.      | 20            |
| orderBy   | string  | Specifies the field to order the results by (id, first_name, last_name, email, hire_date). | created_at     |
| sort      | string  | Specifies the sorting order (ASC/DESC).      | DESC           |

*if the value that passed into parameter is invalid, then default value will be used*


### Request Body

### Response

HTTP Status Codes:

**Success Response**

**200 OK**
```json
{
    "code": "OK",
    "message": "Successfully get all employee data",
    "data": {
        "page_number": 1,
        "page_size": 5,
        "total_page": 1,
        "data": [
            {
                "id": 1,
                "first_name": "Abc",
                "last_name": "Def",
                "email": "abc.def@gmail.com",
                "hire_date": "2024-05-01T00:00:00Z",
                "created_at": "2024-05-05T08:57:10.729112Z",
                "updated_at": "2024-05-05T08:57:44.453695Z"
            },
            {
                "id": 2,
                "first_name": "Zxc",
                "last_name": "Xcv",
                "email": "zxc.xcv@gmail.com",
                "hire_date": "2024-05-02T00:00:00Z",
                "created_at": "2024-05-05T11:42:21.962919Z",
                "updated_at": "2024-05-05T11:42:21.962919Z"
            }
        ]
    },
    "serverTime": 1714913519908
}

```

**Error Response**
```json
{
    "code": "Internal Server Error",
    "message": "something went wrong",
    "serverTime": 1714913519908
}
```

**Error Codes**

**500 Internal Server Error :** Something went wrong on the server side.

## Update Employee by Id

Endpoint to update data for a specific employee.

- **URL:** `http://127.0.0.1:8080/api/employee/{employee_id}`
- **Method:** `PUT`
- **Content-Type:** `application/json`

### Request Body


| Field      | Type   | Description                                                      |
|------------|--------|------------------------------------------------------------------|
| first_name | string | First name of the employee. Should contain alphabets and spaces only. |
| last_name  | string | Last name of the employee. Should contain alphabets and spaces only.                                      |
| email      | string | Email address of the employee. Should be unique for each user.   |
| hire_date  | string | Hire date of the employee (YYYY-MM-DD).                           |

### Example
```json
{
    "first_name": "abc update",
    "last_name": "def",
    "email": "abc.def@gmail.com",
    "hire_date": "2024-05-01"
}
```


### Response

HTTP Status Codes:

**Success Response**

**200 OK**
```json
{
    "code": "OK",
    "message": "Successfully update data for employee id 1",
    "data": {
        "id": 1,
        "first_name": "Abc Update",
        "last_name": "Def",
        "email": "abc.def@gmail.com",
        "hire_date": "2024-05-01T00:00:00Z",
        "updated_at": "2024-05-05T13:06:40.856149334Z"
    },
    "serverTime": 1714914400863
}
```

**Error Response**
```json
{
    "code": "Bad Request",
    "message": "duplicate email",
    "serverTime": 1714914455192
}
```

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.

**404 Not Found :** Employee with the specified ID does not exist.

**500 Internal Server Error :** Something went wrong on the server side.

## Delete Employee by Id

Endpoint to delete data for a specific employee (soft delete).

- **URL:** `http://127.0.0.1:8080/api/employee/{employee_id}`
- **Method:** `DELETE`
- **Content-Type:** `application/json`

### Path Parameters
| Parameter    | Type     | Description                                 |
|--------------|----------|---------------------------------------------|
| employee_id  | integer  | The unique identifier of the employee.      |

### Request Body

### Response

HTTP Status Codes:

**Success Response**

**200 OK**
```json
{
    "code": "OK",
    "message": "Successfully delete data for employee id 2",
    "serverTime": 1714914792382
}
```

**Error Response**
```json
{
    "code": "Bad Request",
    "message": "invalid id",
    "serverTime": 1714914828635
}
```

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.

**404 Not Found :** Employee with the specified ID does not exist

**500 Internal Server Error :** Something went wrong on the server side.