To run the application:
1. Install Golang following the instruction on https://go.dev/doc/install
2. Install make for windows:
   Using Chocolatey. First you need to install this package manager. 
Once installed you simply need to install make (you may need to run it in an elevated/admin command prompt) :
choco install make
For macos: brew install make
For Linux: sudo apt install make
3. Install Docker all instructions can be found on https://www.docker.com/products/docker-desktop/
4. Create app.env file in the root of the project with following fields:
5. Install golang migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
DB_SOURCE=
SERVER_ADDRESS=localhost:8080
JWT_SECRET=
JWT_ISSER=
AUDIENCE=
TOKEN_EXPIRY=15m
REFRESH_EXPIRY=480m
COOKIE_DOMAIN=localhost
COOKIE_PATH=/
COOKIE_NAME=_Host-refresh_token

And set values appropriately
5. Run make commands in the Makefile:
make postgres,
make createdb,
make migrateup.
