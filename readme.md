go mod init go-swaggo-errorcontract
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/sirupsen/logrus
go get github.com/swaggo/echo-swagger
go install github.com/swaggo/swag/cmd/swag@latest

swag --version

swag init

go run .


<!-- CURL Commands -->
curl -X GET http://localhost:8080/users
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@mail.com"}'
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"Bob","email":"bob@mail.com"}'
curl -X POST http://localhost:8080/submit-form -d "name=Alice&email=alice@mail.com"