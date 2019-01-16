# Server

## Run on Windows

Install [Docker Toolbox](https://docs.docker.com/toolbox/overview/)

Create Docker Machine
```
docker-machine create rocket
```

Connect console to created docker machine. Run command from machine environmet to configure your shell.
```
docker-machine env rocket

For example: @FOR /f "tokens=*" %i IN ('docker-machine env rocket') DO @%i
```

Verify docker.

```
docker info
```

## Build Server

Go to server's root (rocket-server\server) directory to run Docker Compose.

```
docker-compose build
```

## Run Server

```
docker-compose up
```
