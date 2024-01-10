package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lielalmog/be-file-streaming/database"
)

const addr = ":8080"

var app *fiber.App

func Serve() {
	//  set fiber port to 8080

	app = fiber.New()

	setupRouter(app)

	// // fiber to http server
	// r.Server()

	// server = &http.Server{
	// 	Addr:              addr,
	// 	Handler:           r.Server(),
	// 	ReadHeaderTimeout: 3 * time.Second,
	// }

	fmt.Println("Server strating on port", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Shutdown(ctx context.Context) error {
	database.GetDB().Close()

	err := app.Shutdown()

	if err != nil {
		return err
	}

	return nil
}
