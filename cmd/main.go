package main

import (
	"booking/internal/app"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	app.Run()
}
