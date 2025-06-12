package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(
		fiber.Config{
			IdleTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 5,
			ReadTimeout:  time.Second * 5,
		},
	)

	app.Use("/api", func(ctx fiber.Ctx) error {
		fmt.Println("I'm middleware before processing request")
		err := ctx.Next()
		fmt.Println("I'm middleware after processing request")
		return err
	})

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello Fiber")
	})

	if fiber.IsChild() {
		fmt.Println("I'm child process")
	} else {
		fmt.Println("I'm parent process")
	}

	err := app.Listen("localhost:3000", fiber.ListenConfig{
		EnablePrefork: true,
	})

	if err != nil {
		panic(err)
	}
}
