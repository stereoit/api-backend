# Eventival backend

This is experimental project where I learn about hexagonal architecture.

At this moment I am focusing to learn the concepts, have User model available with all basic operations (list, add..) expose via gRPC GW. Optionally through REST as well.

## Prerequisities

```
# npm install --unsafe-perm -g grpcc
$ make protoc
$ echo "PORT=8000" > .env
```


## Running

Start a server

`$ make run` 

Then in another terminal

`$ make grpcc`

Register and list all users

```
    UserService@127.0.0.1:8000> client.registerUser({'email':'robert.smol@stereoit.com'}, pr)
    EventEmitter {}
    serService@127.0.0.1:8000> client.listUser({}, pr)
    EventEmitter {}
    UserService@127.0.0.1:8000> 
    {
    "users": [
        {
        "id": "cceb464d-b689-11e9-872d-1c3947113383",
        "email": "robert.smol@stereoit.com"
        }
    ]
    }
```