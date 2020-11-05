# REST API Server
Простой REST API Server<br>
База данных - PostgreSQL


## Setup Зависимости
**Бибилотека для чтения toml файлов**
``` sh
$ go get github.com/BurntSushi/toml
```
**Бибилотека для логирования**
``` sh
$ go get github.com/sirupsen/logrus
```
**Бибилотека для routing входящих соединений**
``` sh
$ go get github.com/gorilla/mux
```
**Бибилотека для routing входящих соединений**
``` sh
$ go get github.com/stretchr/testify
```
**Бибилотека для работы с PostgresSQL Golang - driver**
``` sh
$ go get github.com/lib/pq
```
**Бибилотека для migrate БД**
``` sh
$ go get github.com/golang-migrate/migrate
```

## Дополнительные материалы
**Каким образом работать с SQL в Golang | Библиотека database/sql**
``` sh
http://go-database-sql.org/
```

## UpLevelCode
**Бибилотека для DataBase**
+ Позволяет mapping sql запроса в struct golang или использовать именованные параметры и т.д.
``` sh
$ go get github.com/jmoiron/sqlx
```
**Бибилотека ORM для Golang** *пишут, что не очень хорошая из-за производительности*
``` sh
https://gorm.io/index.html
```