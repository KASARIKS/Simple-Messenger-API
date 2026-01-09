For start run or build ./cmd/main.go
API starts at http://localhost

For the error case, the response is in json:
{
  "error": "some error"
}
The responses below are labeled for the success case,
and have Headers:
  - Content-Type: application/json;

POST /register - create user
Headers:
  - Content-Type: application/json;
Input raw json data:
{
  "id":"custom user id",
  "password":"password",
  "nickname":"nickname"
}
Response:
{
  nil
}

POST /login - get autherization token
Headers:
  - Content-Type: application/json;
Input raw json data:
{
  "id":"customUserId",
  "password":"password"
}
Response:
{
  "token":"token"
}

DELETE /delete - delete user
Headers:
  - Content-Type: application/json;
  - Authorization: token;
Response:
{
  nil
}

POST /sendmessage - create message
Headers:
  - Content-Type: application/json;
  - Authorization: token;
Input raw json data:
{
  "recipientId":"recipientId",
  "value":"message itself"
}
Response:
{
  nil
}

An example of a message for the GET requests below:
{
  "id": 1,
  "authorId": "authorId",
  "recipientId": "recipientId",
  "value": "a message itself",
  "createdAt": time.Now().String() // A value from this construction, this is from Golang.
}

GET /getmessages?page=pageNumber - get sended messages
pageNumber - integer of one page, bigger or equal 0. There are 10 messages on one page
Headers:
  - Content-Type: application/json;
  - Authorization: token;
Response:
{
  [
    There are messages here or an empty array
  ]
}

GET /getgottenmessages?page=pageNumber - get gotten messages
pageNumber - integer of one page, bigger or equal 0. There are 10 messages on one page
Headers:
  - Content-Type: application/json;
  - Authorization: token;
Response:
{
  [
    There are messages here or an empty array
  ]
}
