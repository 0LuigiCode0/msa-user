module x-msa-user

go 1.16

replace (
	x-msa-auth v0.0.0 => ./modules/x-msa-auth
	x-msa-core v0.0.0 => ./modules/x-msa-core
)

require (
	github.com/0LuigiCode0/go-utill v1.0.9
	github.com/0LuigiCode0/logger v1.1.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/klauspost/compress v1.13.1 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.7.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	google.golang.org/genproto v0.0.0-20210729151513-df9385d47c1b // indirect
	x-msa-auth v0.0.0
	x-msa-core v0.0.0
)
