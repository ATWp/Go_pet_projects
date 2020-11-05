package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

/*
createdb restapi_test
migrate -path migrations - database "postgres://localhost/restapi_test?sslmode=disable" up */
func TestMain(m *testing.M) {
	//Вызывается 1 раз перед всеми тестами
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
