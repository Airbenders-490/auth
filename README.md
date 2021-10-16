Uses Fiber for web framework.

To run server, run `go run ./main.go`

to test register:
- on postman, choose POST, url: http://localhost:3000/api/register
- set body as raw json with content:
{
  "name": "my name",
  "email": "myemail@gmail.com",
  "password": "mypassword12"
}
- send
- should see success message


to test login: 
- on postman, choose POST, url: http://localhost:3000/api/login
- set body as raw json with content:
{
  "email": "myemail@gmail.com",
  "password": "mypassword12"
}
- send
- should see success message and cookie


to test Retrieve user:
- login first
- on postman, choose GET, url: http://localhost:3000/api/user
- send
- should see user id, name, email



to test logout:
- on postman, choose POST, url: http://localhost:3000/api/logout
- send
- will see the cookie is gone