# Enpoints Api Devices

## Endpoints

All the endpoints in the list below have middleware to validate that the user is authenticated against the server. The mapping and setting of the middleware is done in the file `cmd\pkg\routes\route.handler.go`

Below are the available endpoints in the application.

```
[GET]               /ping                    
[GET-POST]          /dcnearshore/v1/devices  
[GET-PUT-DELETE]    /dcnearshore/v1/devices/:id 
[GET-POST]          /dcnearshore/v1/firmwares
[GET-PUT-DELETE]    /dcnearshore/v1/firmwares/:id
```

## How to use

In the postman collection you will find the functional structure of the endpoints along with their respective json, environment variables and authentication.