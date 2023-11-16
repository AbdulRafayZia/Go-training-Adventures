package main

import "github.com/godzillaframework/godzilla"

func main() {
	gz := godzilla.New()

	gz.Get("/", func(ctx godzilla.Context) {
		ctx.SendString("Hello world!!!")
	})

	gz.Start(":9090")
}