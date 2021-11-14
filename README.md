[![Codacy Badge](https://app.codacy.com/project/badge/Grade/60f6d7c81fb4415cad928fd9daa17a9c)](https://www.codacy.com/gh/Airbenders-490/auth/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Airbenders-490/auth&amp;utm_campaign=Badge_Grade)

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
- should see token returned


to test Retrieve user:
- login first
- on postman, choose GET, url: http://localhost:3000/api/user
- send
- should see user id, name, email
