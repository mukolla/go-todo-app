**Run project** 

go run cmd/main.go 

**Run Docker postgres**

docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

**Install migrate**

$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
$ mv migrate.linux-amd64 $GOPATH/bin/migrate
ln -s $GOPATH/bin/migrate /usr/local/bin/migrate

**Run migrate**

migrate create -ext sql -dir ./schema -seq init


**Migrate up**

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up