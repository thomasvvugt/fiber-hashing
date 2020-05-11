# Fiber Hashing
A hashing library built for Fiber. Supports both argon2id and bcrypt, uses argon2id by default.

## Example
```go
package main

import (
	"github.com/gofiber/fiber"
	"github.com/thomasvvugt/fiber-hashing"
)

func main() {
	app := fiber.New(&fiber.Settings{
		CaseSensitive: true,
	})

	hasher := hashing.New()

	app.Get("/", func(c *fiber.Ctx) {
		hash, _ := hasher.CreateHash("Password")
		c.Send(hash)
	})

	app.Get("/match/*", func(c *fiber.Ctx) {
		match, _ := hasher.MatchHash("Password", c.Params("*"))
		if match {
			c.Send("Matches!")
			return
		}
		c.Send("Does not match :c")
	})

	app.Listen(3000)
}
```
