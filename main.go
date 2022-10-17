package main

import (
	"./database"
	"./routes"
	"net/http"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	http.HandleFunc("/",googleDef)
	app.Listen(":8000")
	http.ListenAndServe(":8000", nil)
}

func googleDef(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w "Haveli me aapka swaagat hai %s", r.URL.Path[1:])
}
