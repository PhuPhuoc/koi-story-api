package main

import (
	"log"

	"github.com/PhuPhuoc/koi-story-api/api"
	"github.com/PhuPhuoc/koi-story-api/db/mysql"
)

//	@title		Koi Story API
//	@version	1.0

func main() {
	db, appport, host_production := mysql.ConnectDB()
	server_port := ":" + appport
	sv := api.InitServer(server_port, host_production, db)
	if err := sv.RunApp(); err != nil {
		log.Fatal("main: ", err)
	}
}
