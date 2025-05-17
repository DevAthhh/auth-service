# Auth Service
___
This project was made purely for myself in order to write less of the same code, in it I tried to recreate the DDD architecture in order to improve this experience.

# EndPoint's
___
- `/check` - validate jwt token
- `/register` - registration
- `/login` - getting a token
- `/refresh` - update a token

General json request:
``` json
{
  "email": "",
  "username": "",
  "passwrod": "",
  "token": ""
}
```
