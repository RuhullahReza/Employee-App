# Employee-App

# How To Run App Localy
To run this application locally, follow these steps

## 1. Clone the Repository

Clone the GitHub repository to your local machine using the following command:

```
git clone https://github.com/RuhullahReza/Employee-App.git
```

This will create a local copy of the project on your machine.


## 2. Set Up Docker Compose
Navigate to the root directory of the cloned repository in your terminal. Once there, you can run the Docker Compose command to set up and start the application

```
docker compose up
```
This command will orchestrate the necessary containers and services defined in the docker-compose.yml file to bring up the application.

## 3. Access the Application
Once Docker Compose has successfully started the application, you can access it using curl, postman, or web broser on `http://localhost:8080`.


# Unit Test
In this codebase, unit tests are primarily focused on testing the business logic within the handlers layer, service layer, and utility functions.

## Testing Strategy
- **Handlers Layer**: Unit tests for the handlers layer focus on verifying the behavior of HTTP request handlers and their interaction with service layer interfaces.
- **Service Layer**: Tests for the service layer ensure that business logic is correctly implemented and that services interact with data sources on repository layer.
- **Utility Functions**: Utility functions are also unit tested to ensure they perform their intended tasks accurately.

## Mocking
To isolate the components being tested and remove dependencies on external systems, we utilize mocking frameworks. Specifically, we use Mockery to automatically generate mocks for interfaces used within the service and handler layers. This allows us to simulate the behavior of dependencies during testing.

## Checking Test Coverage
To isolate the components being tested and remove dependencies on external systems, we utilize mocking frameworks. Specifically, we use Mockery to automatically generate mocks for interfaces used within the service and handler layers. This allows us to simulate the behavior of dependencies during testing.

```
go test ./... -coverprofile cover.out
go tool cover -func cover.out | grep total:
```
These commands will execute all tests in the project and output the total coverage percentage, giving insight into how much of the codebase is covered by unit tests.

# Employee API Documentation

## Create Employee

Endpoint to create a new employee.

- **URL:** `http://127.0.0.1:8080/api/employees`
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

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.
```json
{
    "code": "Bad Request",
    "message": "invalid name format",
    "serverTime": 1714909901155
}
```

```json
{
    "code": "Bad Request",
    "message": "duplicate email",
    "serverTime": 1714997222595
}
```

```json
{
    "code": "Bad Request",
    "message": "invalid date format",
    "serverTime": 1714997245643
}
```

**500 Internal Server Error :** Something went wrong on the server side.
```json
{
    "code": "Internal Server Error",
    "message": "something went wrong",
    "serverTime": 1714913519908
}
```

## Get Employee By Id

Endpoint to retrieve data for a specific employee.

- **URL:** `http://127.0.0.1:8080/api/employees/{employee_id}`
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

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.
```json
{
    "code": "Bad Request",
    "message": "invalid id",
    "serverTime": 1714997949081
}
```

**404 Not Found :** Employee with the specified ID does not exist
```json
{
    "code": "Not Found",
    "message": "employee with id 11 not found",
    "serverTime": 1714909901155
}
```

**500 Internal Server Error :** Something went wrong on the server side.
```json
{
    "code": "Internal Server Error",
    "message": "something went wrong",
    "serverTime": 1714913519908
}
```

## Get All Employee

Endpoint to retrieve all employee data with pagination and sorting options.

- **URL:** `http://127.0.0.1:8080/api/employees`
- **Method:** `GET`
- **Content-Type:** `application/json`

### Path Parameters
| Parameter | Type    | Description                                  | Default Value |
|-----------|---------|----------------------------------------------|---------------|
| pageNum   | integer | Specifies the page number.                   | 1             |
| pageSize  | integer | Specifies the number of items per page.      | 20            |
| orderBy   | string  | Specifies the field to order the results by (id, first_name, last_name, email, hire_date, created_at, updated_at). | created_at     |
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
```json
{
    "code": "Internal Server Error",
    "message": "something went wrong",
    "serverTime": 1714913519908
}
```

## Update Employee by Id

Endpoint to update data for a specific employee.

- **URL:** `http://127.0.0.1:8080/api/employees/{employee_id}`
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

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.
```json
{
    "code": "Bad Request",
    "message": "invalid id",
    "serverTime": 1714998143690
}
```

```json
{
    "code": "Bad Request",
    "message": "invalid name format",
    "serverTime": 1714909901155
}
```

```json
{
    "code": "Bad Request",
    "message": "duplicate email",
    "serverTime": 1714997222595
}
```

```json
{
    "code": "Bad Request",
    "message": "invalid date format",
    "serverTime": 1714997245643
}
```

**404 Not Found :** Employee with the specified ID does not exist.
```json
{
    "code": "Not Found",
    "message": "employee with id 11 not found",
    "serverTime": 1714998175547
}
```

**500 Internal Server Error :** Something went wrong on the server side.
```json
{
    "code": "Internal Server Error",
    "message": "something went wrong",
    "serverTime": 1714913519908
}
```

## Delete Employee by Id

Endpoint to delete data for a specific employee (soft delete).

- **URL:** `http://127.0.0.1:8080/api/employees/{employee_id}`
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

**Error Codes**

**400 Bad Request :** Invalid request body or parameters.
```json
{
    "code": "Bad Request",
    "message": "invalid id",
    "serverTime": 1714914828635
}
```

**404 Not Found :** Employee with the specified ID does not exist
```json
{
    "code": "Not Found",
    "message": "employee with id 1111 not found",
    "serverTime": 1714998220649
}
```

**500 Internal Server Error :** Something went wrong on the server side.
```json
{
    "code": "Internal Server Error",
    "message": "something went wrong",
    "serverTime": 1714913519908
}
```