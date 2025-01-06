package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


type Todo struct {
	ID 	  int    `json:"id"`
	Completed bool `json:"completed"`
	Body 	string `json:"body"`

}


func main(){
		

	app := fiber.New()
	fmt.Println("fiber creation✅")

	err:=godotenv.Load(".env")
	if err!=nil {
		log.Fatal(" .env load nhi ho rha mama")
	}

	PORT:=os.Getenv("PORT")

	// list of todos
	todos:=[]Todo{}

	// GET⁡	
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})


	//POST
	app.Post("/api/todos", func (c *fiber.Ctx) error{

		todo:=&Todo{}
		
		if err:=c.BodyParser(todo);
		err!=nil{
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
		if todo.Body==""{
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID=len(todos)+1
		todos=append(todos, *todo)

		
		return c.Status(201).JSON(todo)
		
		
	})


	// UPDATE a todo 
    app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")

        for i, todo := range todos {
            if fmt.Sprint(todo.ID) == id {
                todos[i].Completed = true
                return c.Status(200).JSON(todos[i])
            }
        }

		for temp:= range todos {
			fmt.Println(temp)
		}
		

        return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
    })
	



	app.Delete("/api/todos/:id",func (c *fiber.Ctx) error{
		id:=c.Params("id")
		for i,todo :=range todos {
			if(fmt.Sprint(todo.ID)==id){
				todos=append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success":true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error":"todo not found"})
	})



	log.Fatal(app.Listen(":"+PORT))
}
