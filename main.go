package main

import (
	orm "github.com/bigby/project/Database"
	router "github.com/bigby/project/Routers"

	_ "github.com/bigby/project/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a server for sorting wrong problems.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 159.138.3.194:8080
func main() {
	orm.InitMysql()
	orm.InitRedis()
	defer orm.Eloquent.Close()

	router := router.InitRouter()
	router.Run()
}
