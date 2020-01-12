# Eventival backend

This is experimental project where I learn about clean architecture.

At this moment I am focusing to learn the concepts, have User model available with all basic operations (list, add..) expose via gRPC GW and REST server as well.

## Prerequisities

```
# npm install --unsafe-perm -g grpcc
$ make protoc
$ cat << EOF > .env
REST_PORT=8000 
GRPC_PORT=8001
EOF
```


## Running

Start a GRPC server

`$ make grpc` 

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

## Container

[Podman](https://podman.io/getting-started/installation) is used instead of Docker. If you are using docker, just substitute the `podman` for `docker` (they are options compatible).

Build image

`$ make docker`

Run image:

`$ make run-docker`


## Testing

For unit tests run:

    $ make test

For integration tests:

    $ make integration-test


### Rest API

Run UI for provided openapi specification:

    $ docker run --rm -d --name swagger-ui -p 8888:8000 -v /path/to/code/api/:/docs shotat/swagger-ui-watcher -- /docs/openapi.yaml
    $ xdg-open http://localhost:8888

Then you can [browse](http://localhost:8888) the documentation.

