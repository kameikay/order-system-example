package main

import (
	"database/sql"
	"fmt"

	"github.com/kameikay/order-system-example/configs"
	"github.com/kameikay/order-system-example/internal/infra/web/webserver"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	webserver := webserver.NewWebServer(configs.WebServerPort)
	// webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	// webserver.AddHandler("/endpoint", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start()
}
