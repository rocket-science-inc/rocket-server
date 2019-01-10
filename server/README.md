# rocket-server

Setup new project
glide init
glide get github.com/go-kit/kit

# Create new microservice

Create new service
kit new service [name]

Update types and service

Generate updated service
kit generate service [name] --dmw --gorilla 

Get dependencies
glide get github.com/gorilla/mux
glide get github.com/gorilla/handlers

Test Run
go run [name]/cmd/main.go

Create DB
mkdir [name]/pkg/db
touch [name]/pkg/db/redis.go

```
package db

import (
	"github.com/go-redis/redis"
)
```

Get dependencies
glide get github.com/go-redis/redis

Update services

Generate Docker with Glide
kit generate docker --glide

Add new services endpoints

Generate new endpoints
kit generate service [name] --dmw --gorilla

Add middleware with endpoints
kit generate middleware auth -s [name] -e

Add GRPC transport
kit generate service [name] -w -t grpc
