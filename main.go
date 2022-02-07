package main

import (
	// "fmt"
	"GO-MAIN/config"
	routes "GO-MAIN/routers"
	"log"
	"net/http"
	"os"

 //	socketio "github.com/googollee/go-socket.io"
)

func main() {



	
	os.Setenv("port", "8000")
	db := config.SetupDB()
	// db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run(":" + os.Getenv("PORT"))

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe( os.Getenv("PORT"), nil))

}
