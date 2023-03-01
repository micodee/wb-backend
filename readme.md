install 1 group route :
  go mod init backend
  go get github.com/labstack/echo/v4
  npm i -g nodemon => nodemon --exec go run .

install 2 fetching query with gorm :
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/mysql
  go get github.com/go-playground/validator/v10


go mod tidy && go mod download = npm i

c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON) || c.Response().Header().Set("Content-Type", "application/json")
c.Response().WriteHeader(200) || c.Response().WriteHeader(http.statusOK)

buat database = mysql > models > database
buat product = models > repo > dto > controllers > routes > routes