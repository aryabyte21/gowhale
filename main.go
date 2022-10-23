package main

import (
	"./database"
	"./routes"
	"net/http"
	"fmt"
	"io"
  	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

    http.ListenAndServe(":5050", nil)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	mux := http.NewServeMux()
 
    mux.HandleFunc("/welcome", welcome)
    http.ListenAndServe(":5050", mux)
	routes.Setup(app)
	http.HandleFunc("/",googleDef)
	app.Listen(":3000")
	http.ListenAndServe(":3000", nil)
}

func googleDef(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w "Haveli me aapka swaagat hai %s", r.URL.Path[1:])
}
