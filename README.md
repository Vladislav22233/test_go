Documentation for the Code
Description
This code is a simple web server that provides an API for managing segments and users. It uses a PostgreSQL database to store the data.
Dependencies
To run this code, you will need to have the following dependencies installed:
Go: https://golang.org/dl/
PostgreSQL: https://www.postgresql.org/download/
Installation
Clone the repository to your local machine.
Create a database in PostgreSQL with the name "segments".
Specify the database connection parameters: DBHost, DBPort, DBUser, DBPassword in the main.go file.
Run the following command to install the necessary dependencies:
go mod download
Start the application using the following command:
go run main.go
Usage
Create Segment
Method: POST
Path: /segment
Request Body:
{
    "id": 1,
    "slug": "segment1"
}
Example Response:
{
    "id": 1,
    "slug": "segment1"
}
Delete Segment
Method: DELETE
Path: /segment/{slug}
Example Request:
DELETE /segment/segment1
Example Response:
Status 200 OK
Add User Segments
Method: POST
Path: /user/{id}/segments
Request Body:
[
    {
        "id": 1,
        "slug": "segment1"
    },
    {
        "id": 2,
        "slug": "segment2"
    }
]
Example Response:
Status 200 OK
Delete User Segments
Method: DELETE
Path: /user/{id}/segments
Request Body:
[
    {
        "id": 1,
        "slug": "segment1"
    },
    {
        "id": 2,
        "slug": "segment2"
    }
]
Example Response:
Status 200 OK
Get User Segments
Method: GET
Path: /user/{id}/segments
Example Request:
GET /user/1/segments
Example Response:
[
    {
        "id": 1,
        "slug": "segment1"
    },
    {
        "id": 2,
        "slug": "segment2"
    }
]Documentation for the Code
Description
This code is a simple web server that provides an API for managing segments and users. It uses a PostgreSQL database to store the data.
Dependencies
To run this code, you will need to have the following dependencies installed:
Go: https://golang.org/dl/
PostgreSQL: https://www.postgresql.org/download/
Installation
Clone the repository to your local machine.
Create a database in PostgreSQL with the name "segments".
Specify the database connection parameters: DBHost, DBPort, DBUser, DBPassword in the main.go file.
Run the following command to install the necessary dependencies:
go mod download
Start the application using the following command:
go run main.go
Usage
Create Segment
Method: POST
Path: /segment
Request Body:
{
    "id": 1,
    "slug": "segment1"
}
Example Response:
{
    "id": 1,
    "slug": "segment1"
}
Delete Segment
Method: DELETE
Path: /segment/{slug}
Example Request:
DELETE /segment/segment1
Example Response:
Status 200 OK
Add User Segments
Method: POST
Path: /user/{id}/segments
Request Body:
[
    {
        "id": 1,
        "slug": "segment1"
    },
    {
        "id": 2,
        "slug": "segment2"
    }
]
Example Response:
Status 200 OK
Delete User Segments
Method: DELETE
Path: /user/{id}/segments
Request Body:
[
    {
        "id": 1,
        "slug": "segment1"
    },
    {
        "id": 2,
        "slug": "segment2"
    }
]
Example Response:
Status 200 OK
Get User Segments
Method: GET
Path: /user/{id}/segments
Example Request:
GET /user/1/segments
Example Response:
[
    {
        "id": 1,
        "slug": "segment1"
    },
    {
        "id": 2,
        "slug": "segment2"
    }
]
