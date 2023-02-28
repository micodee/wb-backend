install :
go mod init backend
go get github.com/labstack/echo/v4
npm i -g nodemon => nodemon --exec go run .

go mod tidy && go mod download = npm i

c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON) || c.Response().Header().Set("Content-Type", "application/json")
c.Response().WriteHeader(200) || c.Response().WriteHeader(http.statusOK)