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

// @host 1.116.126.208:8080
func main() {
	orm.InitMysql()
	orm.InitRedis()
	defer orm.Eloquent.Close()

	router := router.InitRouter()
	router.RunTLS(":8080", "certs/www.bigby.love.pem", "certs/www.bigby.love.key")

}
