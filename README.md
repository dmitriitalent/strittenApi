# strittenApi
Golang api for stritten-project 

Database - Postgres at docker container

config.yaml example: 
```
port: "8083"

db:
  host: "localhost"
  port: "5431"
  username: "postgres"
  password: "password"
  dbname: "postgres"
  sslmode: "disable"

crypto:
  hashpasswordcost: 20
  accessTokenSalt: "1"
  refreshTokenSalt: "2"
```
