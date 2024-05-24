package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

func main() {
	fmt.Println("hello world")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "POST,GET,PUT,PATCH,DELETE",
	}))

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		err := c.BodyParser(todo)

		if err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)

	})

	app.Put("api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		p := new(Todo)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Title = p.Title
				todos[i].Body = p.Body
				todos[i].Done = p.Done
			}
		}

		return c.JSON(todos)

	})

	app.Patch("api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = !todos[i].Done
				break
			}
		}
		return c.JSON(todos)
	})

	app.Delete("api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}
		for i, t := range todos {
			if t.ID == id {
				log.Println("Todo Deleted Successfully")
				todos[i] = todos[len(todos)-1]
				todos = todos[:len(todos)-1]
				break
			}
		}
		return c.JSON(todos)
	})

	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":3000"))

}
