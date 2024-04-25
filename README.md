# Fiber Hashing
A hashing library built for Fiber. Supports both argon2id and bcrypt, uses argon2id by default.

## Example
```go
package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/thomasvvugt/fiber-hashing"
)

func main() {
	app := fiber.New()

	hasher := hashing.New()

	app.Get("/", func(c fiber.Ctx) error {
		hash, _ := hasher.CreateHash("ourlittlesecret")
		return c.SendString(hash)
	})

	app.Get("/match/*", func(c fiber.Ctx) error {
		match, _ := hasher.MatchHash("ourlittlesecret", c.Params("*"))
		if match {
			return c.SendString("Matches!")
		}
		return c.SendString("Does not match :c")
	})

	app.Listen(":3000")
}
```
